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
	)

	psp := models.PsProduct{}
	features := map[string]string{
		"Anyag":          "", // Acél|Rozsdamentes
		"Felületkezelés": "", // Natúr|Horganyzott
		"Szemforma":      "", // Csomózott|
		"Huzal átmérő":   "",
		"Belső hossz":    "",
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

	// Minden lánc acél
	features["Anyag"] = "Acél"

	// ********************************************
	// `N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)(?:_(H))?$`
	// N-SZL-9-3x16 vagy N-SZL-9-3x16_H
	// Horganyzott vagy natúr szemeslánc
	match = regExpSZL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]

		features["Felületkezelés"] = "Natúr" // Natúr | Horganyzott
		features["Szemforma"] = "N/A"        // Egyenes | Csomózott
		features["Huzal átmérő"] = match[3]
		features["Belső hossz"] = match[4]

		// Ha H-ra végződik, horganyzott
		if match[5] == "H" {
			features["Felületkezelés"] = "Horganyzott"
			imageTag = "SZL_H"
			psp.Tags = "Horganyzott,Szemeslánc," + pStr
			psp.MetaKeywords = pStr + ",horganyzott,szemeslánc"
		} else {
			psp.Tags = "Szemeslánc," + pStr
			psp.MetaKeywords = pStr + ",szemeslánc"
			imageTag = family
		}

		psp.Name = p.Name
		psp.Description = fmt.Sprintf(
			"%s gyártmányú, %s mm huzalátmérőjű, %s mm belső hosszúságú, %s szemformájú, %s felületű, %s szemeslánc.",
			psp.Manufacturer, features["Huzal átmérő"],
			features["Belső hossz"], features["Szemforma"],
			strings.ToLower(features["Felületkezelés"]),
			strings.ToLower(features["Anyag"]),
		)
	}

	// ********************************************
	// Rozsdamentes szemeslánc
	// regExpSSSZL = regexp.MustCompile(`N-(SSSZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)

	match = regExpSSSZL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]

		features["Anyag"] = "Rozsdamentes acél" // Acél
		features["Felületkezelés"] = "Natúr"    // Natúr | Horganyzott
		features["Szemforma"] = "N/A"           // Egyenes | Csomózott
		features["Huzal átmérő"] = match[3]
		features["Belső hossz"] = match[4]

		psp.Name = p.Name
		psp.Description = fmt.Sprintf(
			"%s gyártmányú, %s mm huzalátmérőjű, %s mm belső hosszúságú, %s szemformájú, %s felületű, %s szemeslánc.",
			psp.Manufacturer, features["Huzal átmérő"],
			features["Belső hossz"], features["Szemforma"],
			strings.ToLower(features["Felületkezelés"]),
			strings.ToLower(features["Anyag"]),
		)

		psp.Tags = "rozsdamentes,saválló,Szemeslánc," + pStr
		psp.MetaKeywords = pStr + ",rozsdamentes,saválló,szemeslánc"
		imageTag = family
	}

	// ********************************************
	// Szemeslánc patentszeme
	// regExpSZLPSZ = regexp.MustCompile(`N-(SZLPSZ)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)

	match = regExpSZLPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]

		features["Anyag"] = "Acél"                  // Acél
		features["Felületkezelés"] = "Rozsdamentes" // Natúr | Horganyzott
		features["Szemforma"] = "N/A"               // Egyenes | Csomózott
		features["Huzal átmérő"] = match[3]
		features["Belső hossz"] = match[4]

		psp.Name = p.Name
		psp.Description = fmt.Sprintf(
			"%s gyártmányú, %s mm huzalátmérőjű, %s mm belső hosszúságú, %s szemformájú, %s felületű, %s patentszem.",
			psp.Manufacturer, features["Huzal átmérő"],
			features["Belső hossz"], features["Szemforma"],
			strings.ToLower(features["Felületkezelés"]),
			strings.ToLower(features["Anyag"]),
		)

		psp.Tags = "rozsdamentes,saválló,Szemeslánc patentszem," + pStr
		psp.MetaKeywords = pStr + ",rozsdamentes,saválló,szemeslánc patentszem"
		imageTag = family
	}

	// ********************************************
	// Nem felületkezelt szemes bányalánc patentszem 3 mérettel
	// regExpSZL3 = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
	match = regExpSZL3.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]

		features["Felületkezelés"] = "Natúr" // Natúr | Horganyzott
		features["Szemforma"] = "N/A"        // Egyenes | Csomózott
		features["Huzal átmérő"] = match[3]
		features["Belső hossz"] = match[4]

		psp.Name = p.Name
		psp.Description = fmt.Sprintf(
			"%s gyártmányú, %s mm huzalátmérőjű, %s mm belső hosszúságú, %s szemformájú, %s felületű, %s bányalánc.",
			psp.Manufacturer, features["Huzal átmérő"],
			features["Belső hossz"], features["Szemforma"],
			strings.ToLower(features["Felületkezelés"]),
			strings.ToLower(features["Anyag"]),
		)

		psp.Tags = "bányalánc,patentszem," + pStr
		psp.MetaKeywords = pStr + ",bányalánc patentszem"
		imageTag = family
	}

	// ********************************************
	// További paraméterek beállítása

	// Gyártó beállítása
	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

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

	// Speciális tulajdonságok beállítása
	psp.Features = models.MkFeaturesList(features)

	// Képek előállítása (a Velonál egyedileg készült)
	if psp.ImageURLs == "" && imageTag != "" {
		if features["Szemforma"] != "Csomózott" {
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

	//fmt.Printf("Szl: %s %s %s\n", p.Code, p.Name, psp.ImageURLs)
	//fmt.Printf("Szl: %s: %s %s \n", imageTag, p.Code, p.Name)

	// Ha nincs ára, akkor csak rendelésre
	if psp.UnitPrice == "0" && psp.Quantity == "0.0" {
		psp.ShowPrice = "0"
		psp.AvailableForOrder = "1"
		fmt.Printf("Z-0-0: %-20s %8s Ft. %6s m\n", p.Code, psp.UnitPrice, psp.Quantity)
	}

	// Ha nincs ára, de van belőle, rendelhető
	if psp.UnitPrice == "0" && psp.Quantity != "0.0" {
		fmt.Printf("Z-0-1: %-20s %8s Ft. %6s m\n", p.Code, psp.UnitPrice, psp.Quantity)
	}

	return psp

}

/*

Ezeket be lehetne árazni:
Z-0-1: N-SZLPSZ-0-8x24             0 Ft.   11.0 m
Z-0-1: N-SZLPSZ-0-9x27             0 Ft.   14.0 m
Z-0-1: N-SZL-0-1,8x25_H            0 Ft.    5.0 m
Z-0-1: N-SZLPSZ-0-10x28            0 Ft.    6.0 m
Z-0-1: N-SZLPSZ-0-11x31            0 Ft.    2.0 m
Z-0-1: N-SZL-0-3x26_H              0 Ft.   23.0 m
Z-0-1: N-SZL-0-3x16_H              0 Ft.    0.5 m
Z-0-1: N-SZL-0-3,1x41_H            0 Ft.   39.0 m
Z-0-1: N-SZL-0-4x32_H              0 Ft.   31.2 m
Z-0-1: N-SZL-0-4x19_H              0 Ft.   42.6 m
Z-0-1: N-SZL-9-5x18,5              0 Ft.   53.0 m
Z-0-1: N-SZL-0-5x35_H              0 Ft.   34.0 m
Z-0-1: N-SZL-0-6x42_H              0 Ft.    4.0 m
Z-0-1: N-SZL-0-8x52_H              0 Ft.   20.2 m
Z-0-1: N-SZL-11-8x31               0 Ft.  238.0 m
Z-0-1: N-SZL-11-11x31              0 Ft. 2953.5 m
Z-0-1: N-SZL-11-13x82              0 Ft.   98.5 m

{if $product.quantity <= 0}
        {if !$product.allow_oosp}
            <li class="product-flag call-to-order">
                <i class="fa fa-phone"></i> Hívjon minket!
            </li>
        {else}
            <li class="product-flag order-on-demand">
                <i class="fa fa-shopping-cart"></i> Csak rendelésre!
            </li>
        {/if}
    {/if}

*/
