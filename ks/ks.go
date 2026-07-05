package ks

// Lánckerekek

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

func ProcessKs(p models.KsProduct, prodCodes *[]string) models.PsProduct {
	var (
		regExpKS   = regexp.MustCompile(`N-(KS)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
		regExpKS_G = regexp.MustCompile(`N-(KS)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
		// Agyas lánckerekek
		regExpKR   = regexp.MustCompile(`N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
		regExpKR_G = regexp.MustCompile(`N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
		// Sajnos ezt ebben a formában is felvitték
		regExpGKR      = regexp.MustCompile(`N-(GKR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
		match          []string
		family         string
		manufacturerId string
		sorokszama     string
		productType    string
	)

	psp := models.PsProduct{}
	features := map[string]string{
		"Anyag":     "N/A", // Acél |
		"Típus":     "N/A", // Típus: Agyas lánckerék | Laplánckerék
		"Fogedzett": "N/A", // Fogedzett: Igen | Nem
		"Fogszám":   "N/A", // Fogak száma
		"Kivitel":   "N/A", // Egysoros | kétsoros | hármosoros (sorokszama alapján képezve)
	}

	psp.DeliveryTimeInStock = "5 nap"     // TODO
	psp.DeliveryTimeOutOfStock = "14 nap" // TODO

	// TODO: mi a mértékegység? w.WeightClass = "kg."
	// fmt.Printf(p.Unit)

	psp.ID = ""            // Az ID üres, cikkszámokkal dogozunk
	psp.Reference = p.Code // Ez a cikkszám, ID helyett használjuk
	psp.Active = "1"
	psp.PriceTaxExcluded = fmt.Sprintf("%.0f", p.WebPrice)
	psp.UnitPrice = fmt.Sprintf("%.0f", p.WebPrice) // TODO
	psp.TaxRulesID = "1"                            // ÁFA kulcs 27%
	psp.Quantity = fmt.Sprintf("%.1f", p.Stock)     // Mennyiség
	psp.AvailableForOrder = "1"
	psp.Weight = fmt.Sprintf("%.1f", p.Weight)
	psp.Unity = "db"
	psp.TextInStock = "db"          // darab
	psp.TextBackorderAllowed = "db" // darab
	psp.Categories = "Lánckerekek"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// ---------------

	// N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
	// N-KR-0-08B2_Z30
	match = regExpKS.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		if ManufacturerId, err := strconv.Atoi(match[2]); err == nil {
			psp.Manufacturer = models.Manufacturers[ManufacturerId]
		}
		productType = match[3]
		features["Anyag"] = "Acél"
		sorokszama = match[4]
		features["Fogszám"] = match[5]
		features["Fogedzett"] = "Nem"
		features["Típus"] = "Laplánckerék"
	}

	// N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$
	// N-KS-0-08B2_Z21_G
	match = regExpKS_G.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		features["Anyag"] = "Acél"
		sorokszama = match[4]
		features["Fogszám"] = match[5]
		features["Fogedzett"] = "Igen"
		features["Típus"] = "Laplánckerék"
	}

	// N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
	// N-KR-0-08B2_Z30
	match = regExpKR.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		features["Anyag"] = "Acél"
		sorokszama = match[4]
		features["Fogszám"] = match[5]
		features["Fogedzett"] = "Nem"
		features["Típus"] = "Agyas lánckerék"
	}

	// N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$
	// N-KR-0-08B2_Z21_G
	match = regExpKR_G.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		features["Anyag"] = "Acél"
		sorokszama = match[4]
		features["Fogszám"] = match[5]
		features["Fogedzett"] = "Igen"
		features["Típus"] = "Agyas lánckerék"
	}

	// N-(GKR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
	// N-GKR-0-08B2_Z30
	match = regExpGKR.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		features["Anyag"] = "Acél"
		sorokszama = match[4]
		features["Fogszám"] = match[5]
		features["Fogedzett"] = "Nem"
		features["Típus"] = "Agyas lánckerék"
	}

	// Gyártó beállítása
	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

	psp.Description = fmt.Sprintf("%s gyártmányú %s %s fogszámú %s fogedzett %s.",
		psp.Manufacturer, strings.ToLower(models.Sornevek[sorokszama]), features["Fogszám"], strings.ToLower(features["Fogedzett"]), strings.ToLower(features["Típus"]))

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

	psp.Tags = "Lánckerék"
	psp.MetaTitle = "Lánckerék"
	psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros lánckerék"
	psp.URLRewritten = p.Code

	// Elérhető mennyiség
	psp.Quantity = strconv.Itoa(int(p.Stock))

	// Termék név előállítása
	psp.Name = fmt.Sprintf("%s-%s%sZ%s %s",
		family, productType, sorokszama, features["Fogszám"], strings.ToLower(features["Típus"]))
	if features["Fogedzett"] == "Fogedzett" {
		psp.Name += "G"
	}

	// A termékkép előállítása
	if psp.ImageURLs == "" {
		if features["Fogedzett"] == "Nem" {
			psp.ImageURLs = fmt.Sprintf(
				"%s/N-%s-%s.png,%s/D-%s-%s.png",
				models.ImagesBase, family, sorokszama,
				models.ImagesBase, family, sorokszama)
		} else {
			psp.ImageURLs = fmt.Sprintf(
				"%s/N-%s-%s_G.png,%s/D-%s-%s_G.png",
				models.ImagesBase, family, sorokszama,
				models.ImagesBase, family, sorokszama)
		}
		psp.ImageAltTexts = psp.Name
	}

	// Speciális tulajdonságok beállítása
	features["Kivitel"] = models.Sornevek[sorokszama]
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

	//fmt.Printf("%s: %v\n", psp.Reference, psp.AvailableForOrder)

	return psp
}
