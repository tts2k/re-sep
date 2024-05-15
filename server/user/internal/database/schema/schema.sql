CREATE TABLE Users (
	id UUID PRIMARY KEY,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	email TEXT NOT NULL,
	created TEXT NOT NULL,
	updated TEXT NOT NULL,

	UNIQUE(username)
)

CREATE TABLE Tokens (
	id UUID PRIMARY KEY,
	userId UUID NOT NULL,
	expires TEXT NOT NULL,

	FOREIGN KEY(userId) REFERENCES Users(Id)
)
