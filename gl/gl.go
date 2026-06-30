package gl

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"ws-updater/models"
)

var GLParms = map[string]map[string]string{
	"08A1":  {"Osztas": "12,7,7", "Belsoheveder": "7,85", "Csapatmero": "3,96", "Gorgoatmero": "7,95", "Csaphossz": "21,7"},
	"10A1":  {"Osztas": "15,875", "Belsoheveder": "9,4", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "25,9"},
	"12A1":  {"Osztas": "19,05", "Belsoheveder": "12,57", "Csapatmero": "5,94", "Gorgoatmero": "11,91", "Csaphossz": "31,5"},
	"16A1":  {"Osztas": "25,4,4", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": "38,9"},
	"20A1":  {"Osztas": "31,75", "Belsoheveder": "18,9", "Csapatmero": "9,53", "Gorgoatmero": "19,05", "Csaphossz": "45,2"},
	"24A1":  {"Osztas": "38,1", "Belsoheveder": "25,22", "Csapatmero": "11,1", "Gorgoatmero": "22,23", "Csaphossz": "55,5"},
	"28A1":  {"Osztas": "44,45", "Belsoheveder": "25,22", "Csapatmero": "12,7", "Gorgoatmero": "25,4", "Csaphossz": "59,3"},
	"32A1":  {"Osztas": "50,8", "Belsoheveder": "31,55", "Csapatmero": "14,27", "Gorgoatmero": "28,58", "Csaphossz": "69,6"},
	"36A1":  {"Osztas": "57,15", "Belsoheveder": "35,48", "Csapatmero": "17,46", "Gorgoatmero": "35,71", "Csaphossz": ""},
	"40A1":  {"Osztas": "63,5", "Belsoheveder": "37,85", "Csapatmero": "19,85", "Gorgoatmero": "39,68", "Csaphossz": "85,4"},
	"48A1":  {"Osztas": "76,2", "Belsoheveder": "47,35", "Csapatmero": "23,81", "Gorgoatmero": "47,63", "Csaphossz": "103,1"},
	"08A2":  {"Osztas": "12,7,7", "Belsoheveder": "7,85", "Csapatmero": "3,96", "Gorgoatmero": "7,95", "Csaphossz": "36,2"},
	"10A2":  {"Osztas": "15,875", "Belsoheveder": "9,4", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "44"},
	"12A2":  {"Osztas": "19,05", "Belsoheveder": "12,57", "Csapatmero": "5,94", "Gorgoatmero": "11,91", "Csaphossz": "44,4"},
	"16A2":  {"Osztas": "25,4,4", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": "68,1"},
	"20A2":  {"Osztas": "31,75", "Belsoheveder": "18,9", "Csapatmero": "9,53", "Gorgoatmero": "19,05", "Csaphossz": "81,2"},
	"24A2":  {"Osztas": "38,1", "Belsoheveder": "25,22", "Csapatmero": "11,1", "Gorgoatmero": "22,23", "Csaphossz": "100,9"},
	"28A2":  {"Osztas": "44,45", "Belsoheveder": "25,22", "Csapatmero": "12,7", "Gorgoatmero": "25,4", "Csaphossz": "108,2"},
	"32A2":  {"Osztas": "50,8", "Belsoheveder": "31,55", "Csapatmero": "14,27", "Gorgoatmero": "28,58", "Csaphossz": "128,2"},
	"36A2":  {"Osztas": "57,15", "Belsoheveder": "35,48", "Csapatmero": "17,46", "Gorgoatmero": "35,71", "Csaphossz": ""},
	"40A2":  {"Osztas": "63,5", "Belsoheveder": "37,85", "Csapatmero": "19,85", "Gorgoatmero": "39,68", "Csaphossz": "157"},
	"48A2":  {"Osztas": "76,2", "Belsoheveder": "47,35", "Csapatmero": "23,81", "Gorgoatmero": "47,63", "Csaphossz": "191"},
	"08A3":  {"Osztas": "12,7,7", "Belsoheveder": "7,85", "Csapatmero": "3,96", "Gorgoatmero": "7,95", "Csaphossz": "50,6"},
	"10A3":  {"Osztas": "15,875", "Belsoheveder": "9,4", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "62"},
	"12A3":  {"Osztas": "19,05", "Belsoheveder": "12,57", "Csapatmero": "5,94", "Gorgoatmero": "11,91", "Csaphossz": "77,2"},
	"16A3":  {"Osztas": "25,4,4", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": "97,1"},
	"20A3":  {"Osztas": "31,75", "Belsoheveder": "18,9", "Csapatmero": "9,53", "Gorgoatmero": "19,05", "Csaphossz": "117,2"},
	"24A3":  {"Osztas": "38,1", "Belsoheveder": "25,22", "Csapatmero": "11,1", "Gorgoatmero": "22,23", "Csaphossz": "146,4"},
	"28A3":  {"Osztas": "44,45", "Belsoheveder": "25,22", "Csapatmero": "12,7", "Gorgoatmero": "25,4", "Csaphossz": "157"},
	"32A3":  {"Osztas": "50,8", "Belsoheveder": "31,55", "Csapatmero": "14,27", "Gorgoatmero": "28,58", "Csaphossz": "186,7"},
	"36A3":  {"Osztas": "57,15", "Belsoheveder": "35,48", "Csapatmero": "17,46", "Gorgoatmero": "35,71", "Csaphossz": ""},
	"40A3":  {"Osztas": "63,5", "Belsoheveder": "37,85", "Csapatmero": "19,85", "Gorgoatmero": "39,68", "Csaphossz": "228,5"},
	"48A3":  {"Osztas": "76,2", "Belsoheveder": "47,35", "Csapatmero": "23,81", "Gorgoatmero": "47,63", "Csaphossz": "278,8"},
	"08AH1": {"Osztas": "12,7,7", "Belsoheveder": "7,85", "Csapatmero": "3,96", "Gorgoatmero": "7,95", "Csaphossz": "19,9"},
	"10AH1": {"Osztas": "15,875", "Belsoheveder": "9,4", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "24,3"},
	"12AH1": {"Osztas": "19,05", "Belsoheveder": "12,57", "Csapatmero": "5,94", "Gorgoatmero": "11,91", "Csaphossz": "31"},
	"16AH1": {"Osztas": "25,4,4", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": "42,4"},
	"20AH1": {"Osztas": "31,75", "Belsoheveder": "18,9", "Csapatmero": "9,53", "Gorgoatmero": "19,05", "Csaphossz": "50,6"},
	"24AH1": {"Osztas": "38,1", "Belsoheveder": "25,22", "Csapatmero": "11,1", "Gorgoatmero": "22,23", "Csaphossz": "61,4"},
	"28AH1": {"Osztas": "44,45", "Belsoheveder": "25,22", "Csapatmero": "12,7", "Gorgoatmero": "25,4", "Csaphossz": "66,1"},
	"32AH1": {"Osztas": "50,8", "Belsoheveder": "31,55", "Csapatmero": "14,27", "Gorgoatmero": "28,58", "Csaphossz": "75,4"},
	"36AH1": {"Osztas": "57,15", "Belsoheveder": "35,48", "Csapatmero": "17,46", "Gorgoatmero": "35,71", "Csaphossz": ""},
	"40AH1": {"Osztas": "63,5", "Belsoheveder": "37,85", "Csapatmero": "19,85", "Gorgoatmero": "39,68", "Csaphossz": "95,6"},
	"48AH1": {"Osztas": "76,2", "Belsoheveder": "47,35", "Csapatmero": "23,81", "Gorgoatmero": "47,63", "Csaphossz": "115,5"},
	"12AH2": {"Osztas": "19,05", "Belsoheveder": "12,57", "Csapatmero": "5,94", "Gorgoatmero": "11,91", "Csaphossz": ""},
	"16AH2": {"Osztas": "25,4", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": ""},
	"20AH2": {"Osztas": "31,75", "Belsoheveder": "18,9", "Csapatmero": "9,53", "Gorgoatmero": "19,05", "Csaphossz": ""},
	"24AH2": {"Osztas": "38,1", "Belsoheveder": "25,22", "Csapatmero": "11,1", "Gorgoatmero": "22,23", "Csaphossz": ""},
	"28AH2": {"Osztas": "44,45", "Belsoheveder": "25,22", "Csapatmero": "12,7", "Gorgoatmero": "25,4", "Csaphossz": ""},
	"32AH2": {"Osztas": "50,8", "Belsoheveder": "31,55", "Csapatmero": "14,27", "Gorgoatmero": "28,58", "Csaphossz": ""},
	"40AH2": {"Osztas": "63,5", "Belsoheveder": "37,85", "Csapatmero": "19,85", "Gorgoatmero": "39,68", "Csaphossz": ""},
	"12AH3": {"Osztas": "19,05", "Belsoheveder": "12,57", "Csapatmero": "5,94", "Gorgoatmero": "11,91", "Csaphossz": ""},
	"16AH3": {"Osztas": "25,4", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": ""},
	"20AH3": {"Osztas": "31,75", "Belsoheveder": "18,9", "Csapatmero": "9,53", "Gorgoatmero": "19,05", "Csaphossz": ""},
	"24AH3": {"Osztas": "38,1", "Belsoheveder": "25,22", "Csapatmero": "11,1", "Gorgoatmero": "22,23", "Csaphossz": ""},
	"28AH3": {"Osztas": "44,45", "Belsoheveder": "25,22", "Csapatmero": "12,7", "Gorgoatmero": "25,4", "Csaphossz": ""},
	"32AH3": {"Osztas": "50,8", "Belsoheveder": "31,55", "Csapatmero": "14,27", "Gorgoatmero": "28,58", "Csaphossz": ""},
	"40AH3": {"Osztas": "63,5", "Belsoheveder": "37,85", "Csapatmero": "19,85", "Gorgoatmero": "39,68", "Csaphossz": ""},

	"04B1": {"Osztas": "6", "Belsoheveder": "2,8", "Csapatmero": "1,85", "Gorgoatmero": "4", "Csaphossz": "7,8"},
	"05B1": {"Osztas": "8", "Belsoheveder": "3", "Csapatmero": "2,31", "Gorgoatmero": "5", "Csaphossz": "11,7"},
	"06B1": {"Osztas": "9,525", "Belsoheveder": "5,72", "Csapatmero": "3,28", "Gorgoatmero": "6,35", "Csaphossz": "16,8"},
	"08B1": {"Osztas": "12,7", "Belsoheveder": "7,75", "Csapatmero": "4,45", "Gorgoatmero": "8,51", "Csaphossz": "20,9"},
	"10B1": {"Osztas": "15,875", "Belsoheveder": "9,65", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "23,7"},
	"12B1": {"Osztas": "19,05", "Belsoheveder": "11,68", "Csapatmero": "5,72", "Gorgoatmero": "12,07", "Csaphossz": "27,3"},
	"16B1": {"Osztas": "25,4", "Belsoheveder": "17,02", "Csapatmero": "8,28", "Gorgoatmero": "15,88", "Csaphossz": "41,5"},
	"20B1": {"Osztas": "31,75", "Belsoheveder": "19,56", "Csapatmero": "10,19", "Gorgoatmero": "19,05", "Csaphossz": "46"},
	"24B1": {"Osztas": "38", "Belsoheveder": "25,4", "Csapatmero": "14,63", "Gorgoatmero": "25,4", "Csaphossz": "58,5"},
	"28B1": {"Osztas": "44,45", "Belsoheveder": "30,99", "Csapatmero": "15,9", "Gorgoatmero": "27,94", "Csaphossz": "69,6"},
	"32B1": {"Osztas": "50", "Belsoheveder": "30,99", "Csapatmero": "17,81", "Gorgoatmero": "29,21", "Csaphossz": "73,1"},
	"40B1": {"Osztas": "63", "Belsoheveder": "38,1", "Csapatmero": "22,89", "Gorgoatmero": "39,37", "Csaphossz": "86,3"},
	"48B1": {"Osztas": "76", "Belsoheveder": "45,72", "Csapatmero": "29,24", "Gorgoatmero": "48,26", "Csaphossz": "107,9"},
	"56B1": {"Osztas": "88", "Belsoheveder": "53,34", "Csapatmero": "34,32", "Gorgoatmero": "53,98", "Csaphossz": "137"},
	"64B1": {"Osztas": "101", "Belsoheveder": "60,96", "Csapatmero": "39,4", "Gorgoatmero": "63,5", "Csaphossz": "138,5"},
	"72B1": {"Osztas": "114", "Belsoheveder": "68,58", "Csapatmero": "44,48", "Gorgoatmero": "72,39", "Csaphossz": "156,4"},

	"05B2": {"Osztas": "8", "Belsoheveder": "3", "Csapatmero": "2,31", "Gorgoatmero": "5", "Csaphossz": "17,4"},
	"06B2": {"Osztas": "9,525", "Belsoheveder": "5,72", "Csapatmero": "3,28", "Gorgoatmero": "6,35", "Csaphossz": "27,1"},
	"08B2": {"Osztas": "12", "Belsoheveder": "7,75", "Csapatmero": "4,45", "Gorgoatmero": "8,51", "Csaphossz": "34,9"},
	"10B2": {"Osztas": "15,875", "Belsoheveder": "9,65", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "40,3"},
	"12B2": {"Osztas": "19,05", "Belsoheveder": "11,68", "Csapatmero": "5,72", "Gorgoatmero": "12,07", "Csaphossz": "46,8"},
	"16B2": {"Osztas": "25", "Belsoheveder": "17,02", "Csapatmero": "8,28", "Gorgoatmero": "15,88", "Csaphossz": "73,4"},
	"20B2": {"Osztas": "31,75", "Belsoheveder": "19,56", "Csapatmero": "10,19", "Gorgoatmero": "19,05", "Csaphossz": "82,5"},
	"24B2": {"Osztas": "38", "Belsoheveder": "25,4", "Csapatmero": "14,63", "Gorgoatmero": "25,4", "Csaphossz": "106,9"},
	"28B2": {"Osztas": "44,45", "Belsoheveder": "30,99", "Csapatmero": "15,9", "Gorgoatmero": "27,94", "Csaphossz": "129,2"},
	"32B2": {"Osztas": "50", "Belsoheveder": "30,99", "Csapatmero": "17,81", "Gorgoatmero": "29,21", "Csaphossz": "131,7"},
	"40B2": {"Osztas": "63", "Belsoheveder": "38,1", "Csapatmero": "22,89", "Gorgoatmero": "39,37", "Csaphossz": "158,6"},
	"48B2": {"Osztas": "76", "Belsoheveder": "45,72", "Csapatmero": "29,24", "Gorgoatmero": "48,26", "Csaphossz": "199,1"},
	"56B2": {"Osztas": "88", "Belsoheveder": "53,34", "Csapatmero": "34,32", "Gorgoatmero": "53,98", "Csaphossz": "243,6"},
	"64B2": {"Osztas": "101", "Belsoheveder": "60,96", "Csapatmero": "39,4", "Gorgoatmero": "63,5", "Csaphossz": "258,4"},
	"72B2": {"Osztas": "114", "Belsoheveder": "68,58", "Csapatmero": "44,48", "Gorgoatmero": "72,39", "Csaphossz": "292,7"},

	"05B3": {"Osztas": "8", "Belsoheveder": "3", "Csapatmero": "2,31", "Gorgoatmero": "5", "Csaphossz": "23"},
	"06B3": {"Osztas": "9,525", "Belsoheveder": "5,72", "Csapatmero": "3,28", "Gorgoatmero": "6,35", "Csaphossz": "37,3"},
	"08B3": {"Osztas": "12", "Belsoheveder": "7,75", "Csapatmero": "4,45", "Gorgoatmero": "8,51", "Csaphossz": "48,8"},
	"10B3": {"Osztas": "15,875", "Belsoheveder": "9,65", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "56,9"},
	"12B3": {"Osztas": "19,05", "Belsoheveder": "11,68", "Csapatmero": "5,72", "Gorgoatmero": "12,07", "Csaphossz": "66,3"},
	"16B3": {"Osztas": "25", "Belsoheveder": "17,02", "Csapatmero": "8,28", "Gorgoatmero": "15,88", "Csaphossz": "105,3"},
	"20B3": {"Osztas": "31,75", "Belsoheveder": "19,56", "Csapatmero": "10,19", "Gorgoatmero": "19,05", "Csaphossz": "118,9"},
	"24B3": {"Osztas": "38", "Belsoheveder": "25,4", "Csapatmero": "14,63", "Gorgoatmero": "25,4", "Csaphossz": "155,2"},
	"28B3": {"Osztas": "44,45", "Belsoheveder": "30,99", "Csapatmero": "15,9", "Gorgoatmero": "27,94", "Csaphossz": "188,8"},
	"32B3": {"Osztas": "50", "Belsoheveder": "30,99", "Csapatmero": "17,81", "Gorgoatmero": "29,21", "Csaphossz": "190,2"},
	"40B3": {"Osztas": "63", "Belsoheveder": "38,1", "Csapatmero": "22,89", "Gorgoatmero": "39,37", "Csaphossz": "230,9"},
	"48B3": {"Osztas": "76", "Belsoheveder": "45,72", "Csapatmero": "29,24", "Gorgoatmero": "48,26", "Csaphossz": "293,3"},
	"56B3": {"Osztas": "88", "Belsoheveder": "53,34", "Csapatmero": "34,32", "Gorgoatmero": "53,98", "Csaphossz": "350,2"},
	"64B3": {"Osztas": "101", "Belsoheveder": "60,96", "Csapatmero": "39,4", "Gorgoatmero": "63,5", "Csaphossz": "378,3"},
	"72B3": {"Osztas": "114", "Belsoheveder": "68,58", "Csapatmero": "44,48", "Gorgoatmero": "72,39", "Csaphossz": "429"},

	"04BH":   {"Osztas": "6", "Belsoheveder": "2,8", "Csapatmero": "1,85", "Gorgoatmero": "4", "Csaphossz": "9,4"},
	"06BH":   {"Osztas": "9,525", "Belsoheveder": "5,72", "Csapatmero": "3,58", "Gorgoatmero": "6,35", "Csaphossz": "15,4"},
	"08BH":   {"Osztas": "12", "Belsoheveder": "7,85", "Csapatmero": "4,45", "Gorgoatmero": "8,51", "Csaphossz": "19,9"},
	"10BH":   {"Osztas": "15,875", "Belsoheveder": "9,65", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "21,6"},
	"12BH":   {"Osztas": "19,05", "Belsoheveder": "11,68", "Csapatmero": "5,94", "Gorgoatmero": "12,07", "Csaphossz": "26,8"},
	"12BV":   {"Osztas": "19,05", "Belsoheveder": "11,68", "Csapatmero": "6,1", "Gorgoatmero": "12,07", "Csaphossz": "26,5"},
	"12BHF2": {"Osztas": "19,05", "Belsoheveder": "13,5", "Csapatmero": "5,72", "Gorgoatmero": "12,07", "Csaphossz": "30,3"},
	"16BH":   {"Osztas": "25", "Belsoheveder": "17,02", "Csapatmero": "8,9", "Gorgoatmero": "15,88", "Csaphossz": "38,9"},
	"24BH":   {"Osztas": "38", "Belsoheveder": "25,4", "Csapatmero": "14,63", "Gorgoatmero": "25,4", "Csaphossz": "63,4"},
	"24BHF2": {"Osztas": "38", "Belsoheveder": "25,4", "Csapatmero": "14,63", "Gorgoatmero": "25,4", "Csaphossz": "62,2"},

	"208A1":  {"Osztas": "25,4", "Belsoheveder": "7,85", "Csapatmero": "3,96", "Gorgoatmero": "7,95", "Csaphossz": "17,8"},
	"208B1":  {"Osztas": "25,4", "Belsoheveder": "7,75", "Csapatmero": "4,45", "Gorgoatmero": "8,51", "Csaphossz": "18,2"},
	"210A1":  {"Osztas": "31,75", "Belsoheveder": "9,4", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "22,2"},
	"210B1":  {"Osztas": "31,75", "Belsoheveder": "9,65", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "20,9"},
	"212A1":  {"Osztas": "38,1", "Belsoheveder": "12,57", "Csapatmero": "5,94", "Gorgoatmero": "11,91", "Csaphossz": "27,7"},
	"212B1":  {"Osztas": "38,1", "Belsoheveder": "11,68", "Csapatmero": "5,72", "Gorgoatmero": "12,07", "Csaphossz": "25,2"},
	"216A1":  {"Osztas": "50,8", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": "36,5"},
	"216A1H": {"Osztas": "50,8", "Belsoheveder": "15,75", "Csapatmero": "7,92", "Gorgoatmero": "15,88", "Csaphossz": "39,4"},
	"216B1":  {"Osztas": "50,8", "Belsoheveder": "17,02", "Csapatmero": "8,28", "Gorgoatmero": "15,88", "Csaphossz": "41,5"},
	"220A1":  {"Osztas": "63,5", "Belsoheveder": "18,9", "Csapatmero": "9,53", "Gorgoatmero": "19,05", "Csaphossz": "44,7"},
	"220B1":  {"Osztas": "63,5", "Belsoheveder": "19,56", "Csapatmero": "10,19", "Gorgoatmero": "19,05", "Csaphossz": "46"},
	"224A1":  {"Osztas": "76,2", "Belsoheveder": "25,22", "Csapatmero": "11,1", "Gorgoatmero": "22,23", "Csaphossz": "54,3"},
	"224B1":  {"Osztas": "76,2", "Belsoheveder": "25,4", "Csapatmero": "14,63", "Gorgoatmero": "25,4", "Csaphossz": "58,5"},
	"228B1":  {"Osztas": "88,9", "Belsoheveder": "30,99", "Csapatmero": "15,9", "Gorgoatmero": "27,94", "Csaphossz": "69,5"},
	"232B1":  {"Osztas": "101,6", "Belsoheveder": "30,99", "Csapatmero": "17,81", "Gorgoatmero": "29,21", "Csaphossz": "71"},

	// Manuális úton felvéve
	"04C1":  {"Osztas": "6,35", "Belsoheveder": "3,18", "Csapatmero": "2,31", "Gorgoatmero": "3,3", "Csaphossz": "8,4"},
	"031":   {"Osztas": "5", "Belsoheveder": "2,5", "Csapatmero": "1,49", "Gorgoatmero": "3,2", "Csaphossz": "9,9"},
	"0811":  {"Osztas": "12,7", "Belsoheveder": "3,3", "Csapatmero": "3,66", "Gorgoatmero": "7,75", "Csaphossz": "11,7"},
	"062C1": {"Osztas": "9,525", "Belsoheveder": "9,52", "Csapatmero": "4,18", "Gorgoatmero": "6", "Csaphossz": ""},
	"06C1":  {"Osztas": "9,525", "Belsoheveder": "4,77", "Csapatmero": "3,58", "Gorgoatmero": "5,08", "Csaphossz": "13,3"},
	"06C2":  {"Osztas": "9,525", "Belsoheveder": "4,55", "Csapatmero": "3,58", "Gorgoatmero": "5,08", "Csaphossz": "23,45"},
	"0831":  {"Osztas": "12,7", "Belsoheveder": "4,88", "Csapatmero": "4,09", "Gorgoatmero": "7,75", "Csaphossz": "12,9"},
	"0861":  {"Osztas": "12,7", "Belsoheveder": "5,3", "Csapatmero": "4,45", "Gorgoatmero": "8,51", "Csaphossz": "15,9"},
	"1011":  {"Osztas": "15,88", "Belsoheveder": "6,48", "Csapatmero": "5,08", "Gorgoatmero": "10,16", "Csaphossz": "17,2"},
	"411":   {"Osztas": "12,7", "Belsoheveder": "6,25", "Csapatmero": "3,58", "Gorgoatmero": "7,77", "Csaphossz": "15"},
	"04B2":  {"Osztas": "", "Belsoheveder": "", "Csapatmero": "", "Gorgoatmero": "", "Csaphossz": ""},

	// VELO
	"1/2x3/16": {"Osztas": "12,7", "Belsoheveder": "4,88", "Csapatmero": "7,75", "Gorgoatmero": "3,66", "Csaphossz": "12,3"},
	"1/2x1/8":  {"Osztas": "12,7", "Belsoheveder": "3,3", "Csapatmero": "7,75", "Gorgoatmero": "3,66", "Csaphossz": "10,2"},
	// MOFA
	"0841_MOFA": {"Osztas": "12,7", "Belsoheveder": "4,88", "Csapatmero": "4,18", "Gorgoatmero": "7,75", "Csaphossz": "14,3"},
	"0851_MOFA": {"Osztas": "12,7", "Belsoheveder": "6,4", "Csapatmero": "4,18", "Gorgoatmero": "7,75", "Csaphossz": "15,9"},
}

//*******************************************************************
// Görgősláncok
//*******************************************************************

func ProcessGl(p models.KsProduct, prodCodes *[]string) models.PsProduct {
	var (
		// Alap görgőslánc
		regExpGL = regexp.MustCompile(`N-(GL)-([0-9]+)-([0-9ABC]+)([123])$`)
		// Rozsdamentes görgőslánc
		regExpSSGL = regexp.MustCompile(`N-(SSGL)-([0-9]+)-([0-9ABC]+)([0-9])$`)
		// Heavy görgőslánc
		regExpGL_H = regexp.MustCompile(`N-(GL)-([0-9]+)-([0-9ABC]+)([123])_H$`)
		// Mofa görgőslánc
		regExpGLMOFA = regexp.MustCompile(`N-(GL)-([0-9]+)-([0-9ABC]+)([123])_MOFA(_[0-9]+)?$`)
		// VELO görgőslánc
		regExpGLVELO = regexp.MustCompile(`N-(GL)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$`)
		// CSCS
		regExpCSCSGL = regexp.MustCompile(`N-(CSCSGL)-([0-9]+)-([0-9ABC]+)1$`)
		// Párhuzamos profilú görgősláncok
		regExpPPGL = regexp.MustCompile(`N-(PPGL)-([0-9]+)-([0-9ABC]+)([1-3])$`)

		match          []string
		family         string
		manufacturerId string
		sorokszama     string
		productType    string
		pStr           string
	)
	psp := models.PsProduct{}
	features := map[string]string{
		"Anyag":                 "", // Acél |
		"Erősített":             "", // Igen|Nem
		"Osztás":                "", // Típus: Agyas lánckerék | Laplánckerék
		"Belső hevedertávolság": "", // Fogedzett: Igen | Nem
		"Csapátmérő":            "", // Fogak száma
		"Görgőátmérő":           "", // mm
		"Csaphossz":             "", // mm
		"Csaptípus":             "", // Tömör|Csőcsap
		"Szemforma":             "", // Piskóta|Egyenes|Hajlított
		"Kivitel":               "", // Egysoros | kétsoros | hármosoros (sorokszama alapján képezve)
	}

	psp.ID = ""            // Az ID üres, cikkszámokkal dogozunk
	psp.Reference = p.Code // Ez a cikkszám, ID helyett használjuk
	psp.Active = "1"
	psp.PriceTaxExcluded = fmt.Sprintf("%.0f", p.WebPrice)
	psp.UnitPrice = fmt.Sprintf("%.0f", p.WebPrice) // TODO
	psp.TaxRulesID = "1"                            // ÁFA kulcs 27%
	psp.Quantity = fmt.Sprintf("%.1f", p.Stock)     // Mennyiség
	psp.Weight = fmt.Sprintf("%.1f", p.Weight)
	psp.Unity = "m"
	psp.TextInStock = "m"          // méter
	psp.TextBackorderAllowed = "m" // méter
	psp.Categories = "Láncok/Görgősláncok"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// Görgőslánc: N-GL-5-24B3
	// N-(GL)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[productType+sorokszama]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Standard Görgőslánc", productType, sorokszama)
		psp.Tags = "Standard,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,standard,görgőslánc"
		//fmt.Printf("%s: (%s) -> %s\n", p.Code, productType, features["Csaphossz"])

	}

	// Rozsdamentes görgőslánc
	// N-(SSGL)-([0-9]+)-([0-9ABC]+)([0-9])$
	match = regExpSSGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Rozsdamentes acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Görgőslánc", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,görgőslánc"

	}

	// Heavy görgőslánc
	// N-(GL)-([0-9]+)-([0-9ABC]+)([123])_H$
	match = regExpGL_H.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Igen"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Erősített Görgőslánc", productType, sorokszama)
		psp.Tags = "Erősített,Heavy,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,erősített,heavy,görgőslánc"
	}

	// Mofa görgőslánc
	// N-(GL)-([0-9]+)-([0-9ABC]+)([123])_MOFA(_[0-9]+)?$
	match = regExpGLMOFA.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s MOFA Görgőslánc", productType, sorokszama)
		psp.Tags = "MOFA,Görgőslánc"
		psp.MetaKeywords = "Mofa,görgőslánc"
	}

	// Velo görgőslánc
	// N-(GL)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$
	match = regExpGLVELO.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = "1" // Velo mindig egysoros

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Velo Görgőslánc", productType, sorokszama)
		psp.Tags = "Standard,Görgőslánc"
		psp.MetaKeywords = "Velo,görgőslánc"
	}

	// Csőcsapos
	// N-(CSCSGL)-([0-9]+)-([0-9ABC]+)1$
	match = regExpCSCSGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = "1" // Mindig egysoros

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Csőcsapos"
		features["Szemforma"] = "Piskóta"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Csőcsapos Standard Görgőslánc", productType, sorokszama)
		psp.Tags = "Csőcsapos,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,csőcsapos,görgőslánc"
	}

	// Párhuzamos profilú
	// N-(PPGL)-([0-9]+)-([0-9ABC]+)([1-3])$
	match = regExpPPGL.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Párhuzamos profilú"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Párhuzamos Profilú Standard Görgőslánc", productType, sorokszama)
		psp.Tags = "Standard,Párhuzamos profilú,Görgőslánc"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,párhuzamos profilú,standard,görgőslánc"
	}

	// Gyártó beállítása
	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

	kivitelTmp := "Standard"
	if features["Erősített"] == "Igen" {
		kivitelTmp = "Erősített"
	}
	psp.Summary = fmt.Sprintf("%s , %s soros, %s mm osztású, %s mm belső hevedertávolságú, %s mm görgőátmérőjű %s %s görgőslánc.", psp.Manufacturer, sorokszama, features["Osztás"], features["Belső hevedertávolság"], features["Görgőátmérő"], strings.ToLower(kivitelTmp), strings.ToLower(features["Anyag"]))
	// Ebben a fázisában kell beállítani és nem lehet pont a végén.
	psp.MetaDescription = strings.TrimRight(psp.Summary, ".")
	psp.Summary += models.Zaradek

	psp.MetaTitle = psp.Name
	psp.URLRewritten = p.Code

	// Képek előállítása
	if psp.ImageURLs == "" {
		psp.ImageURLs = fmt.Sprintf(
			"%s/N-%s-%s.png,%s/D-%s-%s.png",
			models.ImagesBase, family, sorokszama,
			models.ImagesBase, family, sorokszama)
		psp.ImageAltTexts = psp.Name
	}

	// Speciális tulajdonságok beállítása
	psp.Features = models.MkFeaturesList(features)

	// Kapcsolódó termékek
	psp.Accessories = ""
	// TODO Ideiglenesen kivesszük, mert ha nem ltezik a termék, nagyon lelassul
	//rgxStr := fmt.Sprintf(`^N-GL-[0-9]+-%s%s.*`, productType, sorokszama)
	//psp.Accessories += models.getRelatedProductIds(rgxStr, prodCodes)

	// Rendelhető?
	psp.AvailableForOrder = "1"
	if slices.Contains(models.CsakRendelesre, family) {
		psp.AvailableForOrder = "0"
	}

	return psp
}

