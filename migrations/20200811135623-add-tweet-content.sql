
-- +migrate Up
ALTER TABLE tweets
  ADD COLUMN content char(140);

-- +migrate Down
ALTER TABLE tweets
  DROP COLUMN content;
