package models

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

/************************************************************
* Kapcsolódó termékek listájának előállítása
************************************************************/
func getRelatedProductIds(rgxStr string, prodCodes *[]string) string {
	accessories := ""
	rgx := regexp.MustCompile(rgxStr)
	for _, code := range *prodCodes {
		match := rgx.FindStringSubmatch(code)
		if match != nil {
			accessories += code + ","
		}
	}
	if accessories != "" {
		accessories = strings.TrimSuffix(accessories, ",")
	}
	return accessories
}

func MkFeaturesList(features map[string]string) string {
	accessoriesList := ""
	i := 1
	for k, v := range features {
		if v != "" {
			v = strings.ReplaceAll(v, ",", ".")
			accessoriesList += fmt.Sprintf("%s:%s:%d,", k, v, i)
			i++
		}
	}
	if accessoriesList != "" {
		accessoriesList = strings.TrimSuffix(accessoriesList, ",")
	}
	return accessoriesList
}

func SetLabels(p *PsProduct, features map[string]string, family string) {
	p.Summary = p.Description

	// Rendelhetőség
	qtyTmp, _ := strconv.ParseFloat(p.Quantity, 64)
	if qtyTmp == 0 {
		// Nincs a termékből raktáron
		p.AvailableForOrder = "0"
		p.OutOfStockAction = "2"
		p.ShowPrice = "1"
		features["Elérhető"] = "Rendelésre"
		p.Summary += JelenlegNemElerheto
		p.Summary += "<hr>A szállítási határidőről egyeztessen munkatársunkkal!"

		// ha nincs ára, és nincs is belőle, akkor csak rendelésre
		if p.UnitPrice == "0" {
			p.AvailableForOrder = "1"
			p.OutOfStockAction = "2"
			p.ShowPrice = "0"
			features["Elérhető"] = "Árajánlattal"
			//fmt.Printf("Z-0-0: %-20s %8s Ft. %6s m\n", p.Reference, p.UnitPrice, p.Quantity)
		}
	} else {
		// Ha nincs ára, de van belőle, árajánlatot kell kérnie
		if p.UnitPrice == "0" {
			p.AvailableForOrder = "1"
			p.OutOfStockAction = "2"
			p.ShowPrice = "0"
			features["Elérhető"] = "Árajánlattal"
			p.Summary += JelenlegNemElerheto
			p.Summary += "<hr>A szállítási határidőről egyeztessen munkatársunkkal!"
			//fmt.Printf("Z-0-1: %-20s %8s Ft. %6s m\n", p.Reference, p.UnitPrice, p.Quantity)
		} else {
			// Ha van raktáron, az beragadt, ezért 5%-os engedménnyel akciózzuk.
			// p.OnSale = "1"
			// p.DiscountPercent = KiarusitasSzazalek
			// Most van belőle, lehet rendelni
			p.AvailableForOrder = "1"
			p.OutOfStockAction = "0"
			p.ShowPrice = "1"
			features["Elérhető"] = "Raktáron"
			p.Summary += fmt.Sprintf("<hr>Készleten: %g %s", qtyTmp, p.Unity)
			p.Summary += fmt.Sprintf("<hr>Súly: %s kg.", p.Weight)
			p.Summary += "<hr>Szállítási idő: 1-2 nap."
		}
	}
	// Egyáltalán lehet rendelni a terméket?
	if slices.Contains(CsakRendelesre, family) {
		p.AvailableForOrder = "0"
		p.OutOfStockAction = "2"
		p.ShowPrice = "0"
		features["Elérhető"] = "Árajánlattal"
		p.Summary = p.Description
		p.Summary += CsakRendelesreLeiras
		p.Summary += "<hr>A szállítási határidőről egyeztessen munkatársunkkal!"
	}

	p.Summary += Zaradek
	p.MetaDescription = strings.TrimRight(p.Description, ".")
	p.MetaTitle = p.Name
	p.URLRewritten = p.Reference
}
