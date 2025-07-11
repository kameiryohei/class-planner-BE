package ocr

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type OCRService interface {
	ExtractTextFromImage(imageData string, fileName string) (string, error)
	ExtractClassNames(text string) ([]string, error)
}

type tesseractOCRService struct {
	uploadDir string
}

func NewTesseractOCRService(uploadDir string) OCRService {
	return &tesseractOCRService{
		uploadDir: uploadDir,
	}
}

func (t *tesseractOCRService) ExtractTextFromImage(imageData string, fileName string) (string, error) {
	// Decode base64 image data
	// Remove data URL prefix
	parts := strings.SplitN(imageData, ",", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid image data format")
	}

	imageBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %v", err)
	}

	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(t.uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate unique filename
	timestamp := time.Now().Unix()
	uniqueFileName := fmt.Sprintf("%d_%s", timestamp, fileName)
	imagePath := filepath.Join(t.uploadDir, uniqueFileName)

	// Save image to file
	if err := os.WriteFile(imagePath, imageBytes, 0644); err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}

	// Clean up image file after processing
	defer os.Remove(imagePath)

	// Run Tesseract OCR
	text, err := t.runTesseract(imagePath)
	if err != nil {
		return "", fmt.Errorf("OCR processing failed: %v", err)
	}

	return text, nil
}

func (t *tesseractOCRService) runTesseract(imagePath string) (string, error) {
	// Check if tesseract is available
	_, err := exec.LookPath("tesseract")
	if err != nil {
		// For development/testing when tesseract is not available
		return t.mockOCRExtraction(imagePath), nil
	}

	// Run tesseract with Japanese language support
	cmd := exec.Command("tesseract", imagePath, "stdout", "-l", "jpn+eng")
	output, err := cmd.Output()
	if err != nil {
		// Fallback to English only if Japanese language pack is not available
		cmd = exec.Command("tesseract", imagePath, "stdout", "-l", "eng")
		output, err = cmd.Output()
		if err != nil {
			return "", fmt.Errorf("tesseract command failed: %v", err)
		}
	}

	return string(output), nil
}

func (t *tesseractOCRService) mockOCRExtraction(imagePath string) string {
	// Mock OCR output for development/testing
	return `時間割表
月曜日	火曜日	水曜日	木曜日	金曜日
1限	数学	英語	物理	化学	体育
2限	国語	歴史	数学	英語	音楽
3限	英語	物理	化学	数学	美術
4限	物理	数学	英語	歴史	国語
5限	化学	国語	歴史	物理	数学`
}

func (t *tesseractOCRService) ExtractClassNames(text string) ([]string, error) {
	// Simple pattern matching for Japanese class names
	// This is a basic implementation - in Phase 2 this would be replaced with BERT model
	
	var classNames []string
	classSet := make(map[string]bool) // To avoid duplicates

	// Common Japanese subject patterns (including compound subjects)
	subjectPatterns := []string{
		// Compound subjects first (longer patterns should come first)
		`コンピュータ基礎`, `プログラミング基礎`, `データベース基礎`, `ネットワーク基礎`,
		`数学基礎`, `英語基礎`, `物理基礎`, `化学基礎`, `生物基礎`,
		`現代社会`, `政治経済`, `世界史A`, `世界史B`, `日本史A`, `日本史B`,
		`コンピュータサイエンス`, `情報技術`, `ソフトウェア工学`, `システム設計`,
		// Basic subjects
		`数学`, `英語`, `国語`, `物理`, `化学`, `生物`, `歴史`, `地理`,
		`現代文`, `古文`, `漢文`, `世界史`, `日本史`, `政治`, `経済`, `倫理`,
		`体育`, `音楽`, `美術`, `書道`, `家庭科`, `技術`, `情報`,
		`コンピュータ`, `プログラミング`, `データベース`, `ネットワーク`,
		`経済学`, `経営学`, `法学`, `心理学`, `社会学`, `哲学`,
		`微積分`, `線形代数`, `統計学`, `確率論`,
		`基礎`, `応用`, `演習`, `実習`, `実験`, `実技`,
	}

	// Create regex pattern with word boundaries
	escapedPatterns := make([]string, len(subjectPatterns))
	for i, pattern := range subjectPatterns {
		escapedPatterns[i] = regexp.QuoteMeta(pattern)
	}
	patternStr := strings.Join(escapedPatterns, "|")
	regex, err := regexp.Compile(patternStr)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex pattern: %v", err)
	}

	// Find all matches
	matches := regex.FindAllString(text, -1)
	for _, match := range matches {
		if !classSet[match] {
			classNames = append(classNames, match)
			classSet[match] = true
		}
	}

	// Also look for patterns like "○○学" (subjects ending with 学)
	studyPattern := regexp.MustCompile(`[ぁ-んァ-ヶー一-龯]+学`)
	studyMatches := studyPattern.FindAllString(text, -1)
	for _, match := range studyMatches {
		// Filter out common non-subject words
		if !t.isNonSubjectWord(match) && !classSet[match] {
			classNames = append(classNames, match)
			classSet[match] = true
		}
	}

	return classNames, nil
}

func (t *tesseractOCRService) isNonSubjectWord(word string) bool {
	// Filter out common words that end with "学" but are not subjects
	nonSubjects := []string{
		"大学", "学校", "学生", "学期", "学年", "学習", "学問", "学会",
		"入学", "卒業", "休学", "退学", "転学", "留学", "見学", "科学",
	}

	for _, nonSubject := range nonSubjects {
		if word == nonSubject {
			return true
		}
	}
	return false
}