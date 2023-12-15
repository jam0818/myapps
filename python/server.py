from fastapi import FastAPI, HTTPException
from sentiment_text.jsonclass import (TextData)

app = FastAPI()


@app.post("/analyze_sentiment/")
async def analyze_sentiment(text_data: TextData):
    try:
        text = text_data.text
        # ここで感情分析を実行し、感情スコアを計算
        sentiment_score = 0.75  # 仮の感情スコア
        return {"sentiment_score": sentiment_score}
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
