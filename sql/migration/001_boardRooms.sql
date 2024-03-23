-- +goose Up

CREATE TABLE boardRooms (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose Down
DROP TABLE boardRooms;