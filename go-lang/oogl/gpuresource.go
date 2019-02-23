package oogl

type GPUResource interface {

	GetID() uint32
	Deallocate()
	Use()
}
