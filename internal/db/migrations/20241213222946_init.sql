-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
  user_id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  age INT
);

CREATE TABLE IF NOT EXISTS "wallets" (
  wallet_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  balance INT NOT NULL DEFAULT 0,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  FOREIGN KEY (user_id) REFERENCES "users" (user_id)
);

CREATE TABLE IF NOT EXISTS "transactions" (
  transaction_id SERIAL PRIMARY KEY,
  wallet_id INT NOT NULL,
  amount INT NOT NULL,
  description VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (wallet_id) REFERENCES "wallets" (wallet_id)
);

CREATE TABLE IF NOT EXISTS "categories" (
  category_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "transactions" CASCADE;
DROP TABLE IF EXISTS "wallets" CASCADE;
DROP TABLE IF EXISTS "categories" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;

-- +goose StatementEnd