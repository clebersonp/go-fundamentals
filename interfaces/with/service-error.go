package with

type ServiceError string

// Error is the error type interface and this type implements the Error method. So this will be treated as an error type
func (s ServiceError) Error() string {
	return string(s)
}
