package validator

const (
	TypeValidate ErrType = iota
	TypeCritical
)

type ErrType int

type Error struct {
	Field   string
	Message string
	Type    ErrType
}

func NewError(field, message string) Error {
	return Error{
		Field:   field,
		Message: message,
		Type:    TypeValidate,
	}
}

// TODO на случай если валидатор должен вернуть например ошибку выполнения запроса в БД
func NewCriticalError(field, message string) Error {
	return Error{
		Field:   field,
		Message: message,
		Type:    TypeCritical,
	}
}
