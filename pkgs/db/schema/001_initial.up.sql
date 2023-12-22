CREATE TABLE IF NOT EXISTS
    device (
        id INTEGER PRIMARY KEY,
        title TEXT NOT NULL,
        address TEXT NOT NULL,
        connect_type TEXT CHECK (connect_type IN ('ssh', 'telnet')) NOT NULL,
        auth_id INTEGER,
        note TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime')),
        updated_at TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS
    device_auth (
        id INTEGER PRIMARY KEY,
        title TEXT NOT NULL,
        type TEXT CHECK (type IN ('username_password', 'private_key')) NOT NULL,
        username TEXT,
        password TEXT,
        private_key TEXT,
        created_at TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime')),
        updated_at TIMESTAMP
    );