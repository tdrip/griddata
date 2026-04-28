package idata

// Parser this is the parser interface
type Parser interface {
	GetProcessors() []Processor
	SetProcessors([]Processor)
	AddProcessor(Processor)

	GetSources() []Source
	SetSources([]Source)

	AddSource(Source)

	Execute() error

	Close() error
}
