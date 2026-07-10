package gl

//*******************************************************************
// Görgősláncok
//*******************************************************************

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

func ProcessGl(p models.KsProduct, prodCodes *[]string) models.PsProduct {
	var (
		match          []string
		family         string
		manufacturerId string
		sorokszama     string
		productType    string
		pStr           string
	)
	psp := models.PsProduct{}
	features := map[string]string{
		"GlFamily":              "", // 08B1 és társai a kereséshez
		"Anyag":                 "", // Acél |
		"Erősített":             "", // Igen|Nem
		"Osztás":                "", // Típus: Agyas lánckerék | Laplánckerék
		"Belső hevedertávolság": "", // Fogedzett: Igen | Nem
		"Csapátmérő":            "", // Fogak száma
		"Görgőátmérő":           "", // mm
		"Csaphossz":             "", // mm
		"Csaptípus":             "", // Tömör|Csőcsap
		"Szemforma":             "", // Piskóta|Egyenes|Hajlított
		"Kivitel":               "", // Egysoros | kétsoros | hármosoros (sorokszama alapján képezve)
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
	psp.Categories = "Láncok/Görgősláncok"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// Görgőslánc: N-GL-5-24B3
	// N-(GL)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[productType+sorokszama]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Standard Görgőslánc", productType, sorokszama)
		psp.Tags = "Standard,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,standard,görgőslánc"
		//fmt.Printf("%s: (%s) -> %s\n", p.Code, productType, features["Csaphossz"])

	}

	// Rozsdamentes görgőslánc
	// N-(SSGL)-([0-9]+)-([0-9ABC]+)([0-9])$
	match = regExpSSGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Rozsdamentes acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Görgőslánc", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,görgőslánc"

	}

	// Heavy görgőslánc
	// N-(GL)-([0-9]+)-([0-9ABC]+)([123])_H$
	match = regExpGL_H.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Igen"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Erősített Görgőslánc", productType, sorokszama)
		psp.Tags = "Erősített,Heavy,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,erősített,heavy,görgőslánc"
	}

	// Mofa görgőslánc
	// N-(GL)-([0-9]+)-([0-9ABC]+)([123])_MOFA(_[0-9]+)?$
	match = regExpGLMOFA.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s MOFA Görgőslánc", productType, sorokszama)
		psp.Tags = "MOFA,Görgőslánc"
		psp.MetaKeywords = "Mofa,görgőslánc"
	}

	// Velo görgőslánc
	// N-(GL)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$
	match = regExpGLVELO.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = "1" // Velo mindig egysoros
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Velo Görgőslánc", productType, sorokszama)
		psp.Tags = "Standard,Görgőslánc"
		psp.MetaKeywords = "Velo,görgőslánc"
	}

	// Csőcsapos
	// N-(CSCSGL)-([0-9]+)-([0-9ABC]+)1$
	match = regExpCSCSGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = "1" // Mindig egysoros
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Csőcsapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Csőcsapos Standard Görgőslánc", productType, sorokszama)
		psp.Tags = "Csőcsapos,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,csőcsapos,görgőslánc"
	}

	// Párhuzamos profilú
	// N-(PPGL)-([0-9]+)-([0-9ABC]+)([1-3])$
	match = regExpPPGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Párhuzamos profilú"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Párhuzamos Profilú Standard Görgőslánc", productType, sorokszama)
		psp.Tags = "Standard,Párhuzamos profilú,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,párhuzamos profilú,standard,görgőslánc"
	}

	// Típus beállítása a kereséshez
	if pStr != "" {
		features["GlFamily"] = pStr
	}
	// Gyártó beállítása
	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

	kemenysegTmp := "Standard"
	if features["Erősített"] == "Igen" {
		kemenysegTmp = "Erősített"
	}
	psp.Description = fmt.Sprintf("%s , %s soros, %s mm osztású, %s mm belső hevedertávolságú, %s mm görgőátmérőjű %s %s görgőslánc.", psp.Manufacturer, sorokszama, features["Osztás"], features["Belső hevedertávolság"], features["Görgőátmérő"], strings.ToLower(kemenysegTmp), strings.ToLower(features["Anyag"]))
	// Ebben a fázisában kell beállítani és nem lehet pont a végén.

	psp.Summary = psp.Description
	qtyTmp, _ := strconv.ParseFloat(psp.Quantity, 64)
	if qtyTmp == 0 {
		// Lehet rendelni
		psp.OutOfStockAction = "2"
		psp.Summary += models.JelenlegNemElerheto
		psp.Summary += "<hr>A szállítási határidőről egyeztessen munkatársunkkal!"
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

	psp.MetaDescription = strings.TrimRight(psp.Description, ".")
	psp.MetaTitle = psp.Name
	psp.URLRewritten = p.Code

	// Képek előállítása
	if psp.ImageURLs == "" {
		psp.ImageURLs = fmt.Sprintf(
			"%s/N-%s-%s.png,%s/D-%s-%s.png",
			models.ImagesBase, family, sorokszama,
			models.ImagesBase, family, sorokszama)
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
