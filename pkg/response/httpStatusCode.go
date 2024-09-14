package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001 // tokein is invalid
	ErrInvaildOTP   = 3002

	// Register code
	ErrcodeUserHasExists = 50001 // user has already registered
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvaildOTP:        "Otp error",
	ErrcodeUserHasExists: "user has already registered",
}
