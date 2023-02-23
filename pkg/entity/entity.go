package entity

import "github.com/google/uuid"

type Entity interface {
	Update()
	Render()
}

type IdentifiableEntity interface {
	Entity
	SetID(id uuid.UUID)
	GetID() uuid.UUID
}
