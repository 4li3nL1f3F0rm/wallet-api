-- +goose Up
-- Seed users
INSERT INTO "user" (email, password, first_name, last_name, age)
VALUES
('user1@example.com', 'password1', 'Alice', 'Smith', 28),
('user2@example.com', 'password2', 'Bob', 'Johnson', 34),
('user3@example.com', 'password3', 'Charlie', 'Williams', 41),
('user4@example.com', 'password4', 'Diana', 'Brown', 23),
('user5@example.com', 'password5', 'Ethan', 'Jones', 37),
('user6@example.com', 'password6', 'Fiona', 'Garcia', 29),
('user7@example.com', 'password7', 'George', 'Martinez', 32),
('user8@example.com', 'password8', 'Hannah', 'Hernandez', 25),
('user9@example.com', 'password9', 'Ivy', 'Moore', 45),
('user10@example.com', 'password10', 'Jack', 'Taylor', 30);

-- Seed wallets (1–4 wallets per user)
INSERT INTO "wallet" (user_id, balance)
VALUES
(1, 1200), (1, 500), 
(2, 800),
(3, 1500), (3, 600), (3, 300), 
(4, 100), 
(5, 900), (5, 450),
(6, 1000),
(7, 2000), 
(8, 300), (8, 50), (8, 100),
(9, 2500), 
(10, 1800);

-- Seed transactions (5–10 transactions per wallet)
INSERT INTO "transaction" (wallet_id, amount, description)
VALUES
(1, -50, 'Groceries'), (1, -200, 'Rent'), (1, -30, 'Utilities'), (1, 300, 'Salary'), (1, -10, 'Coffee'),
(2, -20, 'Snacks'), (2, -100, 'Shopping'), (2, 500, 'Bonus'), (2, -15, 'Lunch'),
(3, -50, 'Fuel'), (3, -200, 'Insurance'), (3, -100, 'Repairs'), (3, 1500, 'Savings Deposit'),
(4, -60, 'Dinner'), (4, 100, 'Gift'),
(5, -90, 'Travel'), (5, -40, 'Groceries'), (5, 100, 'Refund'),
(6, -150, 'Hotel'), (6, -50, 'Snacks'), (6, 500, 'Paycheck'),
(7, -500, 'Car Payment'), (7, -100, 'Electronics'), (7, -200, 'Groceries'), (7, 2000, 'Stock Payout'),
(8, -25, 'Coffee'), (8, -10, 'Parking'), (8, -50, 'Shopping'), (8, 100, 'Freelance Work'),
(9, -300, 'Home Renovation'), (9, -100, 'Groceries'), (9, 2500, 'Salary'),
(10, -400, 'Vacation'), (10, -200, 'Entertainment'), (10, 1800, 'Salary');

-- Seed categories
INSERT INTO "category" (name)
VALUES
('Groceries'), ('Entertainment'), ('Utilities'), ('Rent'), ('Shopping'), ('Travel'), ('Salary'), ('Bonus'), ('Coffee');

-- Seed transaction categories (map transactions to categories)
INSERT INTO "transaction_category" (transaction_id, category_id)
VALUES
(1, 1), (2, 4), (3, 3), (4, 7), (5, 9), 
(6, 9), (7, 5), (8, 8), 
(9, 6), (10, 3);

-- Seed user_wallet relationships
INSERT INTO "user_wallet" (user_id, wallet_id)
VALUES
(1, 1), (1, 2),
(2, 3),
(3, 4), (3, 5), (3, 6),
(4, 7),
(5, 8), (5, 9),
(6, 10),
(7, 11),
(8, 12), (8, 13), (8, 14),
(9, 15),
(10, 16);

-- +goose Down
-- Delete data in reverse order to avoid foreign key violations
DELETE FROM "transaction_category";
DELETE FROM "transaction";
DELETE FROM "user_wallet";
DELETE FROM "wallet";
DELETE FROM "category";
DELETE FROM "user";
