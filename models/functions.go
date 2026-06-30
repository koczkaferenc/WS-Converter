package models

import (
	"fmt"
	"regexp"
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
