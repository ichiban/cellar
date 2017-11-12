-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE bottles (
  id INTEGER PRIMARY KEY,
  account_id INTEGER REFERENCES accounts(id) NOT NULL,
  rating INTEGER NOT NULL DEFAULT 3,
  name TEXT NOT NULL DEFAULT '',
  vineyard TEXT NOT NULL DEFAULT '',
  varietal TEXT NOT NULL DEFAULT '',
  vintage INTEGER NOT NULL DEFAULT 1900,
  color TEXT NOT NULL DEFAULT '',
  sweetness INTEGER,
  country TEXT,
  region TEXT,
  review TEXT,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementBegin
CREATE TRIGGER bottles_updated_at UPDATE ON bottles
BEGIN
  UPDATE bottles SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;
-- +goose StatementEnd

CREATE INDEX bottles_account_id ON bottles(account_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE bottles;
