package aidego

import (
	"fmt"
	"reflect"
)

type Copier struct {
	initialized bool

	sourceKinds              map[int]reflect.Kind
	destinationKinds         map[int]reflect.Kind
	sourceDestinationMapping map[int]int
}

func (copier *Copier) Copy(source, destination interface{}) error {
	if !copier.initialized {
		if err := copier.initialize(source, destination); err != nil {
			return err
		}
	}

	sourceValue, err := copier.getValue(source)
	if err != nil {
		return err
	}

	destinationValue, err := copier.getValue(destination)
	if err != nil {
		return err
	}

	tmpDestination := reflect.Indirect(reflect.New(destinationValue.Type()))

	for k, v := range copier.sourceDestinationMapping {
		switch {
		case copier.destinationKinds[k] == copier.sourceKinds[v] || copier.destinationKinds[k] == reflect.Interface ||
			copier.sourceKinds[v] == reflect.Interface:
			tmpDestination.Field(k).Set(sourceValue.Field(v))
		case copier.sourceKinds[v] == reflect.Ptr:
			tmpDestination.Field(k).Set(sourceValue.Field(v).Elem())
		case copier.destinationKinds[k] == reflect.Ptr:
			tmpDestination.Field(k).Set(sourceValue.Field(v).Addr())
		default:
			return fmt.Errorf("unsupported kind combination: %s x %s", copier.sourceKinds[v].String(),
				copier.destinationKinds[k].String())
		}
	}

	reflect.Indirect(reflect.ValueOf(destination)).Set(tmpDestination)

	return nil
}

func (copier *Copier) initialize(source, destination interface{}) error {
	sourceType, sourceNumFields, err := copier.prepareStructs(source)
	if err != nil {
		return err
	}

	destinationType, destinationNumFields, err := copier.prepareStructs(destination)
	if err != nil {
		return err
	}

	for iDestinationField := 0; iDestinationField < destinationNumFields; iDestinationField++ {
		destinationField := destinationType.Field(iDestinationField)
		copier.destinationKinds[iDestinationField] = destinationField.Type.Kind()

		if tag := destinationField.Tag.Get("copier"); tag != "" {
			if tag == "-" {
				continue
			}

			found := false

			for iSourceField := 0; iSourceField < sourceNumFields; iSourceField++ {
				sourceField := sourceType.Field(iSourceField)
				if sourceField.Name == tag {
					found = true
					copier.sourceDestinationMapping[iDestinationField] = iSourceField
					copier.sourceKinds[iSourceField] = sourceField.Type.Kind()

					break
				}
			}

			if !found {
				return fmt.Errorf("could not find destination field for tag %s", tag)
			}
		} else {
			found := false
			for iSourceField := 0; iSourceField < sourceNumFields; iSourceField++ {
				sourceField := sourceType.Field(iSourceField)
				if sourceField.Name == destinationField.Name {
					found = true
					copier.sourceDestinationMapping[iDestinationField] = iSourceField
					copier.sourceKinds[iSourceField] = sourceField.Type.Kind()

					break
				}
			}

			if !found {
				return fmt.Errorf("could not find on destination field a corresping source field for %s",
					destinationField.Name)
			}
		}
	}

	copier.initialized = true

	return nil
}

func (copier *Copier) prepareStructs(data interface{}) (reflect.Type, int, error) {
	value, err := copier.getValue(data)
	if err != nil {
		return nil, 0, err
	}

	numFields, err := copier.getNumFields(value)
	if err != nil {
		return nil, 0, err
	}

	return value.Type(), numFields, nil
}

func (copier *Copier) getValue(structInterface interface{}) (reflect.Value, error) {
	structKind := reflect.TypeOf(structInterface).Kind()

	switch {
	case structKind == reflect.Struct:
		return reflect.ValueOf(structInterface), nil
	case structKind == reflect.Interface || structKind == reflect.Ptr:
		return copier.getValue(reflect.ValueOf(structInterface).Elem().Interface())
	default:
		return reflect.Value{}, fmt.Errorf("source kind %v is not supported", structKind.String())
	}
}

func (copier *Copier) getNumFields(structValue reflect.Value) (int, error) {
	structKind := reflect.TypeOf(structValue).Kind()

	switch {
	case structKind == reflect.Struct:
		return structValue.NumField(), nil

	case structKind == reflect.Interface || structKind == reflect.Ptr:
		return copier.getNumFields(structValue.Elem())

	default:
		return 0, fmt.Errorf("source kind %v is not supported", structKind.String())
	}
}

func NewCopier() (*Copier, error) {
	return &Copier{
		initialized:              false,
		sourceKinds:              make(map[int]reflect.Kind),
		destinationKinds:         make(map[int]reflect.Kind),
		sourceDestinationMapping: make(map[int]int),
	}, nil
}
