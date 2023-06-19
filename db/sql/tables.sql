CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    birthday TEXT,
    gender TEXT,
    photo_url TEXT,
    current_time INTEGER,
    active INTEGER,
    updated_at INTEGER
);


CREATE TABLE IF NOT EXISTS todo_list (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    created_at INTEGER,
    updated_at INTEGER,
    completed BOOLEAN,
    FOREIGN KEY (user_id) REFERENCES users (id)
);