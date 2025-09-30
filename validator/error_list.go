package validator

type ErrorList []Error

func (e *ErrorList) Add(err Error) {
	*e = append(*e, Error{
		Field:   err.Field,
		Message: err.Message,
	})
}
