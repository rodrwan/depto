-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists items (
  id          uuid primary key default gen_random_uuid(),
  name        varchar(40) not null,
  description text,

  created_at  timestamptz default now(),
  updated_at  timestamptz default now(),
  deleted_at  timestamptz
);

CREATE INDEX if not exists idx_items_name
ON items(name) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists items;
DROP INDEX if exists idx_items_name;
-- +goose StatementEnd
