-- +goose Up

ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down

ALTER TABLE feeds DROP COLMN last_fetched_at ;

