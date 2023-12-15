package router

import (
	"database/sql"
	"db/handler"
	"db/service"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", handler.NewHealthzHandler())
	srv := service.NewSentimentTextService(db)
	mux.Handle("/sentiment_text", handler.NewSentimentTextHandler(srv))
	return mux
}
