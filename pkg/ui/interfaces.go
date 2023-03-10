package ui

import "github.com/dfirebaugh/tortuga/pkg/entity"

type Widget interface {
	entity.Entity
}

type Element interface {
	entity.Entity
}
