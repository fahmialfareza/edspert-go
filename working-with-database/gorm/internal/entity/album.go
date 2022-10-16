package entity

// The entity will be used for migration and album definition
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}
