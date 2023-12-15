package model

import "time"

type (
	SentimentText struct {
		ID             int64     `json:"id"`
		Text           string    `json:"text"`
		SentimentScore float64   `json:"sentiment_score"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	TextData struct {
		Text string `json:"text"`
	}

	SentimentResponse struct {
		SentimentScore float64 `json:"sentiment_score"`
	}

	CreateSentimentTextRequest struct {
		Text string `json:"text"`
	}
	CreateSentimentTextResponse struct {
		SentimentText `json:"sentiment_text"`
	}

	ReadSentimentTextRequest struct {
		PrevID int64 `json:"prev_id"`
		Size   int64 `json:"size"`
	}
	ReadSentimentTextResponse struct {
		SentimentTexts []*SentimentText `json:"sentiment_texts"`
	}

	DeleteSentimentTextRequest struct {
		IDs []int64 `json:"ids"`
	}
	DeleteSentimentTextResponse struct {
		Message string `json:"message"`
	}
)
