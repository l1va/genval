//This file was automatically generated by the genval generator v1.2
//Please don't modify it manually. Edit your entity tags and then
//run go generate

package complicated

import (
	"fmt"

	"github.com/gojuno/genval/errlist"

	"unicode/utf8"
)

type validatable interface {
	Validate() error
}

func validate(i interface{}) error {
	if v, ok := i.(validatable); ok {
		return v.Validate()
	}
	return nil
}

// Validate validates AliasArray
func (r AliasArray) Validate() error {
	var errs errlist.ErrList
	for i, x := range r {
		_ = i
		_ = x
	}
	return errs.ErrorOrNil()
}

// Validate validates AliasChan
func (r AliasChan) Validate() error {
	return nil
}

// Validate validates AliasFunc
func (r AliasFunc) Validate() error {
	return nil
}

// Validate validates AliasOnDogsMapAlias
func (r AliasOnDogsMapAlias) Validate() error {
	if err := DogsMapAlias(r).Validate(); err != nil {
		return fmt.Errorf("%s %v", "r", err)
	}
	return nil
}

// Validate validates AliasString
func (r AliasString) Validate() error {
	return nil
}

// Validate validates Dog
func (r Dog) Validate() error {
	var errs errlist.ErrList
	if utf8.RuneCountInString(r.Name) < 1 {
		errs.AddFieldErrf("Name", "shorter than 1 chars")
	}
	if utf8.RuneCountInString(r.Name) > 64 {
		errs.AddFieldErrf("Name", "longer than 64 chars")
	}
	return errs.ErrorOrNil()
}

// Validate validates DogsMapAlias
func (r DogsMapAlias) Validate() error {
	var errs errlist.ErrList
	for k, v := range r {
		_ = k
		_ = v
		if err := v.Validate(); err != nil {
			errs.AddFieldErr(fmt.Sprintf("r[%v]", k), err)
		}
	}
	return errs.ErrorOrNil()
}

// Validate validates Status
func (r Status) Validate() error {
	switch r {
	case StatusCreated:
	case StatusPending:
	case StatusActive:
	case StatusFailed:
	default:
		return fmt.Errorf("invalid value for enum Status: %v", r)
	}
	return nil
}

// Validate validates User
func (r User) Validate() error {
	var errs errlist.ErrList
	if utf8.RuneCountInString(r.Name) < 3 {
		errs.AddFieldErrf("Name", "shorter than 3 chars")
	}
	if utf8.RuneCountInString(r.Name) > 64 {
		errs.AddFieldErrf("Name", "longer than 64 chars")
	}
	if r.LastName != nil {
		if utf8.RuneCountInString(*r.LastName) < 1 {
			errs.AddFieldErrf("LastName", "shorter than 1 chars")
		}
		if utf8.RuneCountInString(*r.LastName) > 5 {
			errs.AddFieldErrf("LastName", "longer than 5 chars")
		}
	}
	if r.Age < 18 {
		errs.AddFieldErrf("Age", "less than 18")
	}
	if r.Age > 105 {
		errs.AddFieldErrf("Age", "more than 105")
	}
	if r.ChildrenCount == nil {
		errs.AddFieldErrf("ChildrenCount", "cannot be nil")
	} else {
		if *r.ChildrenCount < 0 {
			errs.AddFieldErrf("ChildrenCount", "less than 0")
		}
		if *r.ChildrenCount > 15 {
			errs.AddFieldErrf("ChildrenCount", "more than 15")
		}
	}
	if r.Float < -4.22 {
		errs.AddFieldErrf("Float", "less than -4.22")
	}
	if r.Float > 42.55 {
		errs.AddFieldErrf("Float", "more than 42.55")
	}
	if err := r.Dog.Validate(); err != nil {
		errs.AddFieldErr("Dog", err)
	}
	if r.DogPointer != nil {
		if err := r.DogPointer.Validate(); err != nil {
			errs.AddFieldErr("DogPointer", err)
		}
	}
	if err := r.DogOptional.ValidateOptional(); err != nil {
		errs.AddFieldErr("DogOptional", err)
	}
	if len(r.Urls) < 1 {
		errs.AddFieldErrf("Urls", "less items than 1")
	}
	for i, x := range r.Urls {
		_ = i
		_ = x
		if utf8.RuneCountInString(x) > 256 {
			errs.AddFieldErrf(fmt.Sprintf("Urls[%v]", i), "longer than 256 chars")
		}
	}
	if len(r.Dogs) < 1 {
		errs.AddFieldErrf("Dogs", "less items than 1")
	}
	for i, x := range r.Dogs {
		_ = i
		_ = x
		if x != nil {
			if err := x.Validate(); err != nil {
				errs.AddFieldErr(fmt.Sprintf("Dogs[%v]", i), err)
			}
		}
	}
	if r.Test != nil {
		if len(*r.Test) < 1 {
			errs.AddFieldErrf("Test", "less items than 1")
		}
		for i, x := range *r.Test {
			_ = i
			_ = x
			if x < 4 {
				errs.AddFieldErrf(fmt.Sprintf("Test[%v]", i), "less than 4")
			}
		}
	}
	if err := validateSome(r.Some); err != nil {
		errs.AddFieldErr("Some", err)
	}
	if len(r.SomeArray) < 1 {
		errs.AddFieldErrf("SomeArray", "less items than 1")
	}
	for i, x := range r.SomeArray {
		_ = i
		_ = x
		if err := validateSome(x); err != nil {
			errs.AddFieldErr(fmt.Sprintf("SomeArray[%v]", i), err)
		}
	}
	if len(r.Dict) < 2 {
		errs.AddFieldErrf("Dict", "less items than 2")
	}
	for k, v := range r.Dict {
		_ = k
		_ = v
		if utf8.RuneCountInString(k) > 64 {
			errs.AddFieldErrf(fmt.Sprintf("Dict[%v]", k), "longer than 64 chars")
		}
		if v < -35 {
			errs.AddFieldErrf(fmt.Sprintf("Dict[%v]", k), "less than -35")
		}
		if v > 34 {
			errs.AddFieldErrf(fmt.Sprintf("Dict[%v]", k), "more than 34")
		}
	}
	for k, v := range r.DictDogs {
		_ = k
		_ = v
		if err := v.ValidateOptional(); err != nil {
			errs.AddFieldErr(fmt.Sprintf("DictDogs[%v]", k), err)
		}
		if err := validateMaxDogName(v); err != nil {
			errs.AddFieldErr(fmt.Sprintf("DictDogs[%v]", k), err)
		}
	}
	if err := r.Alias.Validate(); err != nil {
		errs.AddFieldErr("Alias", err)
	}
	if err := r.AliasOnAlias.Validate(); err != nil {
		errs.AddFieldErr("AliasOnAlias", err)
	}
	if err := r.AliasOnAliasWithCustomValidate.ValidateAlias(); err != nil {
		errs.AddFieldErr("AliasOnAliasWithCustomValidate", err)
	}
	for k, v := range r.MapOfMap {
		_ = k
		_ = v
		if len(v) < 1 {
			errs.AddFieldErrf(fmt.Sprintf("MapOfMap[%v]", k), "less items than 1")
		}
		for k, v := range v {
			_ = k
			_ = v
			if utf8.RuneCountInString(v) < 3 {
				errs.AddFieldErrf(fmt.Sprintf("v[%v]", k), "shorter than 3 chars")
			}
		}
	}
	for i, x := range r.ByteArray {
		_ = i
		_ = x
	}
	return errs.ErrorOrNil()
}
