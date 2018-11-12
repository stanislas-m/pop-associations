package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Head struct {
	ID     int   `json:"id,omitempty" db:"id"`
	BodyID int   `json:"-" db:"body_id"`
	Body   *Body `json:"body,omitempty" belongs_to:"body"`
}

// String is not required by pop and may be deleted
func (h Head) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Heads is not required by pop and may be deleted
type Heads []Head

// String is not required by pop and may be deleted
func (h Heads) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (h *Head) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: h.BodyID, Name: "BodyID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (h *Head) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (h *Head) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
