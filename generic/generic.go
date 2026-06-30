package generic

import (
	"fmt"
	"ws-updater/models"
)

var genericTable = map[string]map[string]string{
	"N-KR-4-12B1_Z18_G": {"TargetMachine": "McHale Pro 1", "TargetCategory": "McHale"},
}

func ProcessProd(p models.KsProduct) {
	_, found := genericTable[p.Code]
	if !found {
		return
	}
	machine := genericTable[p.Code]["TargetMachine"]

	fmt.Printf("Generic: %s\n", machine)
}
