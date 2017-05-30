//This file was automatically generated by the genval generator v1.0
//Please don't modify it manually. Edit your entity tags and then
//run go generate

package overriding

import (
	"fmt"
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

// Validate validates Age1
func (r Age1) Validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

// Validate validates Age2
func (r Age2) Validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

// Validate validates Age3
func (r Age3) Validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

// Validate validates Age4
func (r Age4) Validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

// Validate validates Age5
func (r Age5) Validate() error {
	if r.Value < 3 {
		return fmt.Errorf("field Value is less than 3 ")
	}
	if r.Value > 64 {
		return fmt.Errorf("field Value is more than 64 ")
	}
	return nil
}

// Validate validates Request2
func (r Request2) Validate() error {
	if err := r.Age.ValidateMin10(); err != nil {
		return fmt.Errorf("Age is not valid: %v", err)
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}

// Validate validates Request3
func (r Request3) Validate() error {
	if err := validateMin10(r.Age); err != nil {
		return fmt.Errorf("Age is not valid: %v", err)
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}

// Validate validates Request4
func (r Request4) Validate() error {
	if err := r.Age.ValidateMin10(); err != nil {
		return fmt.Errorf("Age is not valid: %v", err)
	}
	if err := validateMax128(r.Age); err != nil {
		return fmt.Errorf("Age is not valid: %v", err)
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return nil
}

// Validate validates Request5
func (r Request5) Validate() error {
	if err := r.Age.Validate(); err != nil {
		return fmt.Errorf("Age is not valid: %v", err)
	}
	if r.Some < 3 {
		return fmt.Errorf("field Some is less than 3 ")
	}
	if r.Some > 64 {
		return fmt.Errorf("field Some is more than 64 ")
	}
	return r.validate()
}
