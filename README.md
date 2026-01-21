# WS-Converter
Agrialánc termékek konverziója webshopba

Velo láncoknál a kódban levő három értéket nem dolgozzuk fel, ez OK?

C:

isql-fb -user SYSDBA -p Informatikai '192.168.1.4/3056:C:/ProgramData/KS/FbDatabaseServer/Databases/KSCOMPANY_AGRIALNCKFT_170203094013.KSFDB'

chown firebird:firebird security2.fdb
isql-fb -u SYSDBA -role RDB$ADMIN "localhost:/eleresi/ut/security2.fdb"
gsec -database "/eleresi/ut/security2.fdb" -user SYSDBA -add SAJATUSER -pw SAJATPWD
GRANT RDB$ADMIN TO SAJATUSER;


## Adatbázis hozzáférés előállítása

1. A data könyvtárba másold be az eredeti szerver /ProgramData/KS/FbDatabaseServer/security2.fdb fájlját. Először ebbe a security2.fdb fájlba vesszük fel a WEBSHOP usert. Ehhez:
2. Indítsd a konténert.
3. Lépj be a konténerbe, és futtasd a következő parancsokat, ezzel felveszed az új usert a security2.fdb-be (WEBSHOP/AisoiDoog1gi):
    docker exec -it firebird-25 bash
    /usr/local/firebird/bin/gsec -database /firebird/data/security2.fdb -user SYSDBA -password masterkey -add WEBSHOP -pw AisoiDoog1gi
3.5 Opcionális:
   Indíts egy dbeavert és csatlakozz:
   jdbc:firebirdsql://localhost:3050//firebird/data/security2.fdb?encoding=UTF8
   Az egyetlen táblában látnod kell az új usert.


4. Állítsd le a KS adatbázis szerverét a Windowson, és másold ide a kulcssoft adatbázisát a C:/ProgramData/KS/FbDatabaseServer/Databases/ könyvtárból, ez gyakran 2 KSFDB fájl.
5. DBeaverben kapcsolódj ehhez SYSDBA és masterkey jelszóval:
   jdbc:firebirdsql://localhost:3050//firebird/data/KSCOMPANY_AGRIALNCKFT_170203094013.KSFDB?encoding=UTF8

6. Válasszd ki az adatbázist és add ki az alábbi parancsokat:
    GRANT ALL ON "Product" TO WEBSHOP;
    GRANT ALL ON "StockAccountingBalance" TO WEBSHOP;
    GRANT ALL ON "StockAccountingItem" TO WEBSHOP;
    GRANT ALL ON "QuantityUnit" TO WEBSHOP;
    GRANT ALL ON "PriceRuleDetail" TO WEBSHOP;
    GRANT ALL ON "PriceRulePrice" TO WEBSHOP;
    COMMIT;

7. Ellenőrzés, a táblák neve mellett S-t kell látni:
      SELECT 
          RDB$RELATION_NAME AS "Tábla neve", 
          RDB$PRIVILEGE AS "Jog típusa"
      FROM RDB$USER_PRIVILEGES 
      WHERE RDB$USER = 'WEBSHOP' 
        AND RDB$USER_TYPE = 8
        AND RDB$OBJECT_TYPE = 0;

Ellenőrzés:
8. Konténer indítása a mostani security2.fdb-vel:
    -v ./data/security2.fdb:/firebird/system/security2.fdb
9. Dbeaver csatlakoztatása a WEBSHOP userrel és AisoiDoog1gi jelszóval
    jdbc:firebirdsql://localhost:3050//firebird/data/KSCOMPANY_AGRIALNCKFT_170203094013.KSFDB?encoding=UTF8

10. Ellenőrizd, hogy az alábbi lekérdezés lefut-e:

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
	WHERE
		PR."Code" LIKE 'N-MGGL%' AND
		PR."OutGoingProduct" = 0
	GROUP BY 1, 2, 3, 4, 6
	HAVING MAX(PRP."Price") > 0
	ORDER BY "Weight"
