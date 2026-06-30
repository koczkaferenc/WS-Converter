package SpecialProducts

import (
	"fmt"
	"strconv"
	"strings"
	"ws-updater/models"
)

type part struct {
	partCode    string
	requiredQty float64
	available   int
}

type specialProduct struct {
	code     string
	parts    []part
	category string
	machine  string

	totalPrice  int
	multiplier  float64 // Ennyivel kell felszorozni a részegységek árát
	description string
	available   int
	weight      float64 // A csomag súlya

	imageUrls string // A képek linkjei
}

var specialProducts = map[string]specialProduct{
	"G-MCHALE-1": {
		parts: []part{
			{partCode: "N-KR-4-12B1_Z18_G", requiredQty: 2},
			{partCode: "N-KR-5-16B1_Z29_G", requiredQty: 1},
			{partCode: "N-GL-6-28B2", requiredQty: 1.8},
		},
		category: "McHale",
		machine:  "McHale Pro 1",
		// Vagy az itt megadott ár, ha 0, akkor a részegysége árának összege
		totalPrice: 0,
		// Árszorzó, ha 0, akkor az ár a részárak összege, egyébként az itt megadott érték
		// Fix ár esetén nem számolunk vele
		multiplier: 0,
		// Vagy itt kitöltjük, vagy számítjuk
		description: "McHale furfangoslánc két elemből.",
		// Ha 0, nem dolgozzuk fel, ha 1, akkor megnézzük, hogy a részegységek
		// Mindegyike rendelkezésre áll-e, és csak akkor érhető el.
		available: 0,
	},
}

func computeParms(psWebProducts *[]models.PsProduct, product *specialProduct) bool {
	maxAvailablePkgs := 9999999
	totalWeight := 0.0
	totalPrice := 0.0
	fmt.Println("Részegység:")
	for _, p := range product.parts {
		partFound := false
		qtyInStock := 0
		for _, ps := range *psWebProducts {
			if p.partCode == ps.Reference {
				partFound = true

				qtyInStock, _ = strconv.Atoi(ps.Quantity)
				availablePkgs := int(float64(qtyInStock) / p.requiredQty)
				if availablePkgs < maxAvailablePkgs {
					maxAvailablePkgs = availablePkgs
				}

				parsedPrice, _ := strconv.ParseFloat(ps.UnitPrice, 64)
				currentPrice := parsedPrice * p.requiredQty
				totalPrice += currentPrice * p.requiredQty

				currentWeight, _ := strconv.ParseFloat(ps.Weight, 64)
				totalWeight += currentWeight

				fmt.Printf("- Kód: %s | Rendelkezésre áll: %d %s | Szükséges: %g %s | Egységár: %.0f Ft. | Súly: %g kg\n",
					p.partCode, qtyInStock, ps.Unity,
					p.requiredQty, ps.Unity,
					currentPrice,
					currentWeight)

				product.imageUrls += ps.ImageURLs + ","
			}
		}
		if !partFound {
			fmt.Printf("%s hiányzik.\n", p.partCode)
			return false
		}
	}

	product.available = maxAvailablePkgs
	product.weight = totalWeight
	product.totalPrice = int(totalPrice)
	product.imageUrls = strings.TrimSuffix(product.imageUrls, ",")

	fmt.Printf("Eladható: %d | Csomagár: %d | ÖsszSúly: %.2f kg. | %s\n",
		product.available, product.totalPrice, product.weight, product.imageUrls)
	return true
}

func ProcessSpecialProducts(psWebProducts []models.PsProduct) []models.PsProduct {
	result := []models.PsProduct{}
	// specialProducts felsolgozása
	for s := range specialProducts {
		item := specialProducts[s]
		// Mennyiség, súly, összár kiszámítása
		computeParms(&psWebProducts, &item)

	}
	return result
}
