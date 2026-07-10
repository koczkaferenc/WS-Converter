package gl

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

// *******************************************************************
// Görgősláncok patentszemei
// *******************************************************************
func ProcessGlPsz(p models.KsProduct, prodCodes *[]string) models.PsProduct {
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
		"Csaptípus":             "", // Tömör|Csőcsapos
		"Szemforma":             "", // Piskóta|Egyenes|Hajlított
		"Kivitel":               "", // Egysoros | kétsoros | hármosoros (sorokszama alapján képezve)
		"Rögzítés":              "", // Rugós lemezes | Sasszeges
	}

	// Általános tulajdonságok
	// w.SKU = p.Code
	// w.Anyag = "Acél"
	// w.Kivitel = "Normál"    // Normál | Heavy
	// w.Rogzites = "Lemezes"  // Lemezes | Sasszeges
	// w.Csaptipus = "Tömör csapos"   // Tömör | Csőcsapos
	// w.Szemforma = "Piskóta" // Piskóta | Párhuzamos Profilú | Hajlított
	// w.WeightClass = "kg."
	// w.Category = "Patentszemek"
	// w.ClassId = "Görgőslánc"
	// w.Quantity = fmt.Sprintf("%.1f", p.Stock)
	// w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
	// w.TaxClass = "27%"
	// w.QuantityUnit = p.Unit
	// w.Weight = fmt.Sprintf("%.1f", p.Weight)

	// Ha nincs belőle raktáron, nem elérhető.
	// qty, _ := strconv.ParseFloat(w.Quantity, 64)
	// if qty == 0 {
	// 	w.ShortDescription += models.JelenlegNemElerheto
	// } else {
	// 	if slices.Contains(models.CsakRendelesre, family) {
	// 		w.ShortDescription += models.CsakRendelesreLeiras
	// 	}
	// }
	// w.ShortDescription += models.Zaradek

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
	psp.TextInStock = "db"          // méter
	psp.TextBackorderAllowed = "db" // méter
	psp.Categories = "Láncok/Patentszemek"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// Normál patetszem
	// N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGLPSZ.FindStringSubmatch(p.Code)
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
		features["Rögzítés"] = "Normál"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Standard Patentszem", productType, sorokszama)
		psp.Tags = "Standard,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,standard,patentszem"
	}

	// Heavy patentszem
	// N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])_H$
	match = regExpGLPSZ_H.FindStringSubmatch(p.Code)
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
		features["Rögzítés"] = "Normál"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Erősített Patentszem", productType, sorokszama)
		psp.Tags = "Erősített,Heavy,patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,erősített,heavy,patentszem"
	}

	// Sasszeges, normál patentszem
	// N-(GLSPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGLSPSZ.FindStringSubmatch(p.Code)
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
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Acél Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Sasszeg,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeg,patentszem"
	}

	// Rozsdamentes görgőslánc patentszem rugós lemezes
	// N-(SSGLPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpSSGLPSZ.FindStringSubmatch(p.Code)
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
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,patentszem"
	}

	// Rozsdamentes görgőslánc patentszem sasszeges
	// N-(SSGLSPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpSSGLSPSZ.FindStringSubmatch(p.Code)
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
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Sasszeg,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeg,patentszem"
	}

	// Sasszeges hajlított patentszem (a hajlított mindig sasszeges)
	// N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGLHOK.FindStringSubmatch(p.Code)
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
		features["Szemforma"] = "Hajlított"
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Hajlított Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Sasszeg,Hajlított,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeg,hajlított,patentszem"
	}

	// Erősített görgőslánc sasszeges hajlított patentszem
	// N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])_H$
	match = regExpGLHOK_H.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = match[3]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Igen"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Hajlított"
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Erősített Hajlított Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Erősített,Heavy,Sasszeg,Hajlított,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,erősített,heavy,sasszeges,hajlított,patentszem"
	}

	// Rozsdamentes görgősláncok rugós sasszeges hajlított patentszeme
	// N-(SSGLHOK)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpSSGLHOK.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = match[3]
		pStr = productType + sorokszama

		features["Anyag"] = "Rozsdamentes Acél"
		features["Erősített"] = "Igen"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Sasszeges"
		features["Szemforma"] = "Hajlított"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Hajlított Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Hajlított,Sasszeg,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeges,hajlított,sasszeg,patentszem"
	}

	// GLPSZ MOFA
	// N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])_MOFA$
	match = regExpGLPSZMOFA.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = fmt.Sprintf("%s%s_MOFA", match[3], sorokszama)

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s MOFA Patentszem", productType, sorokszama)
		psp.Tags = "MOFA,Patentszem"
		psp.MetaKeywords = "Mofa,Patentszem"
	}

	// Hajlított MOFA
	// N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])_MOFA$
	match = regExpGLHOKMOFA.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = fmt.Sprintf("%s%s_MOFA", match[3], sorokszama)

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Hajlított"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s MOFA Patentszem", productType, sorokszama)
		psp.Tags = "MOFA,Patentszem,Összekötő szem"
		psp.MetaKeywords = "Mofa,Patentszem,összekötő szem"
	}

	// VELO Patentszem
	// N-(GLPSZ)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$
	match = regExpGLPSZVELO.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Velo mindig egysoros
		productType = fmt.Sprintf("%s", match[3])

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s VELO Patentszem", productType, sorokszama)
		psp.Tags = "VELO,Patentszem"
		psp.MetaKeywords = "Velo,Patentszem"

		// Egyedi képei vannak
		psp.ImageURLs = fmt.Sprintf(
			"%s/N-GLPSZ-%s.png,%s/D-GLPSZ-%s.png",
			models.ImagesBase, sorokszama,
			models.ImagesBase, sorokszama)
		psp.ImageAltTexts = psp.Name

		// w.Image = fmt.Sprintf("product/N-GLPSZ-%s.png", sorokszama)
		// w.ImageAdditional = fmt.Sprintf("product/D-GLPSZ-%s.png", sorokszama)
	}

	// Hajlított VELO
	// N-(GLHOK)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$
	match = regExpGLHOKVELO.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Velo mindig egysoros
		productType = fmt.Sprintf("%s", match[3])

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Hajlított"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s VELO Patentszem", productType, sorokszama)
		psp.Tags = "VELO,Patentszem,Összekötő szem"
		psp.MetaKeywords = "Velo,Patentszem,Összekötő szem"
	}

	// Csőcsapos egyenes patentszem
	// N-(CSCSGLPSZ)-([0-9]+)-([0-9ABC]+)1$
	match = regExpCSCSGLPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Mindig egysoros
		productType = match[3]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[productType+sorokszama]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Csőcsapos"
		features["Szemforma"] = "Egyenes"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Csőcsapos Egyenes Patentszem", productType, sorokszama)
		psp.Tags = "Csőcsap,Egyenes,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,egyenes,csőcsap,patentszem"
	}

	// Párhuzamos profilú patentszem
	// regExpPPGLPSZ = regexp.MustCompile(`N-(PPGLPSZ)-([0-9])+-([0-9ABC]+)1$`)
	match = regExpPPGLPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Mindig egysoros
		productType = match[3]

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[productType+sorokszama]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Csőcsapos"
		features["Szemforma"] = "Párhuzamos profilú"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Standard Patentszem", productType, sorokszama)
		psp.Tags = "Standard,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,standard,patentszem"
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

	psp.Description = fmt.Sprintf(
		"%s gyártmányú %s, %s mm osztású, %s mm belső hevedertávolságú, %s mm görgőátmérőjű %s szemformájú %s %s %s patentszem.",
		psp.Manufacturer, strings.ToLower(models.Sornevek[sorokszama]),
		features["Osztás"], features["Belső hevedertávolság"], features["Görgőátmérő"],
		strings.ToLower(features["Szemforma"]),
		strings.ToLower(features["Csaptípus"]), strings.ToLower(kemenysegTmp),
		strings.ToLower(features["Anyag"]))

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
	// Ebben a fázisában kell beállítani és nem lehet pont a végén.
	psp.MetaDescription = strings.TrimRight(psp.Description, ".")
	psp.MetaTitle = psp.Name
	psp.URLRewritten = p.Code

	// Képek előállítása (a Velonál egyedileg készült)
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
