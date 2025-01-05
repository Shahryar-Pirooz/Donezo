package domain

import (
	"errors"

	"github.com/google/uuid"
)

const (
	ERROR_ID_NIL = "ID cannot be a Nil UUID"
)

type ProjectID = uuid.UUID

type Project struct {
	UUID   ProjectID
	Name   string
	Parent *Project
}

type ProjectFilter struct {
	Name string
}

func (p *Project) IsValid() error {
	if p.UUID == uuid.Nil {
		return errors.New(ERROR_ID_NIL)
	}
	return uuid.Validate(p.UUID.String())
}
