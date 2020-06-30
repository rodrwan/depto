-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists inventory (
  id  uuid primary key default gen_random_uuid(),

  created_at  timestamptz default now(),
  updated_at  timestamptz default now(),
  deleted_at  timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists inventory;
-- +goose StatementEnd
