package msgs

type RendezvousType interface {
	GetType() string
	SetType()
}

type RendezvousID interface {
	SetID(string)
}

type RendezvousTypeID interface {
	RendezvousType
	RendezvousID
}
