package model

import "time"

type TimetableAnalysis struct {
	ID               uint                  `json:"id" gorm:"primaryKey"`
	ImagePath        string                `json:"image_path" gorm:"not null"`
	OriginalText     *string               `json:"original_text"`
	ProcessingStatus string                `json:"processing_status" gorm:"default:'pending'"` // pending, processing, completed, failed
	UserID           uint                  `json:"user_id" gorm:"not null"`
	CreatedAt        time.Time             `json:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at"`

	User            User             `json:"user" gorm:"foreignKey:UserID"`
	ExtractedClasses []ExtractedClass `json:"extracted_classes" gorm:"foreignKey:AnalysisID"`
}

type ExtractedClass struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	AnalysisID   uint    `json:"analysis_id" gorm:"not null"`
	ClassName    string  `json:"class_name" gorm:"not null"`
	Confidence   float64 `json:"confidence" gorm:"default:0.0"`
	Position     *string `json:"position"` // JSON string for storing position info like "{\"x\":100,\"y\":200,\"width\":150,\"height\":30}"
	IsConfirmed  bool    `json:"is_confirmed" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Analysis TimetableAnalysis `json:"analysis" gorm:"foreignKey:AnalysisID"`
}

type TimetableAnalysisResponse struct {
	ID               uint                        `json:"id"`
	ImagePath        string                      `json:"image_path"`
	OriginalText     *string                     `json:"original_text"`
	ProcessingStatus string                      `json:"processing_status"`
	UserID           uint                        `json:"user_id"`
	CreatedAt        time.Time                   `json:"created_at"`
	UpdatedAt        time.Time                   `json:"updated_at"`
	ExtractedClasses []ExtractedClassResponse    `json:"extracted_classes"`
}

type ExtractedClassResponse struct {
	ID          uint    `json:"id"`
	ClassName   string  `json:"class_name"`
	Confidence  float64 `json:"confidence"`
	Position    *string `json:"position"`
	IsConfirmed bool    `json:"is_confirmed"`
}

type TimetableUploadRequest struct {
	ImageData string `json:"image_data" validate:"required"` // Base64 encoded image
	FileName  string `json:"file_name" validate:"required"`
}

type ConfirmClassesRequest struct {
	ClassNames []string `json:"class_names" validate:"required"`
	PlanID     uint     `json:"plan_id" validate:"required"`
}