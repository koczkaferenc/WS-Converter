package models

import (
	"fmt"
)

// type Product struct {
//     "product_description.name.hu": removeBrandNamesFromProductNames(row["nev"]),
//     "product.weight": row["suly"],
//     "product.weight_class_id": "Kg.",
//     "product.image": "",
//     "product_to_category.category_name": "",
//     "product.manufacturer_id": 0,
//     "product.product_class_id": "",
//     "attr_values.anyag.hu": "",                     # anyag
//     "attr_values.rogzites.hu": "",                  # rogzites
//     "attr_values.csaptipus.hu": "",                 # csaptipus
//     "attr_values.szemforma.hu": row["szemforma"],   # szemforma
//     "attr_values.osztas": row["osztas"],            # osztas
//     "attr_values.belsoheveder": row["belsoheveder"],# gorgoslanc belsoheveder
//     "attr_values.gorgoatmero": row["gorgoatmero"],  # gorgoslanc gorgoatmero
//     "attr_values.csapatmer:"",                    # patentszem csapatmero,
//     "attr_values.csaphoss:"",                     # patentszem csap hossza
//     "attr_values.huvelyatmero": row["huvelyatmero"],# huvelyatmero huvelyes láncok esetén
//     "attr_values.fogszam": "",                      # fogszam, számoljuk skuból
//     "attr_values.fogedzett.hu": "",                 # edzett fogakkal, számoljuk skuból
//     "attr_values.kivitel.hu": "",                   # flyer és görgösláncnál normál (gyengébb) vagy erősített (heavy)
//     "attr_values.hevederek_szama.hu": "",           # flyer láncnál hevederek száma, szöveg, mert pl. 3x4
//     "attr_values.lanckerektipus.hu": "",            # lánckerék: agyas vagy laplánckerék
//     "attr_values.feluletkezeles.hu": row["feluletkezeles"], # if row["feluletkezeles"] else "Nincs",  # felulet
//     "product_special.price": row["akcios_ar"],
//     "product_special.date_start": row["akcio_kezdet"],
//     "product_special.date_end": row["akcio_veg"],
//     "product_description.short_description.hu": "",
// }
// }

// Görgőslácok struktúrája
type GL struct {
	Osztas         string
	BelsoHeveder   string
	Csapatmero     string
	Gorgoatmero    string
	Csaphossz      string
}

