package helpers

import "github.com/go-playground/validator/v10"

func ValidateRequest(requestData interface{}) []string {
	validate := validator.New()
	err := validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, e.Field()+" is "+e.Tag())
		}
		return errorMessages
	}
	return nil
}
