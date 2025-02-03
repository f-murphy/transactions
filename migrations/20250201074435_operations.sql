-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION transfer_money(from_user_id INTEGER, to_user_id INTEGER, amount NUMERIC(10,2))
RETURNS TEXT AS $$
DECLARE 
	from_user_balance DECIMAL(10, 2);
BEGIN
	IF from_user_id = to_user_id THEN
		RETURN 'Нельзя перевести самому себе денежные средства';
	END IF;

	SELECT balance INTO from_user_balance FROM Users where id = from_user_id;

	IF from_user_balance < amount THEN
		RETURN 'Недостаточно средств';
	END IF;

	UPDATE Users SET Balance = Balance - amount WHERE ID = from_user_id;
	UPDATE Users SET Balance = Balance + amount WHERE ID = to_user_id;

	INSERT INTO Transactions VALUES
	(DEFAULT, from_user_id, to_user_id, amount);

	 RETURN 'Транзакция завершена успешно';
EXCEPTION
    WHEN others THEN
        RETURN 'Произошла ошибка при выполнении транзакции';
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION replenishment(user_id INTEGER, amount DECIMAL(10, 2))
RETURNS TEXT AS $$
BEGIN
	UPDATE Users SET Balance = Balance + amount WHERE ID = user_id;

	IF NOT FOUND THEN
		RETURN 'Пользователь не найден';
	END IF;

	INSERT INTO Transactions VALUES
	(DEFAULT, null, user_id, amount);

	RETURN 'Транзакция завершена успешно';
EXCEPTION
    WHEN others THEN
        RETURN 'Произошла ошибка при выполнении транзакции';
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION transfer_money;
DROP FUNCTION replenishment;
-- +goose StatementEnd
