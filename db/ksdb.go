package db

import (
	"database/sql"
	"log"
	"ws-updater/models" // Cseréld a saját projektnevedre

	_ "github.com/nakagami/firebirdsql"
)

func FetchProducts() []models.KsProduct {
	dsn := "SYSDBA:masterkey@localhost:3050//firebird/data/KSCOMPANY_AGRIALNCKFT_170203094013.KSFDB?column_name_to_lower=true&encoding=UTF8"
	db, err := sql.Open("firebirdsql", dsn)
	if err != nil {
		log.Fatalf("Sikertelen kapcsolódás: %v", err)
	}
	defer db.Close()

	sql := `
		SELECT 
		    PR."Id", 
		    PR."Code", 
		    PR."Name",
		    CAST(COALESCE(PR."Weight", 0) AS DOUBLE PRECISION) AS "Weight",
			CAST(
				COALESCE(
					(SELECT SUM(SAB."Balance") 
					FROM "StockAccountingBalance" SAB
					JOIN "StockAccountingItem" SAI ON SAI."Id" = SAB."StockAccountingItem"
					WHERE SAI."Product" = PR."Id" AND SAB."Stock" <> 2053695), 
				0) -- Ha a SUM NULL-t adna vissza, legyen 0
			AS DOUBLE PRECISION) AS "Menny",
		    QU."Name" AS "Unit",
		    CAST(MAX(CASE WHEN PRP."Price" > 0 THEN PRP."Price" ELSE NULL END) AS DOUBLE PRECISION) AS "WebPrice",
		    CAST(MIN(CASE WHEN PRP."Price" > 0 THEN PRP."Price" ELSE NULL END) AS DOUBLE PRECISION) AS "ReSellerPrice"
		FROM "Product" PR
		LEFT JOIN "QuantityUnit" QU ON QU."Id" = PR."QuantityUnit"
		LEFT JOIN "PriceRuleDetail" PRD ON PRD."Product" = PR."Id"
		LEFT JOIN "PriceRulePrice" PRP ON PRP."PriceRuleDetail" = PRD."Id" 
		    AND PRP."Currency" = 257
		    AND (PRP."ValidFrom" IS NULL OR PRP."ValidFrom" <= CURRENT_TIMESTAMP)
		    AND (PRP."ValidTo" IS NULL OR PRP."ValidTo" >= CURRENT_TIMESTAMP)
		WHERE PR."Code" LIKE 'N-%CSCS%'
		GROUP BY 1, 2, 3, 4, 6
		HAVING MAX(PRP."Price") > 0
		ORDER BY "Weight"`

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalf("Lekérdezési hiba: %q (%s)", err, sql)
	}
	defer rows.Close()

	var products []models.KsProduct
	for rows.Next() {
		var p models.KsProduct
		err := rows.Scan(
			&p.ID,
			&p.Code,
			&p.Name,
			&p.Weight,
			&p.Stock,
			&p.Unit,
			&p.WebPrice,
			&p.ReSellerPrice,
		)
		if err != nil {
			log.Fatalf("Lekérdezési hiba: %q", err)
		}
		products = append(products, p)
	}
	return products
}
