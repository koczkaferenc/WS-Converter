package main

import (
	"fmt"
	"log"
	"ws-updater/db"
)

func main() {
	fmt.Println("--- Kulcs-Soft Adatbázis Kapcsolódás ---")

	// 1. Kapcsolat megnyitása
	database, err := db.GetConnection()
	if err != nil {
		log.Fatalf("Sikertelen kapcsolódás: %v", err)
	}
	defer database.Close()

	// 2. Adatok lekérése
	products, err := db.FetchProducts(database)
	if err != nil {
		log.Fatalf("Hiba az adatok lekérésekor: %v", err)
	}

	// 3. Eredmények megjelenítése (vagy továbbküldése a webes felületre)
	for _, p := range products {
		stockVal := 0.0
		if p.Stock.Valid {
			// A .Decimal mezőből hívjuk meg a konverziót
			stockVal = p.Stock.Decimal.InexactFloat64()
		}

		webPriceVal := 0.0
		if p.WebPrice.Valid {
			webPriceVal = p.WebPrice.Decimal.InexactFloat64()
		}

		ReSellerPriceVal := 0.0
		if p.ReSellerPrice.Valid {
			ReSellerPriceVal = p.ReSellerPrice.Decimal.InexactFloat64()
		}

		fmt.Printf("[%s] %s | Készlet: %.2f %s | WebÁr: %.0f Ft. | Viszonteladó: %.0f Ft.\n",
			p.Code, p.Name, stockVal, p.Unit.String, webPriceVal, ReSellerPriceVal)
	}
}
