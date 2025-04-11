package service

type Error struct {
	ControllerErr error
	ServiceErr    error
}

func NewError(controllerErr error, serviceErr error) *Error {
	return &Error{
		ControllerErr: controllerErr,
		ServiceErr:    serviceErr,
	}
}

func (e Error) Error() string {
	if e.ControllerErr == nil {
		return ""
	}
	return e.ControllerErr.Error()
}

func (e Error) Unwrap() []error {
	return []error{e.ControllerErr, e.ServiceErr}
}
