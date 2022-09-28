package errors

import "errors"

var (
	ErrUnableToSave               = errors.New("unable to save data")
	ErrUnableToDelete             = errors.New("unable to delete data")
	ErrUserRecordNotFound         = errors.New("user record not found ")
	ErrInvalidRequest             = errors.New("invalid request")
	ErrorUnableToCreate           = errors.New("sorry,something went wrong try again")
	ErrorUnableToConvert          = errors.New("unable to convert")
	ErrorUnableToBindJsonToStruct = errors.New("valid formatted data is required")
	ErrInvalidField               = errors.New("empty field exist")
	ErrOneOrMoreFieldsInvalid     = errors.New("one or more fields are required")
	ErrRecordNotFound             = errors.New("user record not found")
	ErrrHashPasswordMissMatched   = errors.New("hashedPassword is not the hash of the given password")
	ErrInvalidToken               = errors.New("invalid token found")
	ErrUnableToLoginUser          = errors.New("unable to login")
	ErrTransaction                = errors.New("transaction failed")
	ErrUnableToParse              = errors.New("Unable to parse ")
	UnAuthorized                  = errors.New("unauthorized user")
	NoPermission                  = errors.New("no permission")
	ErrorAccessDenied             = errors.New("access denied")
	ErrUnSupportedFormat          = errors.New("unsupported format")
	ErrDNS_FINISHED               = errors.New("not Found")
	UnverifiedAccount             = errors.New("account is not verified")
)

// Descriptions error description
var Descriptions = map[error]string{
	ErrOneOrMoreFieldsInvalid:     "one or more fields are required",
	ErrUnableToSave:               "Unable to save",
	ErrUnableToDelete:             "Unable to delete",
	ErrInvalidRequest:             "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed",
	ErrorUnableToCreate:           "Sorry,Something went wrong try later",
	ErrorUnableToBindJsonToStruct: "valid formatted user data required",
	ErrorUnableToConvert:          "Unable to convert type conversion",
	ErrInvalidField:               "enter valid field values",
	ErrRecordNotFound:             "user record not found",
	ErrrHashPasswordMissMatched:   "crypto/bcrypt: hashedPassword is not the hash of the given password",
	ErrInvalidToken:               "token is not valid",
	ErrUnableToLoginUser:          "User hasn't logged in yet",
	ErrTransaction:                "Transaction failed unexpectedly",
	ErrUnableToParse:              "Unable to parse your token credential",
	UnAuthorized:                  "UnAuthorized User found",
	NoPermission:                  "No valid permmission",
	ErrorAccessDenied:             "You don't have enough permission",
	ErrUnSupportedFormat:          "File format is not supported",
	ErrDNS_FINISHED:               "This site can’t be reached",
	UnverifiedAccount:             "Sorry,Unverified  account found.",
}

// StatusCodes response error HTTP status code
var StatusCodes = map[error]int{
	ErrInvalidRequest:             400,
	ErrOneOrMoreFieldsInvalid:     400,
	ErrUnableToSave:               422,
	ErrUnableToDelete:             422,
	ErrorUnableToCreate:           422,
	ErrorUnableToBindJsonToStruct: 400,
	ErrorUnableToConvert:          403,
	ErrInvalidField:               400,
	ErrRecordNotFound:             400,
	ErrrHashPasswordMissMatched:   400,
	ErrInvalidToken:               400,
	ErrUnableToLoginUser:          401,
	ErrTransaction:                400,
	ErrUnableToParse:              403,
	UnAuthorized:                  401,
	NoPermission:                  403,
	ErrorAccessDenied:             403,
	ErrUnSupportedFormat:          409,
	ErrDNS_FINISHED:               404,
	UnverifiedAccount:             403,
}

// StatusCodes response error HTTP status code
var ErrCodes = map[error]int{
	ErrInvalidRequest:             4000,
	ErrOneOrMoreFieldsInvalid:     4001,
	ErrUnableToSave:               4003,
	ErrUnableToDelete:             4004,
	ErrorUnableToCreate:           4005,
	ErrorUnableToBindJsonToStruct: 4006,
	ErrorUnableToConvert:          4007,
	ErrInvalidField:               4008,
	ErrRecordNotFound:             4009,
	ErrrHashPasswordMissMatched:   4010,
	ErrInvalidToken:               4011,
	ErrUnableToLoginUser:          4012,
	ErrTransaction:                4013,
	ErrUnableToParse:              4014,
	UnAuthorized:                  4015,
	NoPermission:                  4016,
	UnverifiedAccount:             4017,
	ErrorAccessDenied:             4018,
	ErrUnSupportedFormat:          4019,
	ErrDNS_FINISHED:               4020,
}
