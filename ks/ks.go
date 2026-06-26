package ks

// Lánckerék szelet
import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"ws-updater/models"
)

func ProcessKs1(p models.KsProduct) models.PsProduct {
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
		productType    string
		manufacturerId string
		sorokszama     string
	)

	psp := models.PsProduct{}
	// Anyag: Acél|Rozsdamentes

	features := map[string]string{
		"Anyag":     "N/A", // Acél |
		"Típus":     "N/A", // Típus: Agyas lánckerék | Laplánckerék
		"Fogedzett": "N/A", // Fogedzett: Igen | Nem
		"Fogszám":   "N/A", // Fogak száma
		"Kivitel":   "N/A", // Egysoros | kétsoros | hármosoros (sorokszama alapján képezve)
	}

	// Alap paraméterek feltöltése a KulcsSoftból
	// KS-ek fogszáma és keménysége
	// N-KS-0-08B2_Z30
	// N-KS-0-08B2_Z21_G

	// &p.Weight,
	// &p.Stock,
	// &p.Unit//
	// &p.ReSellerPrice,

	psp.DeliveryTimeInStock = "5 nap"     // TODO
	psp.DeliveryTimeOutOfStock = "14 nap" // TODO

	// TODO: mi a mértékegység? w.WeightClass = "kg."
	// fmt.Printf(p.Unit)

	psp.ID = ""            // Az ID üres, cikkszámokkal dogozunk
	psp.Reference = p.Code // Ez a cikkszám, ID helyett használjuk
	psp.Active = "1"
	psp.Name = "terméknév" //p.Name
	psp.PriceTaxExcluded = fmt.Sprintf("%.0f", p.WebPrice)
	psp.UnitPrice = fmt.Sprintf("%.0f", p.WebPrice) // TODO
	psp.TaxRulesID = "1"                            // ÁFA kulcs 27%
	psp.Quantity = fmt.Sprintf("%.1f", p.Stock)     // Mennyiség
	psp.Weight = fmt.Sprintf("%.1f", p.Weight)
	psp.Unity = "db"

	psp.Categories = "Lánckerekek"
	psp.OnSale = "0"          // TODO: akció számítása
	psp.DiscountAmount = ""   // TODO
	psp.DiscountPercent = "0" // TODO
	psp.DiscountFrom = ""     // TODO
	psp.DiscountTo = ""       // TODO

	psp.SupplierReference = "" // TODO
	psp.Supplier = ""          // TODO
	psp.Manufacturer = ""      // TODO
	psp.EAN13 = ""             // TODO
	psp.UPC = ""               // TODO
	psp.MPN = ""               // TODO
	psp.Ecotax = ""            // TODO
	psp.Width = ""             // TODO
	psp.Height = ""            // TODO
	psp.Depth = ""             // TODO

	// nem használjuk
	psp.MinimalQuantity = ""        // TODO
	psp.LowStockLevel = ""          // TODO
	psp.LowStockAlertEmail = ""     // TODO
	psp.Visibility = ""             // TODO
	psp.AdditionalShippingCost = "" // TODO

	psp.TextInStock = "Raktáron"                                   // TODO
	psp.TextBackorderAllowed = "Nincs raktáron, de utánrendelhető" // TODO
	psp.AvailableForOrder = "1"                                    // Rendelhető?
	psp.ProductAvailableDate = ""                                  // TODO
	psp.ProductCreationDate = ""                                   // TODO
	psp.ShowPrice = "1"                                            // TODO
	psp.DeleteExistingImages = "1"                                 // TODO

	psp.AvailableOnlineOnly = "0"     // TODO
	psp.Condition = ""                // Mindenünk új
	psp.Customizable = ""             // Nem konfigurálható
	psp.UploadableFiles = ""          // Nincs ilyenünk
	psp.TextFields = ""               // TODO
	psp.OutOfStockAction = ""         // TODO
	psp.VirtualProduct = ""           // TODO
	psp.FileURL = ""                  // TODO
	psp.NumberOfAllowedDownloads = "" // TODO
	psp.ExpirationDate = ""           // TODO
	psp.NumberOfDays = ""             // TODO
	psp.ShopIDOrName = ""             // TODO
	psp.AdvancedStockManagement = ""  // TODO
	psp.DependsOnStock = ""           // TODO
	psp.Warehouse = ""                // TODO

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
		psp.Accessories = "N-KR-10-16B2_Z30, "
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

	psp.Summary = fmt.Sprintf("%s gyártmányú %s %s fogszámú %s fogedzett %s.",
		psp.Manufacturer, strings.ToLower(models.Sornevek[sorokszama]), features["Fogszám"], strings.ToLower(features["Fogedzett"]), strings.ToLower(features["Típus"]))
	// Ebben a fázisában, és nem lehet pont a végén.
	psp.MetaDescription = strings.TrimRight(psp.Summary, ".")

	psp.Summary += "<p style='color: red;'>A lánckerék technikai furattal van ellátva. Amennyiben megmunkálva kívánja beszerezni, vegye fel a kapcsolatot a munkatársunkkal a kívánt furat méretének egyeztetése érdekében!</p>"
	psp.Summary += models.Zaradek

	psp.Tags = "Lánckerék"
	psp.MetaTitle = "Lánckerék"
	psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros lánckerék"
	psp.URLRewritten = p.Code

	// Elérhető mennyiség
	psp.Quantity = strconv.Itoa(int(p.Stock))

	// Ha nincs belőle raktáron, nem elérhető.
	// qty, _ := strconv.ParseFloat(w.Quantity, 64)
	// if qty == 0 {
	// 	w.ShortDescription += models.JelenlegNemElerheto
	// } else {
	// 	if slices.Contains(models.CsakRendelesre, family) {
	// 		w.ShortDescription += models.CsakRendelesreLeiras
	// 	}
	// }

	psp.Name = fmt.Sprintf("%s-%s%sZ%s %s",
		family, productType, sorokszama, features["Fogszám"], strings.ToLower(features["Típus"]))

	if features["Típus"] == "Fogedzett" {
		psp.Name += "G"
	}
	psp.ImageURLs = fmt.Sprintf(
		"%s/N-%s-%s.png,%s/D-%s-%s.png",
		models.ImagesBase, family, sorokszama,
		models.ImagesBase, family, sorokszama)
	psp.ImageAltTexts = psp.Name

	// Speciális tulajdonságok
	features["Kivitel"] = models.Sornevek[sorokszama]
	i := 1
	for k, v := range features {
		psp.Features += fmt.Sprintf("%s:%s:%d,", k, v, i)
		i++
	}
	if psp.Features != "" {
		psp.Features = strings.TrimSuffix(psp.Features, ",")
	}
	fmt.Printf("%s: %v\n", psp.Reference, psp.Features)
	return psp
}

