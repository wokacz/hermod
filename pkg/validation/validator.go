package validation

import "github.com/go-playground/validator/v10"

// Validator is used to validate structs.
var validate = validator.New()

// ErrorResponse struct is used to return validation errors. The FailedField
// field contains the name of the field that failed validation.
type ErrorResponse struct {
	// FailedField is the name of the field that failed validation.
	FailedField string `json:"failedField"`
	// Tag is the name of the validation tag that failed.
	Tag string `json:"tag"`
	// Value is the value of the validation tag.
	Value string `json:"value"`
}

// ValidateStruct function is used to validate a struct. It returns an array of
// errors if any are found. The errors are returned as an array of ErrorResponse
// structs. The ErrorResponse struct contains the name of the field that failed
// validation, the name of the validation tag that failed, and the value of the
// validation tag.
func ValidateStruct(value interface{}) (errors []*ErrorResponse) {
	// Validate the struct.
	err := validate.Struct(value)
	if err != nil {
		// Loop through the errors and add them to the errors array.
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
