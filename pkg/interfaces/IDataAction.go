package igrid

type IDataAction interface {
	GetId() string
	Perform(any) error
}
