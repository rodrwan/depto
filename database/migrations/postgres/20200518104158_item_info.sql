-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists item_info (
  id serial primary key,
  item_id  uuid not null references items(id),

  brand  varchar(30) not null,
  price  decimal,
  unit   decimal,

  purchase_date    timestamptz,
  purchase_place   varchar(20),
  expiration_date  timestamptz,

  created_at  timestamptz default now(),
  updated_at  timestamptz default now(),
  deleted_at  timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists item_info;
-- +goose StatementEnd
