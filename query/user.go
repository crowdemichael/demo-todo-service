package query

const (
	QueryGetUserProfile = `
		SELECT
			name,
			email
		FROM users
		WHERE id = ?;
	`

	QueryCreateUser = `
		INSERT INTO users (name, email) VALUES (?, ?);
	`
)
