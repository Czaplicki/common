package oogl

type Resource interface {

	GetID() uint32
	Deallocate()
	Use()
}
