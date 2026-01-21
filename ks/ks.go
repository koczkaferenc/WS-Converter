package ks

// Lánckerék szelet
import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

func ProcessKs(p models.KsProduct) models.WsProduct {
	var (
		regExpKS   = regexp.MustCompile(`N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
		regExpKS_G = regexp.MustCompile(`N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
		// Agyas lánckerekek
		regExpKR   = regexp.MustCompile(`N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
		regExpKR_G = regexp.MustCompile(`N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
		// Sajnos ezt ebben a formában is felvitték
		regExpGKR      = regexp.MustCompile(`N-(GKR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
		match          []string
		family         string
		productType    string
		manufacturerId int
		numOfRows      string
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
	w.Category = "Lánckerekek"
	w.ClassId = "Lánckerék"
	w.LanckerekTipus = "Laplánckerék" // Agyas lánckerék |  Laplánckerék
	w.Fogedzett = "Standard"          // Fogedzett | Standard

	// N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
	// N-KR-0-08B2_Z30
	match = regExpKS.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		productType = match[3]
		numOfRows = match[4]
		w.Fogszam = match[5]
	}

	// N-(KS)-([0-9])+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$
	// N-KS-0-08B2_Z21_G
	match = regExpKS_G.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		productType = match[3]
		numOfRows = match[4]
		w.Fogszam = match[5]
		w.Fogedzett = "Fogedzett"
	}

	// N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
	// N-KR-0-08B2_Z30
	match = regExpKR.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		productType = match[3]
		numOfRows = match[4]
		w.Fogszam = match[5]
		w.LanckerekTipus = "Agyas lánckerék"
	}

	// N-(KR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$
	// N-KR-0-08B2_Z21_G
	match = regExpKR_G.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		productType = match[3]
		numOfRows = match[4]
		w.Fogszam = match[5]
		w.Fogedzett = "Fogedzett"
		w.LanckerekTipus = "Agyas lánckerék"
	}

	// N-(GKR)-([0-9]+)-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$
	// N-GKR-0-08B2_Z30
	match = regExpGKR.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId, _ = strconv.Atoi(match[2])
		productType = match[3]
		numOfRows = match[4]
		w.Fogszam = match[5]
		w.LanckerekTipus = "Agyas lánckerék"
	}

	w.Manufacturer = models.Manufacturers[manufacturerId]
	w.ShortDescription = fmt.Sprintf("%s gyártmányú %ssoros %s fogszámú %s %s. ",
		w.Manufacturer, models.Sornevek[numOfRows], w.Fogszam, strings.ToLower(w.Fogedzett), strings.ToLower(w.LanckerekTipus))

	if w.LanckerekTipus == "Agyas lánckerék" {
		w.ShortDescription += "<p style='color: red;'>Amennyiben a lánckereket megmunkálva kívánja beszerezni, vegye fel a kapcsolatot munkatársunkkal!</p>"
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

	if w.Name == "" {
		w.Name = fmt.Sprintf("%s %s%sZ=%s %s", family, productType, numOfRows, w.Fogszam, strings.ToLower(w.LanckerekTipus))
	}
	if w.Fogedzett == "Fogedzett" {
		w.Name += ` "G"`
	}
	if w.Image == "" {
		w.Image = fmt.Sprintf("product/N-%s-%s.png", family, numOfRows)
	}
	if w.ImageAdditional == "" {
		w.ImageAdditional = fmt.Sprintf("product/D-%s-%s.png", family, numOfRows)
	}

	return w
}
