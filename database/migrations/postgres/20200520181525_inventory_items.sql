-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists inventory_items (
  id            uuid primary key default gen_random_uuid(),
  inventory_id  uuid not null references inventory(id),
  item_id       uuid not null references items(id),

  count  integer not null default 0,
  last_use timestamptz,

  created_at  timestamptz default now(),
  updated_at  timestamptz default now(),
  deleted_at  timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists inventory;
-- +goose StatementEnd
