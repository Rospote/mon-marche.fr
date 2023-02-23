package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	ticket "mon-marche/test/ticket"
	"net/http"
)

func main() {
	r := gin.Default()

	db := ticket.ConnectToDB()

	r.POST("/ticket", func(c *gin.Context) {

		//Recuperation de l'input en tant que raw data
		jsonData, err := c.GetRawData()
		if err != nil {
			panic(err)
		}

		//Conversion en string
		s := string(jsonData)

		// Si l'input est ok, on insere en BDD de maniere classique
		if ticket.IsValidInput(s) {
			ticketHeader, products := ticket.ConvertRawData(s)
			ticket.InsertTicketHeaderToDB(db, ticketHeader)
			// Pour les produits, on utilise une go routine pour améliorer les performances
			for _, product := range products {
				go ticket.InsertProductToDB(db, product)
			}

			c.JSON(http.StatusOK, gin.H{"message": "Ticket inserted in DB"})

		} else { // l'input est incorrect, on l'archive tel quel en base
			ticket.InsertErrorTicketToDb(db, s)

			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "The ticket cannot be processed, registered in DB as an error"})

		}

	})

	r.Run()
}
