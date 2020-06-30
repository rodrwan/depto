-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists repositions (
  id             uuid primary key default gen_random_uuid(),
  collection_id  varchar(100) not null,

  created_at  timestamptz default now(),
  updated_at  timestamptz default now(),
  deleted_at  timestamptz
);

CREATE INDEX if not exists idx_repositions_collection_id
ON repositions(collection_id) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists repositions;
DROP INDEX if exists idx_repositions_collection_id;
-- +goose StatementEnd
