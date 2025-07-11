package validator

import (
	"backend/model"
	"encoding/base64"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITimetableValidator interface {
	TimetableUploadValidate(req model.TimetableUploadRequest) error
	ConfirmClassesValidate(req model.ConfirmClassesRequest) error
}

type timetableValidator struct{}

func NewTimetableValidator() ITimetableValidator {
	return &timetableValidator{}
}

func (tv *timetableValidator) TimetableUploadValidate(req model.TimetableUploadRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.ImageData,
			validation.Required.Error("image data is required"),
			validation.By(tv.validateBase64Image),
		),
		validation.Field(
			&req.FileName,
			validation.Required.Error("file name is required"),
			validation.Length(1, 255).Error("file name must be between 1 and 255 characters"),
			validation.By(tv.validateImageFileName),
		),
	)
}

func (tv *timetableValidator) ConfirmClassesValidate(req model.ConfirmClassesRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.ClassNames,
			validation.Required.Error("class names are required"),
			validation.Length(1, 50).Error("must have between 1 and 50 class names"),
		),
		validation.Field(
			&req.PlanID,
			validation.Required.Error("plan ID is required"),
			validation.Min(uint(1)).Error("plan ID must be greater than 0"),
		),
	)
}

func (tv *timetableValidator) validateBase64Image(value interface{}) error {
	imageData, ok := value.(string)
	if !ok {
		return validation.NewError("INVALID_TYPE", "image data must be a string")
	}

	// Check if it's a valid base64 data URL
	if !strings.HasPrefix(imageData, "data:image/") {
		return validation.NewError("INVALID_FORMAT", "image data must be a valid data URL with image/ prefix")
	}

	// Extract the base64 part
	parts := strings.SplitN(imageData, ",", 2)
	if len(parts) != 2 {
		return validation.NewError("INVALID_FORMAT", "invalid data URL format")
	}

	// Validate base64 encoding
	_, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return validation.NewError("INVALID_BASE64", "invalid base64 encoding")
	}

	// Check image type
	allowedTypes := []string{"data:image/jpeg", "data:image/jpg", "data:image/png", "data:image/gif"}
	isValidType := false
	for _, allowedType := range allowedTypes {
		if strings.HasPrefix(imageData, allowedType) {
			isValidType = true
			break
		}
	}

	if !isValidType {
		return validation.NewError("INVALID_IMAGE_TYPE", "only JPEG, PNG, and GIF images are allowed")
	}

	return nil
}

func (tv *timetableValidator) validateImageFileName(value interface{}) error {
	fileName, ok := value.(string)
	if !ok {
		return validation.NewError("INVALID_TYPE", "file name must be a string")
	}

	// Check file extension
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	isValidExtension := false
	lowerFileName := strings.ToLower(fileName)
	
	for _, ext := range allowedExtensions {
		if strings.HasSuffix(lowerFileName, ext) {
			isValidExtension = true
			break
		}
	}

	if !isValidExtension {
		return validation.NewError("INVALID_EXTENSION", "file must have a valid image extension (.jpg, .jpeg, .png, .gif)")
	}

	return nil
}