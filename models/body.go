package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type Body struct {
	ID   int  `json:"id" db:"id"`
	Head Head `json:"head" has_one:"head"`
}

// String is not required by pop and may be deleted
func (b Body) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Bodies is not required by pop and may be deleted
type Bodies []Body

// String is not required by pop and may be deleted
func (b Bodies) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Body) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Body) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Body) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
