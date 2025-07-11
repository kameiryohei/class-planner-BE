package repository

import (
	"backend/model"
	"gorm.io/gorm"
)

type ITimetableRepository interface {
	CreateAnalysis(analysis *model.TimetableAnalysis) error
	GetAnalysisByID(analysisID uint) (*model.TimetableAnalysis, error)
	GetAnalysesByUserID(userID uint) ([]model.TimetableAnalysis, error)
	UpdateAnalysis(analysis *model.TimetableAnalysis) error
	DeleteAnalysis(analysisID uint) error
	CreateExtractedClasses(classes []model.ExtractedClass) error
	UpdateExtractedClasses(classes []model.ExtractedClass) error
	GetExtractedClassesByAnalysisID(analysisID uint) ([]model.ExtractedClass, error)
}

type timetableRepository struct {
	db *gorm.DB
}

func NewTimetableRepository(db *gorm.DB) ITimetableRepository {
	return &timetableRepository{db}
}

func (tr *timetableRepository) CreateAnalysis(analysis *model.TimetableAnalysis) error {
	if err := tr.db.Create(analysis).Error; err != nil {
		return err
	}
	return nil
}

func (tr *timetableRepository) GetAnalysisByID(analysisID uint) (*model.TimetableAnalysis, error) {
	var analysis model.TimetableAnalysis
	if err := tr.db.Preload("User").Preload("ExtractedClasses").First(&analysis, analysisID).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (tr *timetableRepository) GetAnalysesByUserID(userID uint) ([]model.TimetableAnalysis, error) {
	var analyses []model.TimetableAnalysis
	if err := tr.db.Where("user_id = ?", userID).Preload("ExtractedClasses").Find(&analyses).Error; err != nil {
		return nil, err
	}
	return analyses, nil
}

func (tr *timetableRepository) UpdateAnalysis(analysis *model.TimetableAnalysis) error {
	if err := tr.db.Save(analysis).Error; err != nil {
		return err
	}
	return nil
}

func (tr *timetableRepository) DeleteAnalysis(analysisID uint) error {
	// Delete related extracted classes first
	if err := tr.db.Where("analysis_id = ?", analysisID).Delete(&model.ExtractedClass{}).Error; err != nil {
		return err
	}
	
	// Delete the analysis
	if err := tr.db.Delete(&model.TimetableAnalysis{}, analysisID).Error; err != nil {
		return err
	}
	return nil
}

func (tr *timetableRepository) CreateExtractedClasses(classes []model.ExtractedClass) error {
	if len(classes) == 0 {
		return nil
	}
	if err := tr.db.Create(&classes).Error; err != nil {
		return err
	}
	return nil
}

func (tr *timetableRepository) UpdateExtractedClasses(classes []model.ExtractedClass) error {
	for _, class := range classes {
		if err := tr.db.Save(&class).Error; err != nil {
			return err
		}
	}
	return nil
}

func (tr *timetableRepository) GetExtractedClassesByAnalysisID(analysisID uint) ([]model.ExtractedClass, error) {
	var classes []model.ExtractedClass
	if err := tr.db.Where("analysis_id = ?", analysisID).Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}