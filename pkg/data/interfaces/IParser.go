package idata

// IParser this is the parser interface
type IParser interface {
	GetProcessors() []Processor
	SetProcessors([]Processor)
	AddProcessor(Processor)

	GetSources() []Source
	SetSources([]Source)

	AddSource(Source)

	Execute() error

	Close() error
}