var GLParms = map[string]map[string]string{
	"08A1": {"Osztas":"12,7", "Belsoheveder":"7,85", "Csapatmero":"3,96", "Gorgoatmero":"7,95", "Csaphossz":"21,7"},
	"10A1": {"Osztas":"15,875", "Belsoheveder":"9,4", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"25,9"},
	"12A1": {"Osztas":"19,05", "Belsoheveder":"12,57", "Csapatmero":"5,94", "Gorgoatmero":"11,91", "Csaphossz":"31,5"},
	"16A1": {"Osztas":"25,4", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":"38,9"},
	"20A1": {"Osztas":"31,75", "Belsoheveder":"18,9", "Csapatmero":"9,53", "Gorgoatmero":"19,05", "Csaphossz":"45,2"},
	"24A1": {"Osztas":"38,1", "Belsoheveder":"25,22", "Csapatmero":"11,1", "Gorgoatmero":"22,23", "Csaphossz":"55,5"},
	"28A1": {"Osztas":"44,45", "Belsoheveder":"25,22", "Csapatmero":"12,7", "Gorgoatmero":"25,4", "Csaphossz":"59,3"},
	"32A1": {"Osztas":"50,8", "Belsoheveder":"31,55", "Csapatmero":"14,27", "Gorgoatmero":"28,58", "Csaphossz":"69,6"},
	"36A1": {"Osztas":"57,15", "Belsoheveder":"35,48", "Csapatmero":"17,46", "Gorgoatmero":"35,71", "Csaphossz":""},
	"40A1": {"Osztas":"63,5", "Belsoheveder":"37,85", "Csapatmero":"19,85", "Gorgoatmero":"39,68", "Csaphossz":"85,4"},
	"48A1": {"Osztas":"76,2", "Belsoheveder":"47,35", "Csapatmero":"23,81", "Gorgoatmero":"47,63", "Csaphossz":"103,1"},
	"08A2": {"Osztas":"12,7", "Belsoheveder":"7,85", "Csapatmero":"3,96", "Gorgoatmero":"7,95", "Csaphossz":"36,2"},
	"10A2": {"Osztas":"15,875", "Belsoheveder":"9,4", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"44"},
	"12A2": {"Osztas":"19,05", "Belsoheveder":"12,57", "Csapatmero":"5,94", "Gorgoatmero":"11,91", "Csaphossz":"44,4"},
	"16A2": {"Osztas":"25,4", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":"68,1"},
	"20A2": {"Osztas":"31,75", "Belsoheveder":"18,9", "Csapatmero":"9,53", "Gorgoatmero":"19,05", "Csaphossz":"81,2"},
	"24A2": {"Osztas":"38,1", "Belsoheveder":"25,22", "Csapatmero":"11,1", "Gorgoatmero":"22,23", "Csaphossz":"100,9"},
	"28A2": {"Osztas":"44,45", "Belsoheveder":"25,22", "Csapatmero":"12,7", "Gorgoatmero":"25,4", "Csaphossz":"108,2"},
	"32A2": {"Osztas":"50,8", "Belsoheveder":"31,55", "Csapatmero":"14,27", "Gorgoatmero":"28,58", "Csaphossz":"128,2"},
	"36A2": {"Osztas":"57,15", "Belsoheveder":"35,48", "Csapatmero":"17,46", "Gorgoatmero":"35,71", "Csaphossz":""},
	"40A2": {"Osztas":"63,5", "Belsoheveder":"37,85", "Csapatmero":"19,85", "Gorgoatmero":"39,68", "Csaphossz":"157"},
	"48A2": {"Osztas":"76,2", "Belsoheveder":"47,35", "Csapatmero":"23,81", "Gorgoatmero":"47,63", "Csaphossz":"191"},
	"08A3": {"Osztas":"12,7", "Belsoheveder":"7,85", "Csapatmero":"3,96", "Gorgoatmero":"7,95", "Csaphossz":"50,6"},
	"10A3": {"Osztas":"15,875", "Belsoheveder":"9,4", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"62"},
	"12A3": {"Osztas":"19,05", "Belsoheveder":"12,57", "Csapatmero":"5,94", "Gorgoatmero":"11,91", "Csaphossz":"77,2"},
	"16A3": {"Osztas":"25,4", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":"97,1"},
	"20A3": {"Osztas":"31,75", "Belsoheveder":"18,9", "Csapatmero":"9,53", "Gorgoatmero":"19,05", "Csaphossz":"117,2"},
	"24A3": {"Osztas":"38,1", "Belsoheveder":"25,22", "Csapatmero":"11,1", "Gorgoatmero":"22,23", "Csaphossz":"146,4"},
	"28A3": {"Osztas":"44,45", "Belsoheveder":"25,22", "Csapatmero":"12,7", "Gorgoatmero":"25,4", "Csaphossz":"157"},
	"32A3": {"Osztas":"50,8", "Belsoheveder":"31,55", "Csapatmero":"14,27", "Gorgoatmero":"28,58", "Csaphossz":"186,7"},
	"36A3": {"Osztas":"57,15", "Belsoheveder":"35,48", "Csapatmero":"17,46", "Gorgoatmero":"35,71", "Csaphossz":""},
	"40A3": {"Osztas":"63,5", "Belsoheveder":"37,85", "Csapatmero":"19,85", "Gorgoatmero":"39,68", "Csaphossz":"228,5"},
	"48A3": {"Osztas":"76,2", "Belsoheveder":"47,35", "Csapatmero":"23,81", "Gorgoatmero":"47,63", "Csaphossz":"278,8"},
	"08AH1": {"Osztas":"12,7", "Belsoheveder":"7,85", "Csapatmero":"3,96", "Gorgoatmero":"7,95", "Csaphossz":"19,9"},
	"10AH1": {"Osztas":"15,875", "Belsoheveder":"9,4", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"24,3"},
	"12AH1": {"Osztas":"19,05", "Belsoheveder":"12,57", "Csapatmero":"5,94", "Gorgoatmero":"11,91", "Csaphossz":"31"},
	"16AH1": {"Osztas":"25,4", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":"42,4"},
	"20AH1": {"Osztas":"31,75", "Belsoheveder":"18,9", "Csapatmero":"9,53", "Gorgoatmero":"19,05", "Csaphossz":"50,6"},
	"24AH1": {"Osztas":"38,1", "Belsoheveder":"25,22", "Csapatmero":"11,1", "Gorgoatmero":"22,23", "Csaphossz":"61,4"},
	"28AH1": {"Osztas":"44,45", "Belsoheveder":"25,22", "Csapatmero":"12,7", "Gorgoatmero":"25,4", "Csaphossz":"66,1"},
	"32AH1": {"Osztas":"50,8", "Belsoheveder":"31,55", "Csapatmero":"14,27", "Gorgoatmero":"28,58", "Csaphossz":"75,4"},
	"36AH1": {"Osztas":"57,15", "Belsoheveder":"35,48", "Csapatmero":"17,46", "Gorgoatmero":"35,71", "Csaphossz":""},
	"40AH1": {"Osztas":"63,5", "Belsoheveder":"37,85", "Csapatmero":"19,85", "Gorgoatmero":"39,68", "Csaphossz":"95,6"},
	"48AH1": {"Osztas":"76,2", "Belsoheveder":"47,35", "Csapatmero":"23,81", "Gorgoatmero":"47,63", "Csaphossz":"115,5"},
	"12AH2": {"Osztas":"19,05", "Belsoheveder":"12,57", "Csapatmero":"5,94", "Gorgoatmero":"11,91", "Csaphossz":""},
	"16AH2": {"Osztas":"25,4", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":""},
	"20AH2": {"Osztas":"31,75", "Belsoheveder":"18,9", "Csapatmero":"9,53", "Gorgoatmero":"19,05", "Csaphossz":""},
	"24AH2": {"Osztas":"38,1", "Belsoheveder":"25,22", "Csapatmero":"11,1", "Gorgoatmero":"22,23", "Csaphossz":""},
	"28AH2": {"Osztas":"44,45", "Belsoheveder":"25,22", "Csapatmero":"12,7", "Gorgoatmero":"25,4", "Csaphossz":""},
	"32AH2": {"Osztas":"50,8", "Belsoheveder":"31,55", "Csapatmero":"14,27", "Gorgoatmero":"28,58", "Csaphossz":""},
	"40AH2": {"Osztas":"63,5", "Belsoheveder":"37,85", "Csapatmero":"19,85", "Gorgoatmero":"39,68", "Csaphossz":""},
	"12AH3": {"Osztas":"19,05", "Belsoheveder":"12,57", "Csapatmero":"5,94", "Gorgoatmero":"11,91", "Csaphossz":""},
	"16AH3": {"Osztas":"25,4", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":""},
	"20AH3": {"Osztas":"31,75", "Belsoheveder":"18,9", "Csapatmero":"9,53", "Gorgoatmero":"19,05", "Csaphossz":""},
	"24AH3": {"Osztas":"38,1", "Belsoheveder":"25,22", "Csapatmero":"11,1", "Gorgoatmero":"22,23", "Csaphossz":""},
	"28AH3": {"Osztas":"44,45", "Belsoheveder":"25,22", "Csapatmero":"12,7", "Gorgoatmero":"25,4", "Csaphossz":""},
	"32AH3": {"Osztas":"50,8", "Belsoheveder":"31,55", "Csapatmero":"14,27", "Gorgoatmero":"28,58", "Csaphossz":""},
	"40AH3": {"Osztas":"63,5", "Belsoheveder":"37,85", "Csapatmero":"19,85", "Gorgoatmero":"39,68", "Csaphossz":""},

	"04B1": {"Osztas":"6,000", "Belsoheveder":"2,80", "Csapatmero":"1,85", "Gorgoatmero":"4,00", "Csaphossz":"7,8"},
	"05B1": {"Osztas":"8,000", "Belsoheveder":"3,00", "Csapatmero":"2,31", "Gorgoatmero":"5,00", "Csaphossz":"11,7"},
	"06B1": {"Osztas":"9,525", "Belsoheveder":"5,72", "Csapatmero":"3,28", "Gorgoatmero":"6,35", "Csaphossz":"16,8"},
	"08B1": {"Osztas":"12,700", "Belsoheveder":"7,75", "Csapatmero":"4,45", "Gorgoatmero":"8,51", "Csaphossz":"20,9"},
	"10B1": {"Osztas":"15,875", "Belsoheveder":"9,65", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"23,7"},
	"12B1": {"Osztas":"19,050", "Belsoheveder":"11,68", "Csapatmero":"5,72", "Gorgoatmero":"12,07", "Csaphossz":"27,3"},
	"16B1": {"Osztas":"25,400", "Belsoheveder":"17,02", "Csapatmero":"8,28", "Gorgoatmero":"15,88", "Csaphossz":"41,5"},
	"20B1": {"Osztas":"31,750", "Belsoheveder":"19,56", "Csapatmero":"10,19", "Gorgoatmero":"19,05", "Csaphossz":"46,0"},
	"24B1": {"Osztas":"38,100", "Belsoheveder":"25,40", "Csapatmero":"14,63", "Gorgoatmero":"25,40", "Csaphossz":"58,5"},
	"28B1": {"Osztas":"44,450", "Belsoheveder":"30,99", "Csapatmero":"15,90", "Gorgoatmero":"27,94", "Csaphossz":"69,6"},
	"32B1": {"Osztas":"50,800", "Belsoheveder":"30,99", "Csapatmero":"17,81", "Gorgoatmero":"29,21", "Csaphossz":"73,1"},
	"40B1": {"Osztas":"63,500", "Belsoheveder":"38,10", "Csapatmero":"22,89", "Gorgoatmero":"39,37", "Csaphossz":"86,3"},
	"48B1": {"Osztas":"76,200", "Belsoheveder":"45,72", "Csapatmero":"29,24", "Gorgoatmero":"48,26", "Csaphossz":"107,9"},
	"56B1": {"Osztas":"88,900", "Belsoheveder":"53,34", "Csapatmero":"34,32", "Gorgoatmero":"53,98", "Csaphossz":"137"},
	"64B1": {"Osztas":"101,600", "Belsoheveder":"60,96", "Csapatmero":"39,40", "Gorgoatmero":"63,50", "Csaphossz":"138,5"},
	"72B1": {"Osztas":"114,300", "Belsoheveder":"68,58", "Csapatmero":"44,48", "Gorgoatmero":"72,39", "Csaphossz":"156,4"},

	"05B2": {"Osztas":"8,000", "Belsoheveder":"3,00", "Csapatmero":"2,31", "Gorgoatmero":"5,00", "Csaphossz":"17,4"},
	"06B2": {"Osztas":"9,525", "Belsoheveder":"5,72", "Csapatmero":"3,28", "Gorgoatmero":"6,35", "Csaphossz":"27,1"},
	"08B2": {"Osztas":"12,700", "Belsoheveder":"7,75", "Csapatmero":"4,45", "Gorgoatmero":"8,51", "Csaphossz":"34,9"},
	"10B2": {"Osztas":"15,875", "Belsoheveder":"9,65", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"40,3"},
	"12B2": {"Osztas":"19,050", "Belsoheveder":"11,68", "Csapatmero":"5,72", "Gorgoatmero":"12,07", "Csaphossz":"46,8"},
	"16B2": {"Osztas":"25,400", "Belsoheveder":"17,02", "Csapatmero":"8,28", "Gorgoatmero":"15,88", "Csaphossz":"73,4"},
	"20B2": {"Osztas":"31,750", "Belsoheveder":"19,56", "Csapatmero":"10,19", "Gorgoatmero":"19,05", "Csaphossz":"82,5"},
	"24B2": {"Osztas":"38,100", "Belsoheveder":"25,40", "Csapatmero":"14,63", "Gorgoatmero":"25,40", "Csaphossz":"106,9"},
	"28B2": {"Osztas":"44,450", "Belsoheveder":"30,99", "Csapatmero":"15,90", "Gorgoatmero":"27,94", "Csaphossz":"129,2"},
	"32B2": {"Osztas":"50,800", "Belsoheveder":"30,99", "Csapatmero":"17,81", "Gorgoatmero":"29,21", "Csaphossz":"131,7"},
	"40B2": {"Osztas":"63,500", "Belsoheveder":"38,10", "Csapatmero":"22,89", "Gorgoatmero":"39,37", "Csaphossz":"158,6"},
	"48B2": {"Osztas":"76,200", "Belsoheveder":"45,72", "Csapatmero":"29,24", "Gorgoatmero":"48,26", "Csaphossz":"199,1"},
	"56B2": {"Osztas":"88,900", "Belsoheveder":"53,34", "Csapatmero":"34,32", "Gorgoatmero":"53,98", "Csaphossz":"243,6"},
	"64B2": {"Osztas":"101,600", "Belsoheveder":"60,96", "Csapatmero":"39,40", "Gorgoatmero":"63,50", "Csaphossz":"258,4"},
	"72B2": {"Osztas":"114,300", "Belsoheveder":"68,58", "Csapatmero":"44,48", "Gorgoatmero":"72,39", "Csaphossz":"292,7"},

	"05B3": {"Osztas":"8,000", "Belsoheveder":"3,00", "Csapatmero":"2,31", "Gorgoatmero":"5,00", "Csaphossz":"23"},
	"06B3": {"Osztas":"9,525", "Belsoheveder":"5,72", "Csapatmero":"3,28", "Gorgoatmero":"6,35", "Csaphossz":"37,3"},
	"08B3": {"Osztas":"12,700", "Belsoheveder":"7,75", "Csapatmero":"4,45", "Gorgoatmero":"8,51", "Csaphossz":"48,8"},
	"10B3": {"Osztas":"15,875", "Belsoheveder":"9,65", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"56,9"},
	"12B3": {"Osztas":"19,050", "Belsoheveder":"11,68", "Csapatmero":"5,72", "Gorgoatmero":"12,07", "Csaphossz":"66,3"},
	"16B3": {"Osztas":"25,400", "Belsoheveder":"17,02", "Csapatmero":"8,28", "Gorgoatmero":"15,88", "Csaphossz":"105,3"},
	"20B3": {"Osztas":"31,750", "Belsoheveder":"19,56", "Csapatmero":"10,19", "Gorgoatmero":"19,05", "Csaphossz":"118,9"},
	"24B3": {"Osztas":"38,100", "Belsoheveder":"25,40", "Csapatmero":"14,63", "Gorgoatmero":"25,40", "Csaphossz":"155,2"},
	"28B3": {"Osztas":"44,450", "Belsoheveder":"30,99", "Csapatmero":"15,90", "Gorgoatmero":"27,94", "Csaphossz":"188,8"},
	"32B3": {"Osztas":"50,800", "Belsoheveder":"30,99", "Csapatmero":"17,81", "Gorgoatmero":"29,21", "Csaphossz":"190,2"},
	"40B3": {"Osztas":"63,500", "Belsoheveder":"38,10", "Csapatmero":"22,89", "Gorgoatmero":"39,37", "Csaphossz":"230,9"},
	"48B3": {"Osztas":"76,200", "Belsoheveder":"45,72", "Csapatmero":"29,24", "Gorgoatmero":"48,26", "Csaphossz":"293,3"},
	"56B3": {"Osztas":"88,900", "Belsoheveder":"53,34", "Csapatmero":"34,32", "Gorgoatmero":"53,98", "Csaphossz":"350,2"},
	"64B3": {"Osztas":"101,600", "Belsoheveder":"60,96", "Csapatmero":"39,40", "Gorgoatmero":"63,50", "Csaphossz":"378,3"},
	"72B3": {"Osztas":"114,300", "Belsoheveder":"68,58", "Csapatmero":"44,48", "Gorgoatmero":"72,39", "Csaphossz":"429,0"},

	"04BH": {"Osztas":"6,000", "Belsoheveder":"2,80", "Csapatmero":"1,85", "Gorgoatmero":"4,00", "Csaphossz":"9,4"},
	"06BH": {"Osztas":"9,525", "Belsoheveder":"5,72", "Csapatmero":"3,58", "Gorgoatmero":"6,35", "Csaphossz":"15,4"},
	"08BH": {"Osztas":"12,700", "Belsoheveder":"7,85", "Csapatmero":"4,45", "Gorgoatmero":"8,51", "Csaphossz":"19,9"},
	"10BH": {"Osztas":"15,875", "Belsoheveder":"9,65", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"21,6"},
	"12BH": {"Osztas":"19,050", "Belsoheveder":"11,68", "Csapatmero":"5,94", "Gorgoatmero":"12,07", "Csaphossz":"26,8"},
	"12BV": {"Osztas":"19,050", "Belsoheveder":"11,68", "Csapatmero":"6,10", "Gorgoatmero":"12,07", "Csaphossz":"26,5"},
	"12BHF2": {"Osztas":"19,050", "Belsoheveder":"13,50", "Csapatmero":"5,72", "Gorgoatmero":"12,07", "Csaphossz":"30,3"},
	"16BH": {"Osztas":"25,400", "Belsoheveder":"17,02", "Csapatmero":"8,90", "Gorgoatmero":"15,88", "Csaphossz":"38,9"},
	"24BH": {"Osztas":"38,100", "Belsoheveder":"25,40", "Csapatmero":"14,63", "Gorgoatmero":"25,40", "Csaphossz":"63,4"},
	"24BHF2": {"Osztas":"38,100", "Belsoheveder":"25,40", "Csapatmero":"14,63", "Gorgoatmero":"25,40", "Csaphossz":"62,2"},

	"208A1": {"Osztas":"25,4", "Belsoheveder":"7,85", "Csapatmero":"3,96", "Gorgoatmero":"7,95", "Csaphossz":"17,8"},
	"208B1": {"Osztas":"25,4", "Belsoheveder":"7,75", "Csapatmero":"4,45", "Gorgoatmero":"8,51", "Csaphossz":"18,2"},
	"210A1": {"Osztas":"31,75", "Belsoheveder":"9,4", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"22,2"},
	"210B1": {"Osztas":"31,75", "Belsoheveder":"9,65", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"20,9"},
	"212A1": {"Osztas":"38,1", "Belsoheveder":"12,57", "Csapatmero":"5,94", "Gorgoatmero":"11,91", "Csaphossz":"27,7"},
	"212B1": {"Osztas":"38,1", "Belsoheveder":"11,68", "Csapatmero":"5,72", "Gorgoatmero":"12,07", "Csaphossz":"25,2"},
	"216A1": {"Osztas":"50,8", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":"36,5"},
	"216A1H": {"Osztas":"50,8", "Belsoheveder":"15,75", "Csapatmero":"7,92", "Gorgoatmero":"15,88", "Csaphossz":"39,4"},
	"216B1": {"Osztas":"50,8", "Belsoheveder":"17,02", "Csapatmero":"8,28", "Gorgoatmero":"15,88", "Csaphossz":"41,5"},
	"220A1": {"Osztas":"63,5", "Belsoheveder":"18,9", "Csapatmero":"9,53", "Gorgoatmero":"19,05", "Csaphossz":"44,7"},
	"220B1": {"Osztas":"63,5", "Belsoheveder":"19,56", "Csapatmero":"10,19", "Gorgoatmero":"19,05", "Csaphossz":"46"},
	"224A1": {"Osztas":"76,2", "Belsoheveder":"25,22", "Csapatmero":"11,1", "Gorgoatmero":"22,23", "Csaphossz":"54,3"},
	"224B1": {"Osztas":"76,2", "Belsoheveder":"25,4", "Csapatmero":"14,63", "Gorgoatmero":"25,4", "Csaphossz":"58,5"},
	"228B1": {"Osztas":"88,9", "Belsoheveder":"30,99", "Csapatmero":"15,9", "Gorgoatmero":"27,94", "Csaphossz":"69,5"},
	"232B1": {"Osztas":"101,6", "Belsoheveder":"30,99", "Csapatmero":"17,81", "Gorgoatmero":"29,21", "Csaphossz":"71"},

	// Manuális úton felvéve
	"04C1": {"Osztas":"6,35", "Belsoheveder":"3,18", "Csapatmero":"2,31", "Gorgoatmero":"3,3", "Csaphossz":"8,4"},
	"031" : {"Osztas":"5", "Belsoheveder":"2,5", "Csapatmero":"1,49", "Gorgoatmero":"3,2", "Csaphossz":"9,9"},
	"0811":  {"Osztas":"12,7", "Belsoheveder":"3,3", "Csapatmero":"3,66", "Gorgoatmero":"7,75", "Csaphossz":"11,7"},
	"062C1": {"Osztas":"9,525", "Belsoheveder":"9,52", "Csapatmero":"4,18", "Gorgoatmero":"6", "Csaphossz":""},
	"06C1": {"Osztas":"9,525", "Belsoheveder":"4,77", "Csapatmero":"3,58", "Gorgoatmero":"5,08", "Csaphossz":"13,3"},
	"06C2": {"Osztas":"9,525", "Belsoheveder":"4,55", "Csapatmero":"3,58", "Gorgoatmero":"5,08", "Csaphossz":"23,45"},
	"0831": {"Osztas":"12,7", "Belsoheveder":"4,88", "Csapatmero":"4,09", "Gorgoatmero":"7,75", "Csaphossz":"12,9"},
	"0861": {"Osztas":"12,7", "Belsoheveder":"5,3", "Csapatmero":"4,45", "Gorgoatmero":"8,51", "Csaphossz":"15,9"},
	"1011": {"Osztas":"15,88", "Belsoheveder":"6,48", "Csapatmero":"5,08", "Gorgoatmero":"10,16", "Csaphossz":"17,2"},
	"411": {"Osztas":"12,7", "Belsoheveder":"6,25", "Csapatmero":"3,58", "Gorgoatmero":"7,77", "Csaphossz":"15,0"},
	"04B2": {"Osztas":"", "Belsoheveder":"", "Csapatmero":"", "Gorgoatmero":"", "Csaphossz":""},

	// VELO
	"1/2x3/16": {"Osztas":"12,7", "Belsoheveder":"4,88", "Csapatmero":"7,75", "Gorgoatmero":"3,66", "Csaphossz":"12,3"},
	"1/2x1/8":  {"Osztas":"12,7", "Belsoheveder":"3,3", "Csapatmero":"7,75", "Gorgoatmero":"3,66", "Csaphossz":"10,2"},
	// MOFA
	"0841_MOFA":  {"Osztas":"12,7", "Belsoheveder":"4,88", "Csapatmero":"4,18", "Gorgoatmero":"7,75", "Csaphossz":"14,3"},
	"0851_MOFA":  {"Osztas":"12,7", "Belsoheveder":"6,4",  "Csapatmero":"4,18", "Gorgoatmero":"7,75", "Csaphossz":"15,9"},
}

func main() {
	fmt.Println("GL.")
}
