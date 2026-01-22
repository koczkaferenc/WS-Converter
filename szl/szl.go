//go:build ignore
// +build ignore

package szl

// Szemeslánc
import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"ws-updater/models"
)

func ProcessSzl(p models.KsProduct) models.WsProduct {
	var (
		// Nem felületkezelt szemeslánc
		regExpSZL = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
		// Horganyzott szemeslánc
		//regExpSZLH = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
		// Rozsdamentes szemeslánc
		//regExpSSSZL = regexp.MustCompile(`N-(SSSZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
		// Szemeslánc patentszeme
		//regExpSZLPSZ = regexp.MustCompile(`N-(SZLPSZ)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
		// Nem felületkezelt szemes bányalánc patentszem 3 mérettel
		//regExpSZL3 = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)

		match          []string
		family         string
		productType    string
		manufacturerId int
		kulsoMagassag  string
	)

	w := models.WsProduct{}
	w.SKU = p.Code
	w.Quantity = fmt.Sprintf("%.1f", p.Stock)
	w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
	w.TaxClass = "27%"
	w.QuantityUnit = p.Unit
	w.Weight = fmt.Sprintf("%.1f", p.Weight)
	w.WeightClass = "kg."
	w.Category = "Szemesláncok"
	w.ClassId = "Szemeslánc"

	w.Anyag = "Acél"           // Acél | Rozsdamentes
	w.Feluletkezeles = "Natúr" // Natúr | Horganyzott
	w.Szemforma = "Egyenes"    // Egyenes | Csomózott

	// N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`
	// N-SZL-9-3x16
	match = regExpSZL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		w.HuzalAtmero = match[3]
		w.BelsoHossz = match[4]
		kulsoMagassag = ""

	}

	w.Manufacturer = models.Manufacturers[manufacturerId]
	w.ShortDescription = fmt.Sprintf("%s gyártmányú ", w.Manufacturer)

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

	if w.Name == "" {
		w.Name = fmt.Sprintf("%s ", family)
	}

	if w.Image == "" {
		w.Image = fmt.Sprintf("product/N-%s-%s.png", family, numOfRows)
	}
	if w.ImageAdditional == "" {
		w.ImageAdditional = fmt.Sprintf("product/D-%s-%s.png", family, numOfRows)
	}

	return w
}
