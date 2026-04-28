package iaction

type Action interface {
	GetId() string
	Perform(any) error
}
