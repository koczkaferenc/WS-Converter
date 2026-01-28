package mgbf

// Mezőgazdasági boronafog
// TODO: képekben benne kell lenni az LL|LÉ-nek

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"ws-updater/models"
)

func ProcessMgbf(p models.KsProduct) models.WsProduct {
	var (
		family         string
		kivitel        string
		regExpMGBF     = regexp.MustCompile(`N-(MGBF)([EHK])-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_(LL|LÉ)$`)
		match          []string
		manufacturerId int
		x              string
		y              string
		hossz          string
		m              string
		ph             string
	)

	w := models.WsProduct{}
	w.SKU = p.Code
	w.Category = "Boronafogak"
	w.ClassId = "Boronafog"
	w.Quantity = fmt.Sprintf("%.1f", p.Stock)
	w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
	w.TaxClass = "27%"
	w.QuantityUnit = p.Unit
	w.Weight = fmt.Sprintf("%.1f", p.Weight)
	w.WeightClass = "kg."
	w.Anyag = "Kovácsolt vas"

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
		kivitel = match[2]
		switch kivitel {
		case "E":
			w.Kivitel = "Egyenes"
		case "H":
			w.Kivitel = "Hajlított"
		case "K":
			w.Kivitel = "Kanalas"
		}
		manufacturerId, _ = strconv.Atoi(match[3])
		w.Manufacturer = models.Manufacturers[manufacturerId]

		x = match[4]
		y = match[5]
		hossz = match[6]
		m = match[7]
		switch match[8] {
		case "LL":
			ph = "párhuzamos"
		case "LÉ":
			ph = "elforgatott"
		}
	}

	w.Name = fmt.Sprintf("%sx%sx%sM%s %s %s boronafog", x, y, hossz, m, ph, strings.ToLower(w.Kivitel))
	if w.Image == "" {
		w.Image = fmt.Sprintf("product/N-%s%s.png", family, kivitel)
	}

	w.ShortDescription = fmt.Sprintf("%s %sx%sx%s mm, M%s %s %s boronafog. ", w.Manufacturer, x, y, hossz, m, ph, strings.ToLower(w.Kivitel))
	w.ShortDescription += "<br><font color='red'>A termék csak rendelhető, kérjük, hívja munkatársunkat!</font>"
	w.ShortDescription += models.Zaradek

	return w

}
