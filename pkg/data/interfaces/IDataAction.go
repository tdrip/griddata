package idata

type Action interface {
	GetId() string
	Perform(any) error
}
