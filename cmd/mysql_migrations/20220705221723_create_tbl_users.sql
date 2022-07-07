-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(30) NOT NULL, 
    password VARCHAR(255) NOT NULL, 
    firstname VARCHAR(30) NOT NULL,
    lastname VARCHAR(30) NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    user_role VARCHAR(30) NOT NULL,
    profile_picture varchar(255) NOT NULL,
    is_deleted int default 0,
    refresh_token varchar(255),
    reset_password varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
