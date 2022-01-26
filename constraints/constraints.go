package constraints

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}