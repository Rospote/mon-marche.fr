package ticket

// TicketHeader
// Structure représentant les données propres au ticket/**
type TicketHeader struct {
	ID    int `gorm:"primary_key"`
	VAT   float32
	Total float32
}

// Product
// Structure représentant les données des produits du CSV/**
type Product struct {
	ID      string
	OrderId int
	Name    string
	Price   float32
}
