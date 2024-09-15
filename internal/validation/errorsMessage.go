package validation

import (
	"github.com/ngocthanh06/ecommerce/pkg/utils"
)

var errorMessages = map[string]string{
	"required": "is required.",
	"email":    "must be a valid email address.",
	"phone":    "must be a valid phone number.",
}

func GetErrorMessage(tag string, field string) string {
	if msg, exists := errorMessages[tag]; exists {
		return utils.ConvertFieldName(field) + " " + msg
	}

	return "Invalid value for " + field
}
