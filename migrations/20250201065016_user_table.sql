-- +goose Up
CREATE TABLE users (
    ID SERIAL PRIMARY KEY,
	Name VARCHAR NOT NULL,
	Surname VARCHAR NOT NULL,
	Balance DECIMAL(10, 2) NOT NULL DEFAULT 0 CHECK (Balance >= 0)
);

INSERT INTO users VALUES
(DEFAULT, 'Максим', 'Костромов', 400),
(DEFAULT, 'Иван', 'Васильев', 250);

CREATE TABLE transactions (
	ID SERIAL PRIMARY KEY,
	From_user_id INT REFERENCES users(ID),
	To_user_id INT NOT NULL REFERENCES users(ID),
	Amount DECIMAL(10, 2) NOT NULL CHECK (amount > 0),
	Transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose Down
DROP TABLE users;
DROP TABLE transactions;