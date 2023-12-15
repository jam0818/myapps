CREATE TABLE IF NOT EXISTS sentiment_text
(
    id              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    text            TEXT    NOT NULL DEFAULT '',
    sentiment_score REAL    NOT NULL,
    created_at  DATETIME NOT NULL DEFAULT (DATETIME('now')),
    updated_at  DATETIME NOT NULL DEFAULT (DATETIME('now'))
);

CREATE TRIGGER IF NOT EXISTS trigger_todos_updated_at AFTER UPDATE ON sentiment_text
BEGIN
    UPDATE sentiment_text SET updated_at = DATETIME('now') WHERE id == NEW.id;
END;