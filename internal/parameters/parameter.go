package parameters

type Parameter interface {
	Name() string
	GetAsJSON() []byte
	Set(interface{})
}
