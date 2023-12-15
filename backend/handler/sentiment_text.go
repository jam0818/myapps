package handler

import (
	"context"
	"db/model"
	"db/service"
	"encoding/json"
	"log"
	"net/http"
)

type SentimentTextHandler struct {
	svc *service.SentimentTextService
}

func NewSentimentTextHandler(svc *service.SentimentTextService) *SentimentTextHandler {
	return &SentimentTextHandler{
		svc: svc,
	}
}

func (h *SentimentTextHandler) Create(
	ctx context.Context,
	req *model.CreateSentimentTextRequest,
) (
	*model.CreateSentimentTextResponse,
	error,
) {
	sentimentText, err := h.svc.CreateSentimentText(ctx, req.Text)
	if err != nil {
		return nil, err
	}
	return &model.CreateSentimentTextResponse{SentimentText: *sentimentText}, nil
}

func (h *SentimentTextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// プリフライト用。Content-Type（application/json含む）のヘッダーのリクエストを許可する。
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// プリフライト用。POSTメソッドのリクエストを許可する。
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	switch r.Method {
	case http.MethodPost:
		var createRequest model.CreateSentimentTextRequest
		err := json.NewDecoder(r.Body).Decode(&createRequest)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		log.Printf("%v", createRequest)
		// subjectが空文字列の場合、400 Bad Requestを返す
		if createRequest.Text == "" {
			http.Error(w, "Subject cannot be empty", http.StatusBadRequest)
			return
		}

		// TODOを作成し、データベースに保存
		ctx := r.Context()
		createSentimentText, err := h.Create(ctx, &createRequest)
		if err != nil {
			http.Error(w, "Failed to create sentiment text", http.StatusInternalServerError)
			return
		}

		// レスポンスをJSONエンコードして返す
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(createSentimentText)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
