package helper

import (
	"fmt"
	"time"
)

var (
	indonesiaDayNames = []string{
		"Minggu",
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
	}

	indonesiaMonthNames = []string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
)

func formatTimeToIndonesia(t time.Time) string {
	return fmt.Sprintf("%s, %02d %s %d",
		indonesiaDayNames[t.Weekday()], t.Day(), indonesiaMonthNames[t.Month()-1], t.Year(),
	)
}
