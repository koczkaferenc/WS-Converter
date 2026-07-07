// *******************************************************************
// Szenmesláncok
// *******************************************************************
package szl

// Szemeslánc
import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

func ProcessSzl(p models.KsProduct, prodCodes *[]string) models.PsProduct {

	var (
		match          []string
		family         string
		manufacturerId string
		pStr           string
		imageTag       string
		//kulsoMagassag  string

	)

	psp := models.PsProduct{}
	features := map[string]string{
		"Anyag":          "", // Acél|Rozsdamentes
		"Feluletkezeles": "", // Natúr|Horganyzott
		"Szemforma":      "", // Csomózott|
		"HuzalAtmero":    "",
		"BelsoHossz":     "",
		"KulsoMagassag":  "",
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
	psp.Categories = "Láncok/Szemesláncok"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`
	// N-SZL-9-3x16
	// Nem felületkezelt szemeslánc
	match = regExpSZL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]

		features["Anyag"] = "Acél"           // Acél | Rozsdamentes
		features["Feluletkezeles"] = "Natúr" // Natúr | Horganyzott
		features["Szemforma"] = "Csomózott"  // Egyenes | Csomózott
		features["HuzalAtmero"] = match[3]
		features["BelsoHossz"] = match[4]
		features["KulsoMagassag"] = ""

		psp.Name = p.Name
		psp.Description = fmt.Sprintf("%s %s %s szemeslánc", psp.Manufacturer, features["Szemforma"], p.Name)

		psp.Tags = "Szemeslánc," + pStr
		psp.MetaKeywords = pStr + ",szemeslánc"
		imageTag = "SZL"

	}

	// Gyártó beállítása
	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

	// psp.Description = fmt.Sprintf(
	// 	"%s gyártmányú %s, %s mm osztású, %s mm belső hevedertávolságú, %s mm görgőátmérőjű %s szemformájú %s %s %s patentszem.",
	// 	psp.Manufacturer, strings.ToLower(models.Sornevek[sorokszama]),
	// 	features["Osztás"], features["Belső hevedertávolság"], features["Görgőátmérő"],
	// 	strings.ToLower(features["Szemforma"]),
	// 	strings.ToLower(features["Csaptípus"]), strings.ToLower(kemenysegTmp),
	// 	strings.ToLower(features["Anyag"]))

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

	// Speciális tulajdonságok beállítása
	psp.Features = models.MkFeaturesList(features)

	// Képek előállítása (a Velonál egyedileg készült)
	if psp.ImageURLs == "" && imageTag != "" {
		if features["Szemforma"] == "Csomózott" {
			psp.ImageURLs = fmt.Sprintf(
				"%s/N-%s.png,%s/D-%s.png",
				models.ImagesBase, imageTag,
				models.ImagesBase, imageTag)
		} else {
			psp.ImageURLs = fmt.Sprintf(
				"%s/N-%s-csomozott.png,%s/D-%s-csomozott.png",
				models.ImagesBase, imageTag,
				models.ImagesBase, imageTag)
		}
		psp.ImageAltTexts = psp.Name
	}

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

	fmt.Printf("Szl: %s %s %s\n", p.Code, p.Name, psp.ImageURLs)

	return psp

}
