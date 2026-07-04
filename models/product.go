package models

const ImagesBase = "http://localhost/product-images"

const KiarusitasSzazalek = "5"

// Ezek a termékcsoportok csak rendelésre kaphatók.
var CsakRendelesre = []string{"FGL", "FGLPSZ", "FGLSPSZ", "CSCSGL", "CSCSGLPSZ", "CSCSGLSPSZ", "CSGL", "CSGLPSZ", "CSGLSPSZ"}

var Sornevek = map[string]string{"1": "Egysoros", "2": "Kétsoros", "3": "Háromsoros"}

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
	ID                       string `csv:"Product ID"`
	Active                   string `csv:"Active (0/1)"`
	Name                     string `csv:"Name *"`
	Categories               string `csv:"Categories"`
	PriceTaxExcluded         string `csv:"Price tax excluded"`
	TaxRulesID               string `csv:"Tax rules ID"`
	WholesalePrice           string `csv:"Wholesale price"`
	OnSale                   string `csv:"On sale (0/1)"`
	DiscountAmount           string `csv:"Discount amount"`
	DiscountPercent          string `csv:"Discount percent"`
	DiscountFrom             string `csv:"Discount from (yyyy-mm-dd)"`
	DiscountTo               string `csv:"Discount to (yyyy-mm-dd)"`
	Reference                string `csv:"Reference #"`
	SupplierReference        string `csv:"Supplier reference #"`
	Supplier                 string `csv:"Supplier"`
	Manufacturer             string `csv:"Manufacturer"`
	EAN13                    string `csv:"EAN13"`
	UPC                      string `csv:"UPC"`
	MPN                      string `csv:"MPN"`
	Ecotax                   string `csv:"Ecotax"`
	Width                    string `csv:"Width"`
	Height                   string `csv:"Height"`
	Depth                    string `csv:"Depth"`
	Weight                   string `csv:"Weight"`
	DeliveryTimeInStock      string `csv:"Delivery time of in-stock products"`
	DeliveryTimeOutOfStock   string `csv:"Delivery time of out-of-stock products with allowed orders"`
	Quantity                 string `csv:"Quantity"`
	MinimalQuantity          string `csv:"Minimal quantity"`
	LowStockLevel            string `csv:"Low stock level"`
	LowStockAlertEmail       string `csv:"Receive a low stock alert by email"`
	Visibility               string `csv:"Visibility"`
	AdditionalShippingCost   string `csv:"Additional shipping cost"`
	Unity                    string `csv:"Unity"`
	UnitPrice                string `csv:"Unit price"`
	Summary                  string `csv:"Summary"`
	Description              string `csv:"Description"`
	Tags                     string `csv:"Tags"`
	MetaTitle                string `csv:"Meta title"`
	MetaKeywords             string `csv:"Meta keywords"`
	MetaDescription          string `csv:"Meta description"`
	URLRewritten             string `csv:"URL rewritten"`
	TextInStock              string `csv:"Text when in stock"`
	TextBackorderAllowed     string `csv:"Text when backorder allowed"`
	AvailableForOrder        string `csv:"Available for order"`
	ProductAvailableDate     string `csv:"Product available date"`
	ProductCreationDate      string `csv:"Product creation date"`
	ShowPrice                string `csv:"Show price"`
	ImageURLs                string `csv:"Image URLs"`
	ImageAltTexts            string `csv:"Image alt texts"`
	DeleteExistingImages     string `csv:"Delete existing images"`
	Features                 string `csv:"Feature(Name:Value:Position)"`
	AvailableOnlineOnly      string `csv:"Available online only"`
	Condition                string `csv:"Condition"`
	Customizable             string `csv:"Customizable"`
	UploadableFiles          string `csv:"Uploadable files"`
	TextFields               string `csv:"Text fields"`
	OutOfStockAction         string `csv:"Out of stock action"`
	VirtualProduct           string `csv:"Virtual product"`
	FileURL                  string `csv:"File URL"`
	NumberOfAllowedDownloads string `csv:"Number of allowed downloads"`
	ExpirationDate           string `csv:"Expiration date"`
	NumberOfDays             string `csv:"Number of days"`
	ShopIDOrName             string `csv:"ID / Name of shop"`
	AdvancedStockManagement  string `csv:"Advanced stock management"`
	DependsOnStock           string `csv:"Depends On Stock"`
	Warehouse                string `csv:"Warehouse"`
	Accessories              string `csv:"Acessories"`
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
