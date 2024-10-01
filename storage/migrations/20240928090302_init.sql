-- +goose Up
-- +goose StatementBegin
CREATE TABLE music (
    id SERIAL PRIMARY KEY,
    music_group TEXT NOT NULL,
    song TEXT NOT NULL,
    text_song TEXT NOT NULL,
    link TEXT,
    CONSTRAINT unique_song_group UNIQUE (song, music_group)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE music;
-- +goose StatementEnd
