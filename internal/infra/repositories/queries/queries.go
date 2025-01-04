package queries

const (
	InsertItem = `
		INSERT INTO items (name, description, price, stock, status)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
)
