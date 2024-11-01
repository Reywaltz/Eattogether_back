-- +goose Up
-- +goose StatementBegin

DROP TABLE user_rooms;

ALTER TABLE rooms
ADD COLUMN owner_id INTEGER, ADD CONSTRAINT fk_owner_id
FOREIGN KEY (owner_id) REFERENCES users(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE rooms
DROP CONSTRAINT fk_owner_id;

ALTER TABLE rooms
DROP COLUMN owner_id;

CREATE TABLE user_rooms (
    user_id integer REFERENCES users(id),
    room_Id integer REFERENCES rooms(id)
);
-- +goose StatementEnd
