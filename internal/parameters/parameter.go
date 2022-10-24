package parameters

type Parameter interface {
	Name() string
	GetAsJSON() ([]byte, error)
	Set([]byte) error
}
