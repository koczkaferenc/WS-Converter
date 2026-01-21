package flyer

// Flyer láncok
import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"ws-updater/models"
)

var osztasTab = []string{4: "12,70", 5: "15,88", 6: "19,05", 8: "25,40", 10: "31,75", 12: "38,10", 14: "44,45", 16: "50,80"}

func ProcessFlyer(p models.KsProduct) models.WsProduct {
	var (
		regExpFL       = regexp.MustCompile(`N-(FL)-([0-9]+)-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$`)
		regExpFLCS     = regexp.MustCompile(`N-(FLCS)-([0-9]+)-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$`)
		match          []string
		family         string
		nameTag        string
		osztasTag      int
		kivitelkod     string
		manufacturerId int
	)

	w := models.WsProduct{}

	// KS-ek fogszáma és keménysége
	// N-KS-0-08B2_Z30
	// N-KS-0-08B2_Z21_G
	w.SKU = p.Code
	w.Quantity = fmt.Sprintf("%.1f", p.Stock)
	w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
	w.TaxClass = "27%"
	w.QuantityUnit = p.Unit
	w.Weight = fmt.Sprintf("%.1f", p.Weight)
	w.WeightClass = "kg."

	w.Anyag = "Acél"
	w.Category = "Flyer láncok"
	w.ClassId = "Flyer lánc"

	// N-FL-7-LL1044
	// N-(FL)-([0-9]+)-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$
	//    1      2      3            4       5      6
	//  fam    Manuf    --LL/BL--   osztas  -hevederek-
	match = regExpFL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		w.Manufacturer = models.Manufacturers[manufacturerId]
		kivitelkod = match[3] // LL | BL
		nameTag = fmt.Sprintf("%s %s%s%s", kivitelkod, match[4], match[5], match[6])
		osztasTag, _ = strconv.Atoi(match[4])
		w.Osztas = osztasTab[osztasTag]
		w.HevederSzam = fmt.Sprintf("%sx%s", match[5], match[6])
		w.Name = fmt.Sprintf("%s Flyer Lánc", nameTag)
		w.ShortDescription = fmt.Sprintf("%s %s %s flyer lánc", w.Manufacturer, kivitelkod, w.HevederSzam)
	}

	// `N-(FLCS)-([0-9]+)-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$`
	//     1       2        3           4       5      6
	//    fam     Manuf    --LL/BL--    osztas  -hevederek-
	// N-FLCS-7-BL646
	match = regExpFLCS.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		w.Manufacturer = models.Manufacturers[manufacturerId]
		kivitelkod = match[3] // LL | BL
		nameTag = fmt.Sprintf("%s %s%s%s", kivitelkod, match[4], match[5], match[6])
		osztasTag, _ = strconv.Atoi(match[4])
		w.Osztas = osztasTab[osztasTag]
		w.HevederSzam = fmt.Sprintf("%sx%s", match[5], match[6])
		w.Name = fmt.Sprintf("%s Flyer Lánc Csap", nameTag)
		w.ShortDescription = fmt.Sprintf("%s %s %s flyer lánc csap", w.Manufacturer, kivitelkod, w.HevederSzam)
	}

	// Ha nincs belőle raktáron, nem elérhető.
	qty, _ := strconv.ParseFloat(w.Quantity, 64)
	if qty == 0 {
		w.ShortDescription += models.JelenlegNemElerheto
	} else {
		if slices.Contains(models.CsakRendelesre, family) {
			w.ShortDescription += models.CsakRendelesreLeiras
		}
	}
	w.ShortDescription += models.Zaradek

	if w.Image == "" {
		w.Image = fmt.Sprintf("product/N-%s.png", family)
	}
	if w.ImageAdditional == "" {
		w.ImageAdditional = fmt.Sprintf("product/D-%s.png", family)
	}

	return w
}
