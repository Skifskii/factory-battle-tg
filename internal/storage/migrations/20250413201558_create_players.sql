-- +goose Up
-- +goose StatementBegin
CREATE TABLE players (
    player_id SERIAL PRIMARY KEY,
    room_id INTEGER REFERENCES rooms(room_id) NOT NULL,
    user_id INTEGER REFERENCES users(user_id) NOT NULL,
    score INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS players;
-- +goose StatementEnd
