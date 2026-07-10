package szl

import "regexp"

var (
	// Horganyzott vagy natúr szemeslánc
	// regExpSZL = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
	// regExpSZL = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)(?:_(H))?$`)
	regExpSZL = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)(?:_(H))?$`)
	// Ez nincs külön, összevontuk, Horganyzott szemeslánc
	// regExpSZLH = regexp.MustCompil(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
	// Rozsdamentes szemeslánc
	regExpSSSZL = regexp.MustCompile(`N-(SSSZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)_(H)$`)
	// Szemeslánc patentszeme
	regExpSZLPSZ = regexp.MustCompile(`N-(SZLPSZ)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
	// Nem felületkezelt szemes bányalánc patentszem 3 mérettel
	regExpSZL3 = regexp.MustCompile(`N-(SZL)-([0-9]+)-(\d+(?:,\d)?)x(\d+(?:,\d)?)x(\d+(?:,\d)?)$`)
)
