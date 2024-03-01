package internal

type Factory struct {
}

func NewFactory() (*Factory, error) {
	ft := Factory{}
	return &ft, nil
}

func (f *Factory) Shutdown() error {
	return nil
}
