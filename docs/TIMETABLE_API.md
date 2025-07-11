# Timetable Image Analysis API Documentation

## Overview

The Timetable Image Analysis feature allows users to upload images of their class schedules and automatically extract class names using OCR (Optical Character Recognition) technology. This implementation provides the foundation for Phase 1 MVP with basic text extraction and pattern matching.

## API Endpoints

### 1. Upload Timetable Image

**POST** `/timetable/upload`

Uploads a timetable image and extracts class names from it.

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

**Request Body:**
```json
{
  "image_data": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQ...",
  "file_name": "timetable.jpg"
}
```

**Response:**
```json
{
  "message": "Timetable uploaded and analyzed successfully",
  "data": {
    "id": 1,
    "image_path": "timetable.jpg",
    "original_text": "時間割表\n月曜日\t火曜日\t水曜日\n1限\t数学\t英語\t物理",
    "processing_status": "completed",
    "user_id": 1,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:05Z",
    "extracted_classes": [
      {
        "id": 1,
        "class_name": "数学",
        "confidence": 0.8,
        "position": null,
        "is_confirmed": false
      },
      {
        "id": 2,
        "class_name": "英語",
        "confidence": 0.8,
        "position": null,
        "is_confirmed": false
      }
    ]
  }
}
```

### 2. Get Analysis Result

**GET** `/timetable/analysis/:analysisId`

Retrieves a specific analysis result.

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response:**
```json
{
  "data": {
    "id": 1,
    "image_path": "timetable.jpg",
    "original_text": "時間割表\n月曜日\t火曜日\t水曜日\n1限\t数学\t英語\t物理",
    "processing_status": "completed",
    "user_id": 1,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:05Z",
    "extracted_classes": [...]
  }
}
```

### 3. Get User's Analysis History

**GET** `/timetable/analyses`

Retrieves all analysis results for the authenticated user.

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "image_path": "timetable.jpg",
      "processing_status": "completed",
      "created_at": "2024-01-15T10:30:00Z",
      "extracted_classes": [...]
    },
    {
      "id": 2,
      "image_path": "schedule.png",
      "processing_status": "failed",
      "created_at": "2024-01-14T15:20:00Z",
      "extracted_classes": []
    }
  ]
}
```

### 4. Confirm Classes and Add to Plan

**POST** `/timetable/analysis/:analysisId/confirm`

Confirms selected classes and adds them to a specific plan.

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

**Request Body:**
```json
{
  "class_names": ["数学", "英語", "物理"],
  "plan_id": 5
}
```

**Response:**
```json
{
  "message": "Classes confirmed and added to plan successfully"
}
```

### 5. Delete Analysis

**DELETE** `/timetable/analysis/:analysisId`

Deletes an analysis and its associated extracted classes.

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response:**
```json
{
  "message": "Analysis deleted successfully"
}
```

## Image Format Requirements

- **Supported formats:** JPEG, PNG, GIF
- **Input method:** Base64 encoded data URL
- **File size:** Recommended maximum 5MB
- **Content:** Japanese text timetables work best

### Example Base64 Data URL Format:
```
data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcU...
```

## Class Name Detection

The current implementation uses pattern matching to detect common Japanese subject names:

### Supported Subject Patterns:
- **Basic subjects:** 数学, 英語, 国語, 物理, 化学, 生物, 歴史, 地理
- **Humanities:** 現代文, 古文, 漢文, 世界史, 日本史, 政治経済, 倫理
- **Arts & Skills:** 体育, 音楽, 美術, 書道, 家庭科, 技術, 情報
- **Technology:** コンピュータ, プログラミング, データベース, ネットワーク
- **University level:** 経済学, 経営学, 法学, 心理学, 社会学, 哲学
- **Math advanced:** 微積分, 線形代数, 統計学, 確率論
- **Compound subjects:** コンピュータ基礎, プログラミング基礎, etc.

## Error Handling

### Common Error Responses:

**400 Bad Request:**
```json
{
  "error": "Invalid request format"
}
```

**401 Unauthorized:**
```json
{
  "error": "Invalid or missing JWT token"
}
```

**404 Not Found:**
```json
{
  "error": "analysis not found"
}
```

**500 Internal Server Error:**
```json
{
  "error": "OCR processing failed: tesseract command failed"
}
```

## Processing Status

Analysis processing goes through these states:

- **`pending`**: Analysis record created, processing not started
- **`processing`**: OCR extraction in progress
- **`completed`**: Successfully extracted classes
- **`failed`**: Error occurred during processing

## Usage Example

### Complete Workflow:

1. **Upload Image:**
```bash
curl -X POST http://localhost:8080/timetable/upload \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "image_data": "data:image/jpeg;base64,YOUR_BASE64_IMAGE",
    "file_name": "my_timetable.jpg"
  }'
```

2. **Check Results:**
```bash
curl -X GET http://localhost:8080/timetable/analysis/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

3. **Confirm Classes:**
```bash
curl -X POST http://localhost:8080/timetable/analysis/1/confirm \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "class_names": ["数学", "英語", "物理"],
    "plan_id": 5
  }'
```

## Development Notes

### OCR Service

The system uses Tesseract OCR with the following features:
- **Language support:** Japanese (`jpn`) and English (`eng`)
- **Fallback behavior:** If Tesseract is not installed, mock data is returned for development
- **Image processing:** Automatic format detection and conversion

### Future Enhancements (Phase 2)

- **ML Server Integration:** BERT model for improved class name extraction
- **Image preprocessing:** Contrast adjustment, noise removal, rotation correction
- **Position detection:** X/Y coordinates of detected classes
- **Confidence scoring:** More accurate confidence metrics
- **Batch processing:** Multiple image upload support

## Configuration

### Environment Variables:

- **`UPLOAD_DIR`**: Directory for temporary image storage (default: `./uploads`)
- **Database connection:** Configured via existing DB environment variables

### Dependencies:

- **Tesseract OCR:** Optional for production OCR processing
- **Japanese language pack:** `tesseract-ocr-jpn` for Japanese text recognition