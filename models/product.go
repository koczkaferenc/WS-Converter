package models

const Zaradek = "<br><strong>Figyelem!</strong><br>A termékkép illusztráció, a pontos műszaki tartalmat a cikkszám és a termék megnevezése tartalmazza! Kérdés, kérés esetén hívja munkatársunkat, vagy vegye fel velünk a kapcsolatot!"

// Ezek a termékcsoportok csak rendelésre kaphatók.
var CsakRendelesre = []string{"FGL", "FGLPSZ", "FGLSPSZ", "CSCSGL", "CSCSGLPSZ", "CSCSGLSPSZ", "CSGL", "CSGLPSZ", "CSGLSPSZ"}

const CsakRendelesreLeiras = "<p>Mivel a termék számos egyéb paraméterrel rendelkezik, így csak egyeztetést követően rendelhető. Kérjük, hívja munkatársunkat, vagy vegye fel a kapcsolatot velünk az elérhetőségeink valamelyikén.</p>"

const JelenlegNemElerheto = "<p><strong>A termék jelenleg nincs raktáron.</strong></p>"

// KS Termék adatok - ezeket olvassuk be a Firebirdből
type KsProduct struct {
	ID            int
	Code          string
	Name          string
	Unit          string
	Weight        float64
	Stock         float64
	WebPrice      float64
	ReSellerPrice float64
}

// A WebShop kimenet mezői

// / TODO 0 mennyiségűhöz odaírni, hogy csak rendelésre!!!
type WsProduct struct {
	SKU              string `csv:"product.sku"`
	Name             string `csv:"product_description.name.hu"`
	Weight           string `csv:"product.weight"`
	WeightClass      string `csv:"product.weight_class_id"`
	Image            string `csv:"product.image"`
	Category         string `csv:"product_to_category.category_name"`
	Manufacturer     string `csv:"product.manufacturer_id"`
	ClassId          string `csv:"product.product_class_id"`
	Anyag            string `csv:"attr_values.anyag.hu"`
	Rogzites         string `csv:"attr_values.rogzites.hu"`
	Csaptipus        string `csv:"attr_values.csaptipus.hu"`
	Szemforma        string `csv:"attr_values.szemforma.hu"`
	Osztas           string `csv:"attr_values.osztas"`
	BelsoHeveder     string `csv:"attr_values.belsoheveder"`
	GorgoAtmero      string `csv:"attr_values.gorgoatmero"`
	CsapAtmero       string `csv:"attr_values.csapatmero"`
	CsapHossz        string `csv:"attr_values.csaphossz"`
	HuvelyAtmero     string `csv:"attr_values.huvelyatmero"`
	Fogszam          string `csv:"attr_values.fogszam"`
	Fogedzett        string `csv:"attr_values.fogedzett.hu"`
	Kivitel          string `csv:"attr_values.kivitel.hu"`
	HevederSzam      string `csv:"attr_values.hevederek_szama.hu"`
	LanckerekTipus   string `csv:"attr_values.lanckerektipus.hu"`
	Feluletkezeles   string `csv:"attr_values.feluletkezeles.hu"`
	SpecialPrice     string `csv:"product_special.price"`
	SpecialStart     string `csv:"product_special.date_start"`
	SpecialEnd       string `csv:"product_special.date_end"`
	ShortDescription string `csv:"product_description.short_description.hu"`
	Quantity         string `csv:"product.quantity_2"`
	Alapar           string `csv:"product.alapar"`
	TaxClass         string `csv:"product.tax_class_id"`
	QuantityUnit     string `csv:"product_description.quantity_name.hu"`
	ImageAdditional  string `csv:"product_image.image.0"`
	HuzalAtmero      string `csv:"attr_values.huzal_atmero"`
	BelsoHossz       string `csv:"attr_values.belso_hossz"`
}

// Gyártók
var Manufacturers = map[int]string{
	0:  "Prémium",      // Renold
	1:  "Távol-keleti", // MSC
	2:  "Távol-keleti", // Lovas
	3:  "Távol-keleti", // TEC
	4:  "Európai",      // Codex
	5:  "Európai",      // Vamberk
	6:  "Európai",      // Strakonice
	7:  "Prémium",      // Rexnord
	8:  "Európai",      // Link-Belt
	9:  "Európai",      // Retezarna
	10: "Európai",      // Reiter
	11: "Távol-keleti",
	12: "AL",
}
