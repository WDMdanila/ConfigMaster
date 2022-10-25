package parameters

type Parameter interface {
	Name() string
	Value() interface{}
	Set(interface{}) error
}
