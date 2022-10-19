package parameters

type Parameter interface {
	AsJSON() []byte
}

type SequenceParameter interface {
	Parameter
	Update()
}
