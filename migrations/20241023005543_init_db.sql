-- +goose Up
-- +goose StatementBegin
CREATE TABLE places (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    image TEXT
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    role TEXT DEFAULT 'user'
);

CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ,
    external_id UUID UNIQUE
);

CREATE TABLE rooms_places (
    room_id integer REFERENCES rooms(id),
    place_id integer REFERENCES places(id)
);

CREATE TABLE user_rooms (
    user_id integer REFERENCES users(id),
    room_Id integer REFERENCES rooms(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms_places;
DROP TABLE user_rooms;
DROP TABLE places;
DROP TABLE users;
DROP TABLE rooms;
-- +goose StatementEnd
