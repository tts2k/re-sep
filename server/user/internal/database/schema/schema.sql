CREATE TABLE Users (
	id UUID PRIMARY KEY,
	username TEXT NOT NULL,
	sub TEXT NOT NULL,
	last_login TIME NOT NULL,
	created_at TIME NOT NULL,
	updated_at TIME NOT NULL,

	UNIQUE(username)
)

CREATE TABLE Tokens (
	id UUID PRIMARY KEY,
	userId UUID NOT NULL,
	expires TIME NOT NULL,
	state TEXT NOT NULL,

	FOREIGN KEY(userId) REFERENCES Users(id)
)
