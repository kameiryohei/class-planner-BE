package usecase

import (
	"backend/model"
	"backend/pkg/ocr"
	"backend/repository"
	"backend/validator"
	"fmt"
	"os"
	"path/filepath"
)

type ITimetableUsecase interface {
	UploadAndAnalyze(req model.TimetableUploadRequest, userID uint) (*model.TimetableAnalysisResponse, error)
	GetAnalysis(analysisID uint, userID uint) (*model.TimetableAnalysisResponse, error)
	GetUserAnalyses(userID uint) ([]model.TimetableAnalysisResponse, error)
	ConfirmClasses(analysisID uint, req model.ConfirmClassesRequest, userID uint) error
	DeleteAnalysis(analysisID uint, userID uint) error
}

type timetableUsecase struct {
	tr repository.ITimetableRepository
	cr repository.ICourseRepository
	tv validator.ITimetableValidator
	ocrService ocr.OCRService
}

func NewTimetableUsecase(
	tr repository.ITimetableRepository,
	cr repository.ICourseRepository,
	tv validator.ITimetableValidator,
) ITimetableUsecase {
	// Get upload directory from environment or use default
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = filepath.Join(".", "uploads")
	}
	
	ocrService := ocr.NewTesseractOCRService(uploadDir)
	
	return &timetableUsecase{
		tr: tr,
		cr: cr,
		tv: tv,
		ocrService: ocrService,
	}
}

func (tu *timetableUsecase) UploadAndAnalyze(req model.TimetableUploadRequest, userID uint) (*model.TimetableAnalysisResponse, error) {
	// Validate request
	if err := tu.tv.TimetableUploadValidate(req); err != nil {
		return nil, err
	}

	// Create analysis record
	analysis := &model.TimetableAnalysis{
		ImagePath:        req.FileName, // Store filename for now, in production would be full path
		ProcessingStatus: "processing",
		UserID:           userID,
	}

	if err := tu.tr.CreateAnalysis(analysis); err != nil {
		return nil, fmt.Errorf("failed to create analysis record: %v", err)
	}

	// Process image with OCR
	extractedText, err := tu.ocrService.ExtractTextFromImage(req.ImageData, req.FileName)
	if err != nil {
		// Update status to failed
		analysis.ProcessingStatus = "failed"
		tu.tr.UpdateAnalysis(analysis)
		return nil, fmt.Errorf("OCR processing failed: %v", err)
	}

	// Extract class names from text
	classNames, err := tu.ocrService.ExtractClassNames(extractedText)
	if err != nil {
		// Update status to failed
		analysis.ProcessingStatus = "failed"
		tu.tr.UpdateAnalysis(analysis)
		return nil, fmt.Errorf("class name extraction failed: %v", err)
	}

	// Update analysis with extracted text
	analysis.OriginalText = &extractedText
	analysis.ProcessingStatus = "completed"
	if err := tu.tr.UpdateAnalysis(analysis); err != nil {
		return nil, fmt.Errorf("failed to update analysis: %v", err)
	}

	// Create extracted class records
	var extractedClasses []model.ExtractedClass
	for _, className := range classNames {
		extractedClass := model.ExtractedClass{
			AnalysisID:  analysis.ID,
			ClassName:   className,
			Confidence:  0.8, // Mock confidence for now
			IsConfirmed: false,
		}
		extractedClasses = append(extractedClasses, extractedClass)
	}

	if len(extractedClasses) > 0 {
		if err := tu.tr.CreateExtractedClasses(extractedClasses); err != nil {
			return nil, fmt.Errorf("failed to save extracted classes: %v", err)
		}
	}

	// Return response
	return tu.buildAnalysisResponse(analysis, extractedClasses), nil
}

func (tu *timetableUsecase) GetAnalysis(analysisID uint, userID uint) (*model.TimetableAnalysisResponse, error) {
	analysis, err := tu.tr.GetAnalysisByID(analysisID)
	if err != nil {
		return nil, fmt.Errorf("analysis not found: %v", err)
	}

	// Check if user owns this analysis
	if analysis.UserID != userID {
		return nil, fmt.Errorf("unauthorized access to analysis")
	}

	return tu.buildAnalysisResponseFromModel(analysis), nil
}

