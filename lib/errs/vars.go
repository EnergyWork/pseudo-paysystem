package errs

var (
	ErrInternal = New().SetCode(Internal)
)