// *******************************************************************
// Görgősláncok patentszemei
// *******************************************************************
func ProcessGlPsz(p models.KsProduct, prodCodes *[]string) models.PsProduct {
	var (
		// Normál
		regExpGLPSZ = regexp.MustCompile(`N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])$`)
		// Heavy
		regExpGLPSZ_H = regexp.MustCompile(`N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])_H$`)
		// Sasszeges
		regExpGLSPSZ = regexp.MustCompile(`N-(GLSPSZ)-([0-9]+)-([0-9ABC]+)([123])$`)
		// Rozsdamentes, rugós lemezes
		regExpSSGLPSZ = regexp.MustCompile(`N-(SSGLPSZ)-([0-9]+)-([0-9ABC]+)([123])$`)
		// Rozsdamentes, sasszeges
		regExpSSGLSPSZ = regexp.MustCompile(`N-(SSGLSPSZ)-([0-9]+)-([0-9ABC]+)([123])$`)
		// Sasszeges hajlított patentszem (a hajlított mindig sasszeges)
		regExpGLHOK = regexp.MustCompile(`N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])$`)
		// Erősített görgősláncok rugós sasszeges hajlított patentszem
		regExpGLHOK_H = regexp.MustCompile(`N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])_H$`)
		// Rozsdamentes görgősláncok rugós sasszeges hajlított patentszemei
		regExpSSGLHOK = regexp.MustCompile(`N-(SSGLHOK)-([0-9]+)-([0-9ABC]+)([123])$`)
		// GLPSZ MOFA
		regExpGLPSZMOFA = regexp.MustCompile(`N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])_MOFA$`)
		// Hajlított MOFA
		regExpGLHOKMOFA = regexp.MustCompile(`N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])_MOFA$`)
		// VELO Patentszem
		regExpGLPSZVELO = regexp.MustCompile(`N-(GLPSZ)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$`)
		// Hajlított VELO
		regExpGLHOKVELO = regexp.MustCompile(`N-(GLHOK)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$`)
		// Csőcsapos egyenes patentszem
		regExpCSCSGLPSZ = regexp.MustCompile(`N-(CSCSGLPSZ)-([0-9]+)-([0-9ABC]+)1$`)
		// Párhuzamos profilú lánc patentszeme
		regExpPPGLPSZ = regexp.MustCompile(`N-(PPGLPSZ)-([0-9])+-([0-9ABC]+)1$`)

		match          []string
		family         string
		manufacturerId string
		sorokszama     string
		productType    string
		pStr           string
	)
	psp := models.PsProduct{}
	features := map[string]string{
		"Anyag":                 "", // Acél |
		"Erősített":             "", // Igen|Nem
		"Osztás":                "", // Típus: Agyas lánckerék | Laplánckerék
		"Belső hevedertávolság": "", // Fogedzett: Igen | Nem
		"Csapátmérő":            "", // Fogak száma
		"Görgőátmérő":           "", // mm
		"Csaphossz":             "", // mm
		"Csaptípus":             "", // Tömör|Csőcsapos
		"Szemforma":             "", // Piskóta|Egyenes|Hajlított
		"Kivitel":               "", // Egysoros | kétsoros | hármosoros (sorokszama alapján képezve)
		"Rögzítés":              "", // Rugós lemezes | Sasszeges
	}

	// Általános tulajdonságok
	// w.SKU = p.Code
	// w.Anyag = "Acél"
	// w.Kivitel = "Normál"    // Normál | Heavy
	// w.Rogzites = "Lemezes"  // Lemezes | Sasszeges
	// w.Csaptipus = "Tömör csapos"   // Tömör | Csőcsapos
	// w.Szemforma = "Piskóta" // Piskóta | Párhuzamos Profilú | Hajlított
	// w.WeightClass = "kg."
	// w.Category = "Patentszemek"
	// w.ClassId = "Görgőslánc"
	// w.Quantity = fmt.Sprintf("%.1f", p.Stock)
	// w.Alapar = fmt.Sprintf("%.0f", p.WebPrice)
	// w.TaxClass = "27%"
	// w.QuantityUnit = p.Unit
	// w.Weight = fmt.Sprintf("%.1f", p.Weight)

	// Ha nincs belőle raktáron, nem elérhető.
	// qty, _ := strconv.ParseFloat(w.Quantity, 64)
	// if qty == 0 {
	// 	w.ShortDescription += models.JelenlegNemElerheto
	// } else {
	// 	if slices.Contains(models.CsakRendelesre, family) {
	// 		w.ShortDescription += models.CsakRendelesreLeiras
	// 	}
	// }
	// w.ShortDescription += models.Zaradek

	psp.ID = ""            // Az ID üres, cikkszámokkal dogozunk
	psp.Reference = p.Code // Ez a cikkszám, ID helyett használjuk
	psp.Active = "1"
	psp.PriceTaxExcluded = fmt.Sprintf("%.0f", p.WebPrice)
	psp.UnitPrice = fmt.Sprintf("%.0f", p.WebPrice) // TODO
	psp.TaxRulesID = "1"                            // ÁFA kulcs 27%
	psp.Quantity = fmt.Sprintf("%.1f", p.Stock)     // Mennyiség
	psp.Weight = fmt.Sprintf("%.1f", p.Weight)
	psp.Unity = "db"
	psp.TextInStock = "db"          // méter
	psp.TextBackorderAllowed = "db" // méter
	psp.Categories = "Láncok/Patentszemek"
	psp.TextInStock = "Raktáron"
	psp.TextBackorderAllowed = "Rendelhető"
	psp.ShowPrice = "1"            // TODO
	psp.OnSale = "0"               // Akció számítása
	psp.DiscountAmount = ""        // TODO
	psp.DiscountPercent = "0"      // TODO
	psp.DiscountFrom = ""          // TODO
	psp.DiscountTo = ""            // TODO
	psp.DeleteExistingImages = "0" // TODO

	// Normál patetszem
	// N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGLPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[productType+sorokszama]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Normál"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Standard Patentszem", productType, sorokszama)
		psp.Tags = "Standard,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,standard,patentszem"
	}

	// Heavy patentszem
	// N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])_H$
	match = regExpGLPSZ_H.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Igen"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Normál"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Erősített Patentszem", productType, sorokszama)
		psp.Tags = "Erősített,Heavy,patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,erősített,heavy,patentszem"
	}

	// Sasszeges, normál patentszem
	// N-(GLSPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGLSPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Acél Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Sasszeg,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeg,patentszem"
	}

	// Rozsdamentes görgőslánc patentszem rugós lemezes
	// N-(SSGLPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpSSGLPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Rozsdamentes acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,patentszem"
	}

	// Rozsdamentes görgőslánc patentszem sasszeges
	// N-(SSGLSPSZ)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpSSGLSPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Rozsdamentes acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Sasszeg,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeg,patentszem"
	}

	// Sasszeges hajlított patentszem (a hajlított mindig sasszeges)
	// N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpGLHOK.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		productType = match[3]
		sorokszama = match[4]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Hajlított"
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Hajlított Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Sasszeg,Hajlított,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeg,hajlított,patentszem"
	}

	// Erősített görgőslánc sasszeges hajlított patentszem
	// N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])_H$
	match = regExpGLHOK_H.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = match[3]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Igen"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Hajlított"
		features["Rögzítés"] = "Sasszeges"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Erősített Hajlított Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Erősített,Heavy,Sasszeg,Hajlított,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,erősített,heavy,sasszeges,hajlított,patentszem"
	}

	// Rozsdamentes görgősláncok rugós sasszeges hajlított patentszeme
	// N-(SSGLHOK)-([0-9]+)-([0-9ABC]+)([123])$
	match = regExpSSGLHOK.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = match[3]
		pStr = productType + sorokszama

		features["Anyag"] = "Rozsdamentes Acél"
		features["Erősített"] = "Igen"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Sasszeges"
		features["Szemforma"] = "Hajlított"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Rozsdamentes Hajlított Sasszeges Patentszem", productType, sorokszama)
		psp.Tags = "Rozsdamentes,Hajlított,Sasszeg,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,rozsdamentes,sasszeges,hajlított,sasszeg,patentszem"
	}

	// GLPSZ MOFA
	// N-(GLPSZ)-([0-9]+)-([0-9ABC]+)([123])_MOFA$
	match = regExpGLPSZMOFA.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = fmt.Sprintf("%s%s_MOFA", match[3], sorokszama)

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s MOFA Patentszem", productType, sorokszama)
		psp.Tags = "MOFA,Patentszem"
		psp.MetaKeywords = "Mofa,Patentszem"
	}

	// Hajlított MOFA
	// N-(GLHOK)-([0-9]+)-([0-9ABC]+)([123])_MOFA$
	match = regExpGLHOKMOFA.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = match[4]
		productType = fmt.Sprintf("%s%s_MOFA", match[3], sorokszama)

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Hajlított"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s MOFA Patentszem", productType, sorokszama)
		psp.Tags = "MOFA,Patentszem,Összekötő szem"
		psp.MetaKeywords = "Mofa,Patentszem,összekötő szem"
	}

	// VELO Patentszem
	// N-(GLPSZ)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$
	match = regExpGLPSZVELO.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Velo mindig egysoros
		productType = fmt.Sprintf("%s", match[3])

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Piskóta"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s VELO Patentszem", productType, sorokszama)
		psp.Tags = "VELO,Patentszem"
		psp.MetaKeywords = "Velo,Patentszem"

		// Egyedi képei vannak
		psp.ImageURLs = fmt.Sprintf(
			"%s/N-GLPSZ-%s.png,%s/D-GLPSZ-%s.png",
			models.ImagesBase, sorokszama,
			models.ImagesBase, sorokszama)
		psp.ImageAltTexts = psp.Name

		// w.Image = fmt.Sprintf("product/N-GLPSZ-%s.png", sorokszama)
		// w.ImageAdditional = fmt.Sprintf("product/D-GLPSZ-%s.png", sorokszama)
	}

	// Hajlított VELO
	// N-(GLHOK)-([0-9]+)-([0-9,\/]+x[0-9,\/]+)_VELO?$
	match = regExpGLHOKVELO.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Velo mindig egysoros
		productType = fmt.Sprintf("%s", match[3])

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[pStr]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Tömör csapos"
		features["Szemforma"] = "Hajlított"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s VELO Patentszem", productType, sorokszama)
		psp.Tags = "VELO,Patentszem,Összekötő szem"
		psp.MetaKeywords = "Velo,Patentszem,Összekötő szem"
	}

	// Csőcsapos egyenes patentszem
	// N-(CSCSGLPSZ)-([0-9]+)-([0-9ABC]+)1$
	match = regExpCSCSGLPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Mindig egysoros
		productType = match[3]
		pStr = productType + sorokszama

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[productType+sorokszama]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Csőcsapos"
		features["Szemforma"] = "Egyenes"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Csőcsapos Egyenes Patentszem", productType, sorokszama)
		psp.Tags = "Csőcsap,Egyenes,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,egyenes,csőcsap,patentszem"
	}

	// Párhuzamos profilú patentszem
	// regExpPPGLPSZ = regexp.MustCompile(`N-(PPGLPSZ)-([0-9])+-([0-9ABC]+)1$`)
	match = regExpPPGLPSZ.FindStringSubmatch(p.Code)
	if match != nil {
		family = match[1]
		manufacturerId = match[2]
		sorokszama = "1" // Mindig egysoros
		productType = match[3]

		features["Anyag"] = "Acél"
		features["Erősített"] = "Nem"
		features["Osztás"] = GLParms[pStr]["Osztas"]
		features["Belső hevedertávolság"] = GLParms[pStr]["Belsoheveder"]
		features["Görgőátmérő"] = GLParms[pStr]["Gorgoatmero"]
		features["Csapátmérő"] = GLParms[productType+sorokszama]["Csapatmero"]
		features["Csaphossz"] = GLParms[pStr]["Csaphossz"]
		features["Csaptípus"] = "Csőcsapos"
		features["Szemforma"] = "Párhuzamos profilú"
		features["Rögzítés"] = "Rugós lemezes"
		features["Kivitel"] = models.Sornevek[sorokszama]
		psp.Name = fmt.Sprintf("%s%s Standard Patentszem", productType, sorokszama)
		psp.Tags = "Standard,Patentszem"
		psp.MetaKeywords = models.Sornevek[features["Fogszám"]] + "soros,standard,patentszem"
	}

	// Gyártó beállítása
	mIdTmp, _ := strconv.Atoi(manufacturerId)
	psp.Manufacturer, _ = models.Manufacturers[mIdTmp]

	kemenysegTmp := "Standard"
	if features["Erősített"] == "Igen" {
		kemenysegTmp = "Erősített"
	}
	psp.Summary = fmt.Sprintf(
		"%s gyártmányú %s, %s mm osztású, %s mm belső hevedertávolságú, %s mm görgőátmérőjű %s szemformájú %s %s %s patentszem.",
		psp.Manufacturer, strings.ToLower(models.Sornevek[sorokszama]),
		features["Osztás"], features["Belső hevedertávolság"], features["Görgőátmérő"],
		strings.ToLower(features["Szemforma"]),
		strings.ToLower(features["Csaptípus"]), strings.ToLower(kemenysegTmp),
		strings.ToLower(features["Anyag"]))
	// Ebben a fázisában kell beállítani és nem lehet pont a végén.
	psp.MetaDescription = strings.TrimRight(psp.Summary, ".")
	psp.Summary += models.Zaradek

	psp.MetaTitle = psp.Name
	psp.URLRewritten = p.Code

	// Képek előállítása (a Velonál egyedileg készült)
	if psp.ImageURLs == "" {
		psp.ImageURLs = fmt.Sprintf(
			"%s/N-%s-%s.png,%s/D-%s-%s.png",
			models.ImagesBase, family, sorokszama,
			models.ImagesBase, family, sorokszama)
		psp.ImageAltTexts = psp.Name
	}

	// Speciális tulajdonságok beállítása
	psp.Features = models.MkFeaturesList(features)

	// Kapcsolódó termékek
	psp.Accessories = ""
	// TODO Ideiglenesen kivesszük, mert ha nem ltezik a termék, nagyon lelassul
	//rgxStr := fmt.Sprintf(`^N-GL-[0-9]+-%s%s.*`, productType, sorokszama)
	//psp.Accessories += models.getRelatedProductIds(rgxStr, prodCodes)

	// Rendelhető?
	psp.AvailableForOrder = "1"
	if slices.Contains(models.CsakRendelesre, family) {
		psp.AvailableForOrder = "0"
	}

	return psp
}
