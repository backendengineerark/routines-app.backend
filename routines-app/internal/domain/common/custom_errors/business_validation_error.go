package customerrors

type BusinessValidationError struct {
	Message string
}

func (bve *BusinessValidationError) Error() string {
	return bve.Message
}
