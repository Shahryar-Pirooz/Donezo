package domain

import "github.com/google/uuid"

type ProjectID = uuid.UUID

type Project struct {
	UUID     ProjectID
	Name   string
	Parent *Project
}

type ProjectFilter struct {
	Name string
}

func (p *Project) IsValid() bool {
	if p.UUID ==uuid.Nil {
		return false
	}
	return true
}
