package parameters

type Describer interface {
	Describe() map[string]interface{}
}

type Parameter interface {
	Name() string
	Value() interface{}
	Set(interface{}) error
	Describer
}
