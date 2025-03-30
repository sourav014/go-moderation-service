package helpers

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetJSONFieldName(obj interface{}, fieldName string) string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if field, found := objType.FieldByName(fieldName); found {
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			return strings.Split(jsonTag, ",")[0]
		}
	}
	return fieldName
}

func FormatValidationErrors(obj interface{}, err error) []ValidationErrorResponse {
	var errors []ValidationErrorResponse

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			jsonKey := GetJSONFieldName(obj, fieldErr.StructField())

			errors = append(errors, ValidationErrorResponse{
				Field:   jsonKey,
				Message: "The field '" + jsonKey + "' is required",
			})
		}
	}
	return errors
}

func HandleValidationError(ctx *gin.Context, obj interface{}, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": FormatValidationErrors(obj, validationErrors),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
