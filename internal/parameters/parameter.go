package parameters

type Parameter interface {
	ToJSON() []byte
}

type SequenceParameter interface {
	Parameter
	Update()
}
