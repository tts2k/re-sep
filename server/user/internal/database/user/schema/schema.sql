CREATE TABLE IF NOT EXISTS Users (
	id UUID PRIMARY KEY,
	name TEXT NOT NULL,
	sub TEXT NOT NULL,
	config BLOB,
	last_login DATETIME NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,

	UNIQUE(sub)
);

CREATE VIEW IF NOT EXISTS v_user_config
AS
SELECT
	sub,
	JSON(config) as config
FROM Users;

CREATE TRIGGER IF NOT EXISTS user_config_update
INSTEAD OF UPDATE OF config ON v_user_config
BEGIN
	UPDATE Users SET config = JSONB(NEW.config)
	WHERE sub = NEW.sub;
END;
