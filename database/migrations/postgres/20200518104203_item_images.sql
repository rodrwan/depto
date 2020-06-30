-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists item_images (
  id serial primary key,
  item_info_id  bigint not null references item_info(id),
  image_url  text,

  created_at  timestamptz default now(),
  updated_at  timestamptz default now(),
  deleted_at  timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists item_images;
-- +goose StatementEnd
