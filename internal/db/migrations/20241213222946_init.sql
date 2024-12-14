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
  balance INT NOT NULL,
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

CREATE TABLE IF NOT EXISTS "transaction_category" (
  transaction_id INT NOT NULL,
  category_id INT NOT NULL,
  FOREIGN KEY (transaction_id) REFERENCES "transactions" (transaction_id),
  FOREIGN KEY (category_id) REFERENCES "categories" (category_id)
);

CREATE TABLE IF NOT EXISTS "user_wallet" (
  user_id INT NOT NULL,
  wallet_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "users" (user_id),
  FOREIGN KEY (wallet_id) REFERENCES "wallets" (wallet_id)
);

CREATE TABLE IF NOT EXISTS "user_category" (
  user_id INT NOT NULL,
  category_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "users" (user_id),
  FOREIGN KEY (category_id) REFERENCES "categories" (category_id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "transaction_category";
DROP TABLE IF EXISTS "user_wallet";
DROP TABLE IF EXISTS "user_category";
DROP TABLE IF EXISTS "transaction";
DROP TABLE IF EXISTS "wallet";
DROP TABLE IF EXISTS "category";
DROP TABLE IF EXISTS "user";

-- +goose StatementEnd