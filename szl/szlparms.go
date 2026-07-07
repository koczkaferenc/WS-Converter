package szl

import "regexp"

var (
	// Nem felületkezelt szemeslánc
	regExpSZL = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
	// Horganyzott szemeslánc
	//regExpSZLH = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
	// Rozsdamentes szemeslánc
	//regExpSSSZL = regexp.MustCompile(`N-(SSSZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
	// Szemeslánc patentszeme
	//regExpSZLPSZ = regexp.MustCompile(`N-(SZLPSZ)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
	// Nem felületkezelt szemes bányalánc patentszem 3 mérettel
	//regExpSZL3 = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
)
