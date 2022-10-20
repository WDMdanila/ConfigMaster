package parameters

type Parameter interface {
	ToJSON() []byte
}
