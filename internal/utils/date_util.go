package utils

import "time"

func FormatDateToISO(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

func FormatDateToDDMMYYYY(t time.Time) string {
	return t.Format("02/01/2006")
}
