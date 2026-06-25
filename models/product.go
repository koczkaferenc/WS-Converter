package models

const Zaradek = "<br><strong>Figyelem!</strong><br>A termékkép illusztráció, a pontos műszaki tartalmat a cikkszám és a termék megnevezése tartalmazza! Kérdés, kérés esetén hívja munkatársunkat, vagy vegye fel velünk a kapcsolatot!"

// Ezek a termékcsoportok csak rendelésre kaphatók.
var CsakRendelesre = []string{"FGL", "FGLPSZ", "FGLSPSZ", "CSCSGL", "CSCSGLPSZ", "CSCSGLSPSZ", "CSGL", "CSGLPSZ", "CSGLSPSZ", "MGBF"}

const CsakRendelesreLeiras = "<p>Mivel a termék számos egyéb paraméterrel rendelkezik, így csak egyeztetést követően rendelhető. Kérjük, hívja munkatársunkat, vagy vegye fel a kapcsolatot velünk az elérhetőségeink valamelyikén.</p>"

const JelenlegNemElerheto = "<p><strong>A termék jelenleg nincs raktáron.</strong></p>"

var Sornevek = map[string]string{"1": "egy", "2": "két", "3": "három"}

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

// A Prestashop kimenet mezői
type PsProduct struct {
	id                 string `csv:"Product ID"`
	aktiv              string `csv:"Active (0/1)"`
	nev                string `csv:"Name *"`
	kategoria_nev      string `csv:"Categories (x,y,z...)"`
	netto_ar           string `csv:"Price tax excluded"`
	ado_kategoria      string `csv:"Tax rules ID"`
	beszerzesi_ar      `csv:"Wholesale price"`
	akcios             `csv:"On sale (0/1)"`
	diszkont_engedmeny `csv:"Discount amount"`
	diszkont_szazalek  `csv:"Discount percent"`
	diszkont_tol       `csv:"Discount from (yyyy-mm-dd)"`
	diszkont_ig        `csv:"Discount to (yyyy-mm-dd)"`
	string             `csv:"Reference #"`
	string             `csv:"Supplier reference #"`
	string             `csv:"Supplier"`
	string             `csv:"Manufacturer"`
	string             `csv:"EAN13"`
	string             `csv:"UPC"`
	string             `csv:"MPN"`
	string             `csv:"Ecotax"`
	string             `csv:"Width"`
	string             `csv:"Height"`
	string             `csv:"Depth"`
	string             `csv:"Weight"`
	string             `csv:"Delivery time of in-stock products"`
	string             `csv:"Delivery time of out-of-stock products with allowed orders"`
	string             `csv:"Quantity"`
	string             `csv:"Minimal quantity"`
	string             `csv:"Low stock level"`
	string             `csv:"Receive a low stock alert by email"`
	string             `csv:"Visibility"`
	string             `csv:"Additional shipping cost"`
	string             `csv:"Unity"`
	string             `csv:"Unit price"`
	string             `csv:"Summary"`
	string             `csv:"Description"`
	string             `csv:"Tags (x,y,z...)"`
	string             `csv:"Meta title"`
	string             `csv:"Meta keywords"`
	string             `csv:"Meta description"`
	string             `csv:"URL rewritten"`
	string             `csv:"Text when in stock"`
	string             `csv:"Text when backorder allowed"`
	string             `csv:"Available for order (0 = No, 1 = Yes)"`
	string             `csv:"Product available date"`
	string             `csv:"Product creation date"`
	string             `csv:"Show price (0 = No, 1 = Yes)"`
	string             `csv:"Image URLs (x,y,z...)"`
	string             `csv:"Image alt texts (x,y,z...)"`
	string             `csv:"Delete existing images (0 = No, 1 = Yes)"`
	string             `csv:"Feature(Name:Value:Position)"`
	string             `csv:"Available online only (0 = No, 1 = Yes)"`
	string             `csv:"Condition"`
	string             `csv:"Customizable (0 = No, 1 = Yes)"`
	string             `csv:"Uploadable files (0 = No, 1 = Yes)"`
	string             `csv:"Text fields (0 = No, 1 = Yes)"`
	string             `csv:"Out of stock action"`
	string             `csv:"Virtual product"`
	string             `csv:"File URL"`
	string             `csv:"Number of allowed downloads"`
	string             `csv:"Expiration date"`
	string             `csv:"Number of days"`
	string             `csv:"ID / Name of shop"`
	string             `csv:"Advanced stock management"`
	string             `csv:"Depends On Stock"`
	string             `csv:"Warehouse"`
	string             `csv:"Acessories  (x,y,z...)"`
}

// A Shoprenter WebShop kimenet mezői

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
	Fogedzett        string `csv:"attr_values.fogedzett.hu"` // Fogedzett | Standared
	Kivitel          string `csv:"attr_values.kivitel.hu"`
	HevederSzam      string `csv:"attr_values.hevederek_szama.hu"`
	LanckerekTipus   string `csv:"attr_values.lanckerektipus.hu"` // Agyas lánckerék |  Laplánckerék
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
	Profil           string `csv:"attr_values.profil"`
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
