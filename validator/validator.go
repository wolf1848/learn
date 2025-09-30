package validator

type Validator struct {
	errors ErrorList
}

func (v *Validator) Validate() ErrorList {
	return v.errors
}

func (v *Validator) Add(errs ...Error) *Validator {
	v.errors = append(v.errors, errs...)

	return v
}

func NewValidator() *Validator {
	return &Validator{}
}