func (tu *timetableUsecase) GetUserAnalyses(userID uint) ([]model.TimetableAnalysisResponse, error) {
	analyses, err := tu.tr.GetAnalysesByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user analyses: %v", err)
	}

	var responses []model.TimetableAnalysisResponse
	for _, analysis := range analyses {
		responses = append(responses, *tu.buildAnalysisResponseFromModel(&analysis))
	}

	return responses, nil
}

func (tu *timetableUsecase) ConfirmClasses(analysisID uint, req model.ConfirmClassesRequest, userID uint) error {
	// Validate request
	if err := tu.tv.ConfirmClassesValidate(req); err != nil {
		return err
	}

	// Get analysis and verify ownership
	analysis, err := tu.tr.GetAnalysisByID(analysisID)
	if err != nil {
		return fmt.Errorf("analysis not found: %v", err)
	}

	if analysis.UserID != userID {
		return fmt.Errorf("unauthorized access to analysis")
	}

	// Create courses from confirmed class names
	for _, className := range req.ClassNames {
		course := &model.Course{
			Name:   className,
			PlanID: req.PlanID,
		}
		if err := tu.cr.CreateCourse(course); err != nil {
			return fmt.Errorf("failed to create course '%s': %v", className, err)
		}
	}

	// Mark extracted classes as confirmed
	extractedClasses, err := tu.tr.GetExtractedClassesByAnalysisID(analysisID)
	if err != nil {
		return fmt.Errorf("failed to get extracted classes: %v", err)
	}

	for i := range extractedClasses {
		for _, confirmedClass := range req.ClassNames {
			if extractedClasses[i].ClassName == confirmedClass {
				extractedClasses[i].IsConfirmed = true
			}
		}
	}

	if err := tu.tr.UpdateExtractedClasses(extractedClasses); err != nil {
		return fmt.Errorf("failed to update extracted classes: %v", err)
	}

	return nil
}

func (tu *timetableUsecase) DeleteAnalysis(analysisID uint, userID uint) error {
	// Get analysis and verify ownership
	analysis, err := tu.tr.GetAnalysisByID(analysisID)
	if err != nil {
		return fmt.Errorf("analysis not found: %v", err)
	}

	if analysis.UserID != userID {
		return fmt.Errorf("unauthorized access to analysis")
	}

	// Delete analysis (cascade will handle extracted classes)
	if err := tu.tr.DeleteAnalysis(analysisID); err != nil {
		return fmt.Errorf("failed to delete analysis: %v", err)
	}

	return nil
}

func (tu *timetableUsecase) buildAnalysisResponse(analysis *model.TimetableAnalysis, extractedClasses []model.ExtractedClass) *model.TimetableAnalysisResponse {
	var classResponses []model.ExtractedClassResponse
	for _, class := range extractedClasses {
		classResponses = append(classResponses, model.ExtractedClassResponse{
			ID:          class.ID,
			ClassName:   class.ClassName,
			Confidence:  class.Confidence,
			Position:    class.Position,
			IsConfirmed: class.IsConfirmed,
		})
	}

	return &model.TimetableAnalysisResponse{
		ID:               analysis.ID,
		ImagePath:        analysis.ImagePath,
		OriginalText:     analysis.OriginalText,
		ProcessingStatus: analysis.ProcessingStatus,
		UserID:           analysis.UserID,
		CreatedAt:        analysis.CreatedAt,
		UpdatedAt:        analysis.UpdatedAt,
		ExtractedClasses: classResponses,
	}
}

func (tu *timetableUsecase) buildAnalysisResponseFromModel(analysis *model.TimetableAnalysis) *model.TimetableAnalysisResponse {
	var classResponses []model.ExtractedClassResponse
	for _, class := range analysis.ExtractedClasses {
		classResponses = append(classResponses, model.ExtractedClassResponse{
			ID:          class.ID,
			ClassName:   class.ClassName,
			Confidence:  class.Confidence,
			Position:    class.Position,
			IsConfirmed: class.IsConfirmed,
		})
	}

	return &model.TimetableAnalysisResponse{
		ID:               analysis.ID,
		ImagePath:        analysis.ImagePath,
		OriginalText:     analysis.OriginalText,
		ProcessingStatus: analysis.ProcessingStatus,
		UserID:           analysis.UserID,
		CreatedAt:        analysis.CreatedAt,
		UpdatedAt:        analysis.UpdatedAt,
		ExtractedClasses: classResponses,
	}
}