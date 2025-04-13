-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms (
    room_id SERIAL PRIMARY KEY,
    leader INTEGER REFERENCES users(user_id) NOT NULL,
    current_round INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rooms;
-- +goose StatementEnd
