DROP TABLE IF EXISTS password_recovery;

CREATE TABLE password_recovery (
    identifier VARCHAR(255) PRIMARY KEY,
    code VARCHAR(6) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL
);