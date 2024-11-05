-- +goose Up
-- +goose StatementBegin
CREATE TABLE music (
    id SERIAL PRIMARY KEY,
    song TEXT NOT NULL,
    text_song TEXT NOT NULL,
    release_date DATE NOT NULL,
    link TEXT,
    CONSTRAINT unique_song_group UNIQUE (song, music_group)
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE music_groups (
    music_id INTEGER REFERENCES music(id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
    PRIMARY KEY (music_id, group_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE music;
-- +goose StatementEnd
