//go:build ignore
// +build ignore

package mgbf

// Mezőgazdasági boronafog

import (
	"fmt"
	"slices"
	"strconv"
	"ws-updater/models"
)

func ProcessMgbf(p models.KsProduct) models.WsProduct {
	var (
		family string
	)
	fmt.Println(p.Code)
	w := models.WsProduct{}

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
	return w

	//#	var (
	//#	regExpMGBF = regexp.MustCompile(`N-(MGBF)([A-Z]+)-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_([A-Z]+)$`)
	//#	match          []string
	//#	family         string
	//#	manufacturerId int
	//#	x string
	//#	y string
	//#	m string
	//#)

	// #w.SKU = p.Code
	// #w.Category = "Csapkinyomók"
	// #w.ClassId = "Csapkinyomó"
	// #w.Quantity = fmt.Sprintf("%.1f", p.Stock)
	// #w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
	// #w.TaxClass = "27%"
	// #w.QuantityUnit = p.Unit
	// #w.Weight = fmt.Sprintf("%.1f", p.Weight)
	// #w.WeightClass = "kg."
	// #
	// #// N-CSK-3-12-20
	// #// N-(MGBF)([A-Z]+)-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_([A-Z]+)$
	// #// Termék kód és példa
	// #    // 1: family
	// #    // 2: E-Egyenes, K-Kanalas, H-Hajlított
	// #    // 3: Gyártó kód
	// #    // 4: Szélesség
	// #    // 5: Magasság
	// #    // 6: Hossz
	// #    // 7: Csavar menet méret
	// #    // 8: LL: lap-lappal, LE: Lap éllel
	// #match = regExpMGBF.FindStringSubmatch(p.Code)
	// # 	if match != nil {
	// #	family = match[1]
	// #	manufacturerId, _ = strconv.Atoi(match[2])
	// #	w.Manufacturer = models.Manufacturers[manufacturerId]
	// #	x = match[3]
	// #	y = match[4]
	// #	hossz=match[5]
	// #	m = match[6]
	// #}
	// #
	// #w.ShortDescription = fmt.Sprintf("%s gyártmányú csapkinyomó szerszám %sB1-%sB1 méretű láncokhoz. ", w.Manufacturer, h, l)
	// #if w.Image == "" {
	// #	w.Image = fmt.Sprintf("product/N-%s.png", family)
	// #}
	return w

}
