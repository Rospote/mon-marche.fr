package ticket

import (
	"database/sql"
	"log"
)

// ConnectToDB
// Fonction de connection à la BDD, appelé au début de l'algorithme/**
func ConnectToDB() *sql.DB {

	connStr := "postgresql://postgres:mdp@localhost:5432/tickets?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// InsertTicketHeaderToDB
// Insertion d'un objet TicketHeader dans la table tblTicketHeaders/**
func InsertTicketHeaderToDB(db *sql.DB, ticket TicketHeader) {

	insert := `insert into "tblTicketHeaders"("order_id", "vat","total") values($1, $2, $3)`
	_, err := db.Exec(insert, ticket.ID, ticket.VAT, ticket.Total)
	if err != nil {
		panic(err)
	}
}

// InsertProductToDB
// Insertion d'un objet Product dans la table tblProducts/**
func InsertProductToDB(db *sql.DB, product Product) {

	insert := `insert into "tblProducts"("id", "fk_orderid","productname","price") values($1, $2, $3, $4)`
	_, err := db.Exec(insert, product.ID, product.OrderId, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
}

// InsertErrorTicketToDb
// Insertion des tickets comprenant des erreurs dans la table tblErrorTickets/**

func InsertErrorTicketToDb(db *sql.DB, ticket string) {

	insert := `insert into "tblErrorTickets"("ticket") values($1)`
	_, err := db.Exec(insert, ticket)
	if err != nil {
		panic(err)
	}
}