// func ProcessKs(p models.KsProduct) models.WsProduct {
// 	var (
// 		regExpKS   = regexp.MustCompile(`N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
// 		regExpKS_G = regexp.MustCompile(`N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
// 		// Agyas lánckerekek
// 		regExpKR   = regexp.MustCompile(`N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
// 		regExpKR_G = regexp.MustCompile(`N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
// 		// Sajnos ezt ebben a formában is felvitték
// 		regExpGKR      = regexp.MustCompile(`N-(GKR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
// 		match          []string
// 		family         string
// 		productType    string
// 		manufacturerId int
// 		numOfRows      string
// 	)
//
// 	w := models.WsProduct{}
//
// 	// KS-ek fogszáma és keménysége
// 	// N-KS-0-08B2_Z30
// 	// N-KS-0-08B2_Z21_G
// 	w.SKU = p.Code
// 	w.Quantity = fmt.Sprintf("%.1f", p.Stock)
// 	w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
// 	w.TaxClass = "27%"
// 	w.QuantityUnit = p.Unit
// 	w.Weight = fmt.Sprintf("%.1f", p.Weight)
// 	w.WeightClass = "kg."
//
// 	w.Anyag = "Acél"
// 	w.Category = "Lánckerekek"
// 	w.ClassId = "Lánckerék"
// 	w.LanckerekTipus = "Laplánckerék" // Agyas lánckerék |  Laplánckerék
// 	w.Fogedzett = "Standard"          // Fogedzett | Standard
//
// 	// N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
// 	// N-KR-0-08B2_Z30
// 	match = regExpKS.FindStringSubmatch(p.Code)
// 	if match != nil {
// 		family = match[1]
// 		manufacturerId, _ = strconv.Atoi(match[2])
// 		productType = match[3]
// 		numOfRows = match[4]
// 		w.Fogszam = match[5]
// 	}
//
// 	// N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$
// 	// N-KS-0-08B2_Z21_G
// 	match = regExpKS_G.FindStringSubmatch(p.Code)
// 	if match != nil {
// 		family = match[1]
// 		manufacturerId, _ = strconv.Atoi(match[2])
// 		productType = match[3]
// 		numOfRows = match[4]
// 		w.Fogszam = match[5]
// 		w.Fogedzett = "Fogedzett"
// 	}
//
// 	// N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
// 	// N-KR-0-08B2_Z30
// 	match = regExpKR.FindStringSubmatch(p.Code)
// 	if match != nil {
// 		family = match[1]
// 		manufacturerId, _ = strconv.Atoi(match[2])
// 		productType = match[3]
// 		numOfRows = match[4]
// 		w.Fogszam = match[5]
// 		w.LanckerekTipus = "Agyas lánckerék"
// 	}
//
// 	// N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$
// 	// N-KR-0-08B2_Z21_G
// 	match = regExpKR_G.FindStringSubmatch(p.Code)
// 	if match != nil {
// 		family = match[1]
// 		manufacturerId, _ = strconv.Atoi(match[2])
// 		productType = match[3]
// 		numOfRows = match[4]
// 		w.Fogszam = match[5]
// 		w.Fogedzett = "Fogedzett"
// 		w.LanckerekTipus = "Agyas lánckerék"
// 	}
//
// 	// N-(GKR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
// 	// N-GKR-0-08B2_Z30
// 	match = regExpGKR.FindStringSubmatch(p.Code)
// 	if match != nil {
// 		family = match[1]
// 		manufacturerId, _ = strconv.Atoi(match[2])
// 		productType = match[3]
// 		numOfRows = match[4]
// 		w.Fogszam = match[5]
// 		w.LanckerekTipus = "Agyas lánckerék"
// 	}
//
// 	w.Manufacturer = models.Manufacturers[manufacturerId]
// 	w.ShortDescription = fmt.Sprintf("%s gyártmányú %ssoros %s fogszámú %s %s. ",
// 		w.Manufacturer, models.Sornevek[numOfRows], w.Fogszam, strings.ToLower(w.Fogedzett), strings.ToLower(w.LanckerekTipus))
//
// 	if w.LanckerekTipus == "Agyas lánckerék" {
// 		w.ShortDescription += "<p style='color: red;'>Amennyiben a lánckereket megmunkálva kívánja beszerezni, vegye fel a kapcsolatot munkatársunkkal!</p>"
// 	}
//
// 	// Ha nincs belőle raktáron, nem elérhető.
// 	qty, _ := strconv.ParseFloat(w.Quantity, 64)
// 	if qty == 0 {
// 		w.ShortDescription += models.JelenlegNemElerheto
// 	} else {
// 		if slices.Contains(models.CsakRendelesre, family) {
// 			w.ShortDescription += models.CsakRendelesreLeiras
// 		}
// 	}
// 	w.ShortDescription += models.Zaradek
//
// 	if w.Name == "" {
// 		w.Name = fmt.Sprintf("%s %s%sZ=%s %s", family, productType, numOfRows, w.Fogszam, strings.ToLower(w.LanckerekTipus))
// 	}
// 	if w.Fogedzett == "Fogedzett" {
// 		w.Name += ` "G"`
// 	}
// 	if w.Image == "" {
// 		w.Image = fmt.Sprintf("product/N-%s-%s.png", family, numOfRows)
// 	}
// 	if w.ImageAdditional == "" {
// 		w.ImageAdditional = fmt.Sprintf("product/D-%s-%s.png", family, numOfRows)
// 	}
//
// 	return w
// }
//
