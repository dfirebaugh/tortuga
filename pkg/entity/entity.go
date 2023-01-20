package entity

import "github.com/google/uuid"

type Entity interface {
	Update()
	Render()
	SetID(id uuid.UUID)
}
