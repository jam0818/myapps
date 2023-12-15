# 階層構造
```
my_sentiment_analysis_app/
    ├── backend/
    │   ├── main.go             # Goバックエンドのエントリーポイント
    │   ├── database/           # データベース関連のコード
    │   │   ├── db              # データベース接続設定
    │   │   ├── model           # データベースモデル定義
    │   │   ├── service         # データベース操作
    │   ├── handlers/           # HTTPリクエストハンドラー
    │   │   ├── router          # router
    │   ├── utils/              # ユーティリティ関数
    │   │   ├── utils.go        # 一般的なユーティリティ関数
    │   ├── go.mod              # Goモジュール定義
    │   ├── go.sum              # Goモジュールの依存関係
    ├── frontend/
    │   ├── index               # フロントエンドのHTMLファイル
    │   ├── script              # フロントエンドのJavaScriptコード
    │   ├── styles              # フロントエンドのCSSスタイルシート
    ├── python_ml/
    │   ├── sentiment_analysis  # Python感情分析モデル
    ├── README.md               # プロジェクトの説明とドキュメンテーション
```