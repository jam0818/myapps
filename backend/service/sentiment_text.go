package service

import (
	"context"
	"database/sql"
	"db/model"
	"fmt"
	"log"
	"strings"
)

type SentimentTextService struct {
	db *sql.DB
}

func NewSentimentTextService(db *sql.DB) *SentimentTextService {
	return &SentimentTextService{
		db: db,
	}
}

func (s *SentimentTextService) CreateSentimentText(ctx context.Context, text string) (*model.SentimentText, error) {
	const (
		insert  = `INSERT INTO sentiment_text(text, sentiment_score) VALUES(?, ?)`
		confirm = `SELECT text, sentiment_score, created_at, updated_at FROM sentiment_text WHERE id = ?`
	)
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
		}
	}(tx)
	// todo: pythonスクリプトからsentiment_scoreを計算して返す
	//textData := model.TextData{Text: text}
	//requestBody, err := json.Marshal(textData)
	//if err != nil {
	//	return nil, err
	//}
	//response, err := http.Post(
	//	os.Getenv("PYTHON_API_END_POINT")+"analyze_sentiment",
	//	"application/json",
	//	bytes.NewBuffer(requestBody),
	//)
	//if err != nil {
	//	return nil, err
	//}
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//	}
	//}(response.Body)
	//if response.StatusCode != http.StatusOK {
	//	return nil, err
	//}
	sentimentResponse := model.SentimentResponse{SentimentScore: 0}
	//decoder := json.NewDecoder(response.Body)
	//if err := decoder.Decode(&sentimentResponse); err != nil {
	//	return nil, err
	//}
	result, err := tx.ExecContext(
		ctx,
		insert,
		text,
		sentimentResponse.SentimentScore,
	)
	if err != nil {
		return nil, err
	}
	// 挿入したTODOのIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, confirm, id)
	var sentimentText model.SentimentText
	sentimentText.ID = id
	err = row.Scan(
		&sentimentText.Text,
		&sentimentText.SentimentScore,
		&sentimentText.CreatedAt,
		&sentimentText.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &sentimentText, nil
}

func (s *SentimentTextService) DeleteSentimentText(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`
	var query string
	var args []interface{}

	if len(ids) == 0 {
		// ids が空の場合はクエリを実行しない
		return nil
	} else if len(ids) == 1 {
		// ids が 1 つの要素を持つ場合は単一のプレースホルダーを使用
		query = "DELETE FROM todos WHERE id = ?"
		args = append(args, ids[0])
	} else {
		// ids が複数の要素を持つ場合は必要な数のプレースホルダーを生成
		query = fmt.Sprintf(deleteFmt, strings.Repeat(",?", len(ids)-1))
		for _, id := range ids {
			args = append(args, id)
		}
	}
	// ids を []interface{} に変換
	idInterfaces := make([]interface{}, len(ids))
	log.Println(query)
	for i, id := range ids {
		idInterfaces[i] = id
	}

	// トランザクションを開始
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 削除クエリの実行
	result, err := tx.ExecContext(ctx, query, idInterfaces...)
	if err != nil {
		return err
	}

	// 削除された行数を取得
	numRows, err := result.RowsAffected()
	log.Printf("numrows: %v", numRows)
	if err != nil {
		return err
	}

	// 削除された行がなかった場合はエラーを返す
	if numRows == 0 {
		return &model.ErrNotFound{}
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
