-- +goose Up
-- +goose StatementBegin
CREATE TABLE place_votes(
    id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(id),
    user_id INTEGER REFERENCES users(id),
    place_id INTEGER REFERENCES places(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE place_votes;
-- +goose StatementEnd
