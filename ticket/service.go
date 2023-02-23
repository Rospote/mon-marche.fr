package ticket

import (
	"regexp"
	"strconv"
	"strings"
)

// ConvertRawData
// Transforme le payload au format string en un objet TicketHeader et une liste d'objets Produit
// /*
func ConvertRawData(input string) (TicketHeader, []Product) {

	var products []Product

	// Split de l'input par ligne
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	// Recuperation et conversion des éléments de TicketHeader
	// Pour chaque ligne on doit enlever les caracteres de fin de ligne et prendre la donnée à partir de l'espace
	orderId, err := strconv.Atoi(strings.Fields(lines[0][:len(lines[0])-2])[1])
	if err != nil {
		panic(err)
	}
	vat, err := strconv.ParseFloat(strings.Fields(lines[1][:len(lines[1])-2])[1], 32)
	if err != nil {
		panic(err)
	}
	total, err := strconv.ParseFloat(strings.Fields(lines[2][:len(lines[2])-2])[1], 32)
	if err != nil {
		panic(err)
	}
	ticketHeader := TicketHeader{ID: orderId, VAT: float32(vat), Total: float32(total)}

	// Récupération des données liées au Produit

	for i := 5; i < len(lines); i++ {

		// Séparation par la virgules des attributs
		productAttibutes := strings.Split(lines[i], ",")

		var price float64
		var err error

		// On doit enlever le \r à la fin de la ligne sauf pour la derniere ligne de l'input
		if i != len(lines)-1 {
			price, err = strconv.ParseFloat(productAttibutes[2][:len(productAttibutes[2])-2], 32)
			if err != nil {
				panic(err)
			}
		} else {
			price, err = strconv.ParseFloat(productAttibutes[2], 32)
			if err != nil {
				panic(err)
			}
		}

		product := Product{ID: productAttibutes[1], OrderId: orderId, Name: productAttibutes[0], Price: float32(price)}
		products = append(products, product)
	}

	return ticketHeader, products
}

// IsValidInput
// fonction qui permet de valider grossièrement l'input
// On considere que l'input est OK si les headers sont respectés et si le csv est bien formé/*
func IsValidInput(input string) bool {

	// Pour les données de base du ticket on vérifie le header
	regexOrder := regexp.MustCompile(`Order: [0-9]`)
	regexVat := regexp.MustCompile(`VAT: [0-9]`)
	regexTotal := regexp.MustCompile(`Total: [0-9]`)
	regexProductsHeader := regexp.MustCompile(`product,product_id,price`)
	// On vérifie que les lignes produit du CSV contiennent exactement 2 virgules
	regexProducts := regexp.MustCompile(`,`)

	// Separation de l'input par lignes
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	// Verification des headers
	if !(len(lines) > 4 && regexOrder.MatchString(lines[0]) && regexVat.MatchString(lines[1]) && regexTotal.MatchString(lines[2]) && regexProductsHeader.MatchString(lines[4])) {
		return false
	}
	// Verification des produits
	for i := 5; i < len(lines); i++ {
		if len(regexProducts.FindAllStringIndex(lines[i], -1)) != 2 {
			return false
		}
	}
	return true
}
