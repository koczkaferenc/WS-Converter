package mgbf

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

const KiarusitasSzazalek = "5"

func ProcessMgbf(p models.KsProduct, prodCodes *[]string) models.PsProduct {
	var (
		regExpMGBF = regexp.MustCompile(`N-(MGBF)([EHK])-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_(LL|LÉ)$`)

		match          []string
		family         string
		manufacturerId string
		formakod       string
		keresztmetszet string
		hossz          string
	)

	psp := models.PsProduct{}
	features := map[string]string{
		"Anyag":          "", // Kovácsolt vas
		"Forma":          "", // Hajlított|Egyenes|Kanalas
		"Keresztmetszet": "", // pl. 40x40-es
		"Hossz":          "", // Hossz mm-ben
		"Csavarmenet":    "", // Csavarmenet méret
		"Kivitel":        "", // Lap-lap|Lap-él
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
	psp.Unity = "darab"
	psp.TextInStock = "darab"          // darab
	psp.TextBackorderAllowed = "darab" // méter
	psp.Categories = "Boronafogak"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// N-MGBFE-0-16x16x190_M12_LÉ
	// N- MGBF       E                         -0      -16x16x190       _M12_   LL
	//    Boronafog  Egyenes/Hallított/Kanalas Gyarto  szél-szél-hossz  Menet   Lap-lappal/lap-Éllel
	//
	// #// N-(MGBF)(EHK)-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_([A-ZÉ]+)$
	// #// Termék kód és példa
	// #    // 1: family
	// #    // 2: E-Egyenes, K-Kanalas, H-Hajlított
	// #    // 3: Gyártó kód
	// #    // 4: Szélesség
	// #    // 5: Magasság
	// #    // 6: Hossz
	// #    // 7: Csavar menet méret
	// #    // 8: LL: lap-lappal, LE: Lap éllel

	match = regExpMGBF.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[3]

		formakod = match[2]
		switch formakod {
		case "E":
			features["Forma"] = "Egyenes"
		case "H":
			features["Forma"] = "Hajlított"
		case "K":
			features["Forma"] = "Kanalas"
		}

		keresztmetszet = fmt.Sprintf("%sx%s", match[4], match[5])
		features["Keresztmetszet"] = fmt.Sprintf("%s mm", keresztmetszet)
		hossz = fmt.Sprintf("%s", match[6])
		features["Hossz"] = fmt.Sprintf("%s mm", hossz)
		features["Csavarmenet"] = fmt.Sprintf("M%s", match[7])

		switch match[8] {
		case "LL":
			features["Kivitel"] = "Párhuzamos"
		case "LÉ":
			features["Kivitel"] = "Elforgatott"
		}
	}

	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

	psp.Name = fmt.Sprintf("%sx%s-%s boronafog",
		keresztmetszet, hossz, features["Csavarmenet"])
	psp.Tags = fmt.Sprintf("%s boronafog, %s boronafog, ", features["Forma"], features["Kivitel"])
	psp.MetaKeywords = fmt.Sprintf("%s boronafog, %s boronafog, ", features["Forma"], strings.ToLower(features["Kivitel"]))

	psp.Description = fmt.Sprintf("%s mm keresztmetszetű, %s mm hosszú, %s formájú %s állású boronafog %s csavarmenettel.",
		keresztmetszet, hossz, strings.ToLower(features["Forma"]), strings.ToLower(features["Kivitel"]), features["Csavarmenet"])

	psp.Summary = psp.Description
	qtyTmp, _ := strconv.ParseFloat(psp.Quantity, 64)
	if qtyTmp == 0 {
		// Lehet rendelni
		psp.OutOfStockAction = "2"
		psp.Summary += models.JelenlegNemElerheto
		psp.Summary += "<hr>Szállítási idő: kb. 14 nap."
	} else {
		// Ha van raktáron, az beragadt, ezért 5%-os engedménnyel akciózzuk.
		psp.OnSale = "1"
		psp.DiscountPercent = models.KiarusitasSzazalek
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

	if psp.ImageURLs == "" {
		psp.ImageURLs = fmt.Sprintf("%s/N-%s%s.png", models.ImagesBase, family, formakod)
		psp.ImageAltTexts = psp.Name
	}

	// Speciális tulajdonságok beállítása
	psp.Features = models.MkFeaturesList(features)
	// TODO Ideiglenesen kivesszük, mert ha nem ltezik a termék, nagyon lelassul
	//rgxStr := fmt.Sprintf(`^N-GL-[0-9]+-%s%s.*`, productType, sorokszama)
	//psp.Accessories += models.getRelatedProductIds(rgxStr, prodCodes)

	// Rendelhető?
	if slices.Contains(models.CsakRendelesre, family) {
		psp.AvailableForOrder = "0"
	}

	// A boronafog.hu-hoz egy terméklista:
	//fmt.Printf("%s|%s|%s|%s|%s\n", features["Forma"], hossz, keresztmetszet, features["Csavarmenet"], features["Kivitel"])
	return psp
}
