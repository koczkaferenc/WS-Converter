package gl

//*******************************************************************
// Flyer láncok
//*******************************************************************

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

func ProcessFlyer(p models.KsProduct, prodCodes *[]string) models.PsProduct {
	var (
		match          []string
		family         string
		manufacturerId string
		pStr           string

		kivitelkod string
	)

	psp := models.PsProduct{}
	features := map[string]string{
		"FlFamily":    "",     // Kereséshez a lánc ismert megnevezése
		"Anyag":       "",     // Acél
		"Osztás":      "",     // Típus: Agyas lánckerék | Laplánckerék
		"Hevederszam": "",     // Hevederek száma
		"Flyer":       "Igen", // "Igen", ha flyer lánc
	}

	psp.ID = ""            // Az ID üres, cikkszámokkal dogozunk
	psp.Reference = p.Code // Ez a cikkszám, ID helyett használjuk
	psp.Active = "1"
	psp.PriceTaxExcluded = fmt.Sprintf("%.0f", p.WebPrice)
	psp.UnitPrice = fmt.Sprintf("%.0f", p.WebPrice) // TODO
	psp.TaxRulesID = "1"                            // ÁFA kulcs 27%
	psp.Quantity = fmt.Sprintf("%.1f", p.Stock)     // Mennyiség
	psp.AvailableForOrder = "1"
	psp.Weight = fmt.Sprintf("%.1f", p.Weight)
	psp.Unity = "m"
	psp.TextInStock = "m"          // méter
	psp.TextBackorderAllowed = "m" // méter
	psp.Categories = "Láncok/Flyer"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// N-FL-7-LL1044
	// N-(FL)-([0-9]+)-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$
	//    1      2      3            4       5      6
	//  fam    Manuf    --LL/BL--   osztas  -hevederek-
	match = regExpFL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		kivitelkod = match[3] // LL | BL
		pStr = fmt.Sprintf("%s-%s-%sx%s", kivitelkod, match[4], match[5], match[6])

		features["Anyag"] = "Acél"
		features["Osztás"] = flyerOsztasTab[match[4]]
		features["Hevederszam"] = fmt.Sprintf("%sx%s", match[5], match[6])
		psp.Name = fmt.Sprintf("%s Flyer Lánc", pStr)
		psp.Description = fmt.Sprintf("%s %s %s flyer lánc", psp.Manufacturer, kivitelkod, features["Hevederszam"])
		psp.Tags = "Flyer lánc," + pStr
		psp.MetaKeywords = pStr + ",flyer lánc"
		psp.Categories = "Láncok/Flyer"
	}

	// `N-(FLCS)-([0-9]+)-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$`
	//     1       2        3           4       5      6
	//    fam     Manuf    --LL/BL--    osztas  -hevederek-
	// N-FLCS-7-BL646
	match = regExpFLCS.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		kivitelkod = match[3] // LL | BL
		pStr = fmt.Sprintf("%s %s%s%s", kivitelkod, match[4], match[5], match[6])

		features["Anyag"] = "Acél"
		features["Osztás"] = flyerOsztasTab[match[4]]
		features["Hevederszam"] = fmt.Sprintf("%sx%s", match[5], match[6])
		psp.Name = fmt.Sprintf("%s Flyer Lánc Csap", pStr)
		psp.Description = fmt.Sprintf("%s %s %s flyer lánc csap", psp.Manufacturer, kivitelkod, features["Hevederszam"])
		psp.Tags = "Flyer lánc," + pStr
		psp.MetaKeywords = pStr + ",flyer lánc csap"
		psp.Categories = "Láncok/Flyer, Láncok/Patentszemek"
	}

	// Típus beállítása a kereséshez
	if pStr != "" {
		features["FlFamily"] = pStr
	}

	// Gyártó beállítása
	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

	psp.Summary = psp.Description
	qtyTmp, _ := strconv.ParseFloat(psp.Quantity, 64)
	if qtyTmp == 0 {
		// Lehet rendelni
		psp.OutOfStockAction = "2"
		psp.Summary += models.JelenlegNemElerheto
		psp.Summary += "<hr>Szállítási idő: kb. 14 nap."
	} else {
		// Ha van raktáron, az beragadt, ezért 5%-os engedménnyel akciózzuk.
		// psp.OnSale = "1"
		// psp.DiscountPercent = models.KiarusitasSzazalek
		// Most van belőle, lehet rendelni
		psp.OutOfStockAction = "0"
		psp.Summary += fmt.Sprintf("<hr>Készleten: %g %s", qtyTmp, psp.Unity)
		psp.Summary += fmt.Sprintf("<hr>Súly: %s kg.", psp.Weight)
		psp.Summary += "<hr>Szállítási idő: 1-2 nap."
	}
	psp.Summary += models.Zaradek
	// Ebben a fázisában kell beállítani és nem lehet pont a végén.
	psp.MetaDescription = strings.TrimRight(psp.Description, ".")
	psp.MetaTitle = psp.Name
	psp.URLRewritten = p.Code

	// Képek előállítása (a Velonál egyedileg készült)
	if psp.ImageURLs == "" {
		psp.ImageURLs = fmt.Sprintf(
			"%s/N-%s.png,%s/D-%s.png",
			models.ImagesBase, family,
			models.ImagesBase, family)
		psp.ImageAltTexts = psp.Name
	}

	// Speciális tulajdonságok beállítása
	psp.Features = models.MkFeaturesList(features)

	// Kapcsolódó termékek
	psp.Accessories = ""
	// TODO Ideiglenesen kivesszük, mert ha nem ltezik a termék, nagyon lelassul
	//rgxStr := fmt.Sprintf(`^N-GL-[0-9]+-%s%s.*`, productType, sorokszama)
	//psp.Accessories += models.getRelatedProductIds(rgxStr, prodCodes)

	// Rendelhető?
	psp.AvailableForOrder = "1"
	if slices.Contains(models.CsakRendelesre, family) {
		psp.AvailableForOrder = "0"
	}
	return psp
}
