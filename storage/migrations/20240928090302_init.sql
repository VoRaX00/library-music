-- +goose Up
-- +goose StatementBegin
CREATE TABLE music (
    id SERIAL PRIMARY KEY,
    music_group TEXT NOT NULL,
    song TEXT NOT NULL UNIQUE ,
    text_song TEXT NOT NULL,
    link TEXT
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE music;
-- +goose StatementEnd
