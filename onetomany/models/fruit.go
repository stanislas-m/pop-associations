package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Fruit struct {
	ID     int   `json:"id,omitempty" db:"id"`
	TreeID int   `json:"-" db:"tree_id"`
	Tree   *Tree `json:"tree,omitempty" belongs_to:"tree"`
}

// String is not required by pop and may be deleted
func (f Fruit) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Fruits is not required by pop and may be deleted
type Fruits []Fruit

// String is not required by pop and may be deleted
func (f Fruits) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Fruit) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: f.TreeID, Name: "TreeID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Fruit) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Fruit) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
