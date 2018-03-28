package lib

import (
	"io/ioutil"

	"github.com/lestrrat/go-libxml2"
	"github.com/lestrrat/go-libxml2/xsd"
)

// ValidateXMLWithXSD REQUIRE THEM TO DOCUMENT THIS FUNCTION
func ValidateXMLWithXSD(xmlPath string, xsdPath string) error {
	xsdFile, err := ioutil.ReadFile(xsdPath)
	if err != nil {
		return err
	}

	schema, err := xsd.Parse(xsdFile)
	if err != nil {
		return err
	}
	defer schema.Free()

	xmlFile, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		return err
	}

	doc, err := libxml2.Parse(xmlFile)
	if err != nil {
		return err
	}

	err = schema.Validate(doc)
	if err != nil {
		return err
	}

	return nil
}
