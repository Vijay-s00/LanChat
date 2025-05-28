-- +goose Up
CREATE TABLE Messages (
    Id SERIAL PRIMARY KEY,
    Name_from VARCHAR(255) NOT NULL,
    Name_to VARCHAR(255) NOT NULL,
    Message TEXT NOT NULL,
    Created_at TIMESTAMP NOT NULL
);
-- +goose Down 
DROP TABLE Messages;
