package entity

// The entity will be used for migration and album definition
type Album struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Artist string  `db:"artist"`
	Price  float32 `db:"price"`
}
