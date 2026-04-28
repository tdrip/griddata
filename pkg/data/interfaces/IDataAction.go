package idata

type IDataAction interface {
	GetId() string
	Perform(any) error
}
