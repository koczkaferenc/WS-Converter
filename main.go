package main

import (
	"fmt"
	"os"
	"regexp"
	"ws-updater/csk"
	"ws-updater/db"
	"ws-updater/flyer"
	"ws-updater/gl"
	"ws-updater/ks"

	// "ws-updater/mgbf"
	// "ws-updater/mggl"
	"ws-updater/models"

	"bufio"

	"github.com/gocarina/gocsv"
)

var (
	// Görgőslánc
	// 1,2 és 3 soros görgősláncok
	regExpGL = regexp.MustCompile(`N-GL-[0-9]+-([0-9ABC]+)([123])$`)
	// Rozsdamentes
	regExpSSGL = regexp.MustCompile(`N-SSGL-[0-9]+-([0-9ABC]+)([0-9])$`)
	// Heavy acél görgőslánc
	regExpGL_H = regexp.MustCompile(`N-GL-[0-9]+-([0-9ABC]+)([123])_H$`)
	// Patentszemek
	// GLPSZ: 1,2 és 3 soros görgősláncok lemezes patentszemei
	regExpGLPSZ    = regexp.MustCompile(`N-GLPSZ-[0-9]+-([0-9ABC]+)([123])$`)
	regExpGLPSZ_H  = regexp.MustCompile(`N-GLPSZ-[0-9]+-([0-9ABC]+)([123])_H$`)
	regExpGLSPSZ   = regexp.MustCompile(`N-GLSPSZ-[0-9]+-([0-9ABC]+)([123])$`)
	regExpSSGLPSZ  = regexp.MustCompile(`N-SSGLPSZ-[0-9]+-([0-9ABC]+)([123])$`)
	regExpSSGLSPSZ = regexp.MustCompile(`N-SSGLSPSZ-[0-9]+-([0-9ABC]+)([123])$`)
	regExpGLHOK    = regexp.MustCompile(`N-GLHOK-[0-9]+-([0-9ABC]+)([123])$`)
	regExpGLHOK_H  = regexp.MustCompile(`N-GLHOK-[0-9]+-([0-9ABC]+)([123])_H$`)
	regExpSSGLHOK  = regexp.MustCompile(`N-SSGLHOK-[0-9]+-([0-9ABC]+)([123])$`)

	// MOFA
	regExpGLMOFA    = regexp.MustCompile(`N-GL-[0-9]+-([0-9ABC]+)([123])_MOFA(_[0-9]+)?$`)
	regExpGLPSZMOFA = regexp.MustCompile(`N-GLPSZ-[0-9]+-([0-9ABC]+)([123])_MOFA$`)
	regExpGLHOKMOFA = regexp.MustCompile(`N-GLHOK-[0-9]+-([0-9ABC]+)([123])_MOFA$`)

	// Velo
	regExpGLVELO    = regexp.MustCompile(`N-GL-[0-9]+-([0-9,\/]+x[0-9,\/]+)_VELO?$`)
	regExpGLPSZVELO = regexp.MustCompile(`N-GLPSZ-[0-9]+-([0-9,\/]+x[0-9,\/]+)_VELO?$`)
	regExpGLHOKVELO = regexp.MustCompile(`N-GLHOK-[0-9]+-([0-9,\/]+x[0-9,\/]+)_VELO?$`)

	// Csőcsapos lánc
	regExpCSCSGL    = regexp.MustCompile(`N-CSCSGL-[0-9]+-([0-9ABC]+)1$`)
	regExpCSCSGLPSZ = regexp.MustCompile(`N-CSCSGLPSZ-[0-9]+-([0-9ABC]+)1$`)

	// Párhuzamos profilú lánc
	regExpPPGL    = regexp.MustCompile(`N-PPGL-[0-9]+-([0-9ABC]+)([1-3])$`)
	regExpPPGLPSZ = regexp.MustCompile(`N-PPGLPSZ-[0-9]+-([0-9ABC]+)1$`)

	// Csapkinyomó
	regExpCSK = regexp.MustCompile(`N-CSK-[0-9]+-([0-9]+)-([0-9]+)$`)

	// Lánckerék
	regExpKS   = regexp.MustCompile(`N-KS-[0-9]+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
	regExpKS_G = regexp.MustCompile(`N-KS-[0-9]+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
	// Agyas lánckerék
	regExpKR   = regexp.MustCompile(`N-KR-[0-9]+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)
	regExpKR_G = regexp.MustCompile(`N-KR-[0-9]+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)_G$`)
	regExpGKR  = regexp.MustCompile(`N-GKR-[0-9]+-([0-9]+[A,B,C])([1-3])_Z([0-9]+)$`)

	// Flyer
	regExpFL   = regexp.MustCompile(`N-FL-[0-9]+-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$`)
	regExpFLCS = regexp.MustCompile(`N-FLCS-[0-9]+-([A-Z][A-Z])([0-9]+)([0-9])([0-9])$`)

	// ----- ITT TARTUNK ------------------------------------------------------------------------

	// Szemeslánc
	// Nem felületkezelt szemeslánc
	regExpSZL = regexp.MustCompile(`N-SZL-[0-9]+-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
	// Horganyzott szemeslánc
	regExpSZLH = regexp.MustCompile(`N-SZL-[0-9]+-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
	// Rozsdamentes szemeslánc
	regExpSSSZL = regexp.MustCompile(`N-SSSZL-[0-9]+-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
	// Szemeslánc patentszeme
	regExpSZLPSZ = regexp.MustCompile(`N-SZLPSZ-[0-9]+-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
	// Nem felületkezelt szemes bányalánc patentszem 3 mérettel
	regExpSZL3 = regexp.MustCompile(`N-SZL-[0-9]+-(\d+(?:,\d)?)x(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)

	// Mezőgazdasági lánc
	regExpMGGL = regexp.MustCompile(`N-MGGL-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_([A-Z]+)$`)

	// Boronafog
	regExpMGBF = regexp.MustCompile(`N-MGBF([A-Z]+)-([0-9]+)-([0-9]+)x([0-9]+)x([0-9]+)_M([0-9]+)_([A-Z]+)$`)

	// Hüvelyes lánc
	// Hüvelyes lánc, a TM után az osztás értéke áll
	regExpHL = regexp.MustCompile(`N-HL-[0-9]+-TM([0-9]+)$`)
	// Hüvelyes lánc, patentszeme, a TM után az osztás értéke áll
	regExpHLPSZ = regexp.MustCompile(`N-HLPSZ-[0-9]+-TM([0-9]+)$`)
)

func SaveWebProducts(products []models.WsProduct) {
	file, err := os.Create("Shoprenter.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = gocsv.MarshalFile(&products, file)
	if err != nil {
		panic(err)
	}
}

func main() {
	var webProducts []models.WsProduct
	products := db.FetchProducts()

	f, _ := os.OpenFile("Kimaradt_Termekek.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()

	processed := 0
	ignored := 0
	for _, p := range products {
		//fmt.Printf("%s\n", p.Code)
		switch {
		// GL
		case regExpGL.MatchString(p.Code),
			regExpSSGL.MatchString(p.Code),
			regExpGL_H.MatchString(p.Code),
			regExpGLMOFA.MatchString(p.Code),
			regExpGLVELO.MatchString(p.Code),
			regExpCSCSGL.MatchString(p.Code),
			regExpPPGL.MatchString(p.Code):
			webProducts = append(webProducts, gl.ProcessGl(p))
			processed++
		// GLPSZ
		case regExpGLPSZ.MatchString(p.Code),
			regExpGLPSZ.MatchString(p.Code),
			regExpGLPSZ_H.MatchString(p.Code),
			regExpGLSPSZ.MatchString(p.Code),
			regExpSSGLPSZ.MatchString(p.Code),
			regExpSSGLSPSZ.MatchString(p.Code),
			regExpGLHOK.MatchString(p.Code),
			regExpGLHOK_H.MatchString(p.Code),
			regExpSSGLHOK.MatchString(p.Code),
			regExpGLPSZVELO.MatchString(p.Code),
			regExpGLHOKVELO.MatchString(p.Code),
			regExpCSCSGLPSZ.MatchString(p.Code),
			regExpPPGLPSZ.MatchString(p.Code):
			webProducts = append(webProducts, gl.ProcessGlPsz(p))
			processed++
		// Csapkinyomók
		case regExpCSK.MatchString(p.Code):
			webProducts = append(webProducts, csk.ProcessCsk(p))
			processed++
		// TODO Boronafogak
		//#case regExpMGBF.MatchString(p.Code):
		//#	webProducts = append(webProducts, mgbf.ProcessMgbf(p))
		//#	processed++
		//#// TODO Mezőgazdasági láncok
		//#case regExpMGGL.MatchString(p.Code):
		//#	webProducts = append(webProducts, mggl.ProcessMggl(p))
		//#	processed++
		//#// Agyas és lemez lánckerekek
		case regExpKS.MatchString(p.Code),
			regExpKS_G.MatchString(p.Code),
			regExpKR.MatchString(p.Code),
			regExpKR_G.MatchString(p.Code),
			regExpGKR.MatchString(p.Code):
			webProducts = append(webProducts, ks.ProcessKs(p))
			processed++

		// Flyer
		case regExpFL.MatchString(p.Code),
			regExpFLCS.MatchString(p.Code):
			webProducts = append(webProducts, flyer.ProcessFlyer(p))
			processed++

		default:
			w.WriteString(p.Code + "\n")
			ignored++
		}
	}
	SaveWebProducts(webProducts)
	fmt.Printf("Feldolgozva: %d, Kihagyva: %d\n", processed, ignored)
}
