//go:build ignore
// +build ignore

package mggl

// Mezőgazdasági láncok
//
import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"ws-updater/models"
)

func ProcessMggl(p models.KsProduct) models.WsProduct {
	var (
		regExpMGGL     = regexp.MustCompile(`N-(MGGL)([A-Z]+)-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_([A-Z]+)$`)
		match          []string
		family         string
		manufacturerId int
	)

	w := models.WsProduct{}

	w.SKU = p.Code
	w.Anyag = "Acél"
	w.Kivitel = "Normál"    // Normál | Heavy (erősített)
	w.Csaptipus = "Tömör"   // Tömör | Csőcsapos
	w.Szemforma = "Piskóta" // Piskóta | Párhuzamos profilú | Hajlított
	w.Category = "Görgőslácok"
	w.ClassId = "Görgőslánc"
	w.Quantity = fmt.Sprintf("%.1f", p.Stock)
	w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
	w.TaxClass = "27%"
	w.QuantityUnit = p.Unit
	w.Weight = fmt.Sprintf("%.1f", p.Weight)
	w.WeightClass = "kg."

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

	// Mezőgazdasági láncok
	// N-(MGGL)([A-Z]+)-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_([A-Z]+)$
	// MGGL-12-CA650_F3.04_4tag
	// MGGL-12-p=50,8_216B1_3tag_U
	// MGGL-12-50,8_216BF1_4tag
	// MGGLPSZ-12-p=50,8_216BF21_U
	// MGGLPSZ-12-p=50,8_CA650
	// MGGLPSZ-12-p=50,8_CA650F8
	match = regExpMGGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		w.Manufacturer = models.Manufacturers[manufacturerId]
	}

	w.ShortDescription = fmt.Sprintf("%s gyártmányú ", w.Manufacturer)
	if w.Image == "" {
		w.Image = fmt.Sprintf("product/N-%s.png", family)
	}
	return w
}
