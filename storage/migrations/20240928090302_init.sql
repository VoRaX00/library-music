-- +goose Up
-- +goose StatementBegin
CREATE TABLE music (
    id SERIAL PRIMARY KEY,
    song TEXT NOT NULL,
    text_song TEXT NOT NULL,
    release_date DATE NOT NULL,
    link TEXT
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE music_groups (
    music_id INTEGER REFERENCES music(id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
    PRIMARY KEY (music_id, group_id)
);

CREATE INDEX idx_music_groups_music_id ON music_groups(music_id);
CREATE INDEX idx_music_groups_group_id ON music_groups(group_id);
CREATE INDEX idx_music_song ON music(song);
CREATE INDEX idx_group_name ON groups(name);
CREATE INDEX idx_music_song_group ON music_groups(music_id, group_id);
CREATE UNIQUE INDEX idx_unique_group_name ON groups(name);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE music_groups;
DROP TABLE music;
DROP TABLE groups;
-- +goose StatementEnd
