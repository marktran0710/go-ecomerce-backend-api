package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken:     "Token is invalid",
}
