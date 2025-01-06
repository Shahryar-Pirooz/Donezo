package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ERROR_ID_NIL = errors.New("ID cannot be a Nil UUID")
	ERROR_ID_VALID = errors.New("ID is not valid")
	ERROR_PARENT = errors.New("Cannot be self Parent")
)

type ProjectID = uuid.UUID

type Project struct {
	UUID     ProjectID
	Name     string
	Parent   ProjectID
	CreateAt time.Time
	DeleteAt time.Time
}

type ProjectFilter struct {
	Name string
}

func (p *Project) IsValid() error {
	if p.UUID == uuid.Nil {
		return ERROR_ID_NIL
	}
	if err:= uuid.Validate(p.UUID.String());err!=nil{
		return fmt.Errorf("%w : %w" , ERROR_ID_VALID , err)
	}
	if p.UUID == p.Parent {
		return ERROR_PARENT
	}
	return nil
}
