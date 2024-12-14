-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user" (
  user_id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  age INT
);

CREATE TABLE IF NOT EXISTS "wallet" (
  wallet_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  balance INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "user" (user_id)
);

CREATE TABLE IF NOT EXISTS "transaction" (
  transaction_id SERIAL PRIMARY KEY,
  wallet_id INT NOT NULL,
  amount INT NOT NULL,
  description VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (wallet_id) REFERENCES "wallet" (wallet_id)
);

CREATE TABLE IF NOT EXISTS "category" (
  category_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "transaction_category" (
  transaction_id INT NOT NULL,
  category_id INT NOT NULL,
  FOREIGN KEY (transaction_id) REFERENCES "transaction" (transaction_id),
  FOREIGN KEY (category_id) REFERENCES "category" (category_id)
);

CREATE TABLE IF NOT EXISTS "user_wallet" (
  user_id INT NOT NULL,
  wallet_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "user" (user_id),
  FOREIGN KEY (wallet_id) REFERENCES "wallet" (wallet_id)
);

CREATE TABLE IF NOT EXISTS "user_category" (
  user_id INT NOT NULL,
  category_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "user" (user_id),
  FOREIGN KEY (category_id) REFERENCES "category" (category_id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "wallet";
DROP TABLE IF EXISTS "transaction";
DROP TABLE IF EXISTS "category";
DROP TABLE IF EXISTS "transaction_category";
DROP TABLE IF EXISTS "user_wallet";
DROP TABLE IF EXISTS "user_category";

-- +goose StatementEnd