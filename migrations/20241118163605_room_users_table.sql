-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms_users( 
    ID SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(id),
    user_id INTEGER REFERENCES users(id),
    UNIQUE (room_id, user_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms_users;
-- +goose StatementEnd
