CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME(3) NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_users_username (username)
);

CREATE TABLE IF NOT EXISTS notes (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(100) NOT NULL,
    content VARCHAR(2000) NOT NULL,
    created_at DATETIME(3) NULL,
    PRIMARY KEY (id),
    KEY idx_notes_user_id (user_id)
);
