package db

import (
	"database/sql"
	"fmt"
	"ws-updater/models" // Cseréld a saját projektnevedre

	_ "github.com/nakagami/firebirdsql"
)

func GetConnection() (*sql.DB, error) {
	// A Go driver szintaxisa: user:password@host:port/path?params
	dsn := "SYSDBA:masterkey@localhost:3050//firebird/data/KSCOMPANY_AGRIALNCKFT_170203094013.KSFDB?column_name_to_lower=true&encoding=UTF8"
	return sql.Open("firebirdsql", dsn)
}

func FetchProducts(db *sql.DB) ([]models.Product, error) {
	query := `
		SELECT 
		    PR."Id", 
		    PR."Code", 
		    PR."Name",
		    CAST(PR."Weight" AS DOUBLE PRECISION) AS "Weight",
		    -- Készlet
		    CAST((SELECT SUM(SAB."Balance") 
		        FROM "StockAccountingBalance" SAB
		        JOIN "StockAccountingItem" SAI ON SAI."Id" = SAB."StockAccountingItem"
		        WHERE SAI."Product" = PR."Id" AND SAB."Stock" <> 2053695) AS DOUBLE PRECISION) AS "Menny",
		    QU."Name" AS "Unit",
		    -- Árak: Itt is kellenek a CAST-ok!
		    CAST(MAX(CASE WHEN PRP."Price" > 0 THEN PRP."Price" ELSE NULL END) AS DOUBLE PRECISION) AS "WebPrice",
		    CAST(MIN(CASE WHEN PRP."Price" > 0 THEN PRP."Price" ELSE NULL END) AS DOUBLE PRECISION) AS "ReSellerPrice"
		FROM "Product" PR
		LEFT JOIN "QuantityUnit" QU ON QU."Id" = PR."QuantityUnit"
		LEFT JOIN "PriceRuleDetail" PRD ON PRD."Product" = PR."Id"
		LEFT JOIN "PriceRulePrice" PRP ON PRP."PriceRuleDetail" = PRD."Id" 
		    AND PRP."Currency" = 257
		    AND (PRP."ValidFrom" IS NULL OR PRP."ValidFrom" <= CURRENT_TIMESTAMP)
		    AND (PRP."ValidTo" IS NULL OR PRP."ValidTo" >= CURRENT_TIMESTAMP)
		WHERE PR."Code" LIKE 'N-%'
		GROUP BY 1, 2, 3, 4, 6
		HAVING MAX(PRP."Price") > 0
		ORDER BY "Weight"`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("lekérdezési hiba: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,            // 0. index: PR."Id"
			&p.Code,          // 1. index: PR."Code"
			&p.Name,          // 2. index: PR."Name"
			&p.Weight,        // 3. index: CAST(PR."Weight"...) -> Weight
			&p.Stock,         // 4. index: CAST(SUM...) -> Menny
			&p.Unit,          // 5. index: QU."Name" -> Unit (Itt volt az eltolódás!)
			&p.WebPrice,      // 6. index: WebPrice
			&p.ReSellerPrice, // 7. index: ReSellerPrice
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
