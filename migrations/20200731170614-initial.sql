
-- +migrate Up
CREATE TABLE tweets (
  id bigserial
);

CREATE UNIQUE INDEX tweets_unique_id_idx ON tweets (id);

-- +migrate Down
DROP TABLE tweets;
