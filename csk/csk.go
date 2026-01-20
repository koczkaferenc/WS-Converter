package csk
// Csőkinyomó
//
import (
	"fmt"
	"regexp"
	"strconv"
	"slices"
	"ws-updater/models"
)

func ProcessCsk(p models.KsProduct) models.WsProduct {
	var (
		regExpCsk = regexp.MustCompile(`N-(CSK)-([0-9]+)-([0-9]+)-([0-9]+)$`)
		match          []string
		family         string
		manufacturerId int
		h      string
		l      string
	)

	w := models.WsProduct{}

	w.SKU = p.Code
	w.Category = "Csapkinyomók"
	w.ClassId = "Csapkinyomó"
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

	// N-CSK-3-12-20
	// N-(CSK)-([0-9]+)-([0-9]+)-([0-9]+)$
	match = regExpCsk.FindStringSubmatch(p.Code)
 	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		w.Manufacturer = models.Manufacturers[manufacturerId]
		l = match[3]
		h = match[4]
	}

	w.ShortDescription = fmt.Sprintf("%s gyártmányú csapkinyomó szerszám %sB1-%sB1 méretű láncokhoz. ", w.Manufacturer, h, l)
	if w.Image == "" {
		w.Image = fmt.Sprintf("product/N-%s.png", family)
	}
	return w
}
