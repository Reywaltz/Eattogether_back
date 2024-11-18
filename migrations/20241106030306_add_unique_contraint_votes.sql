-- +goose Up
-- +goose StatementBegin
ALTER TABLE place_votes 
ADD CONSTRAINT unique_user_vote_places
UNIQUE(room_id, user_id, place_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE place_votes DROP CONSTRAINT unique_user_vote_places;
-- +goose StatementEnd
