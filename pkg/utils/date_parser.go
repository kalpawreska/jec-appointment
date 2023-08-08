package utils

import (
	"time"

	"github.com/go-openapi/strfmt"
)

func DateParser(freeFormattedDatetime *string, format string, mockedTime ...bool) *string {
	if freeFormattedDatetime == nil {
		return nil
	}

	var parsedDate string
	dob, err := strfmt.ParseDateTime(*freeFormattedDatetime)
	if err == nil {
		parsedDate = time.Time(dob).Format(format)
		return &parsedDate
	}

	// retry but with mocked time
	if len(mockedTime) == 0 {
		tryDate := *freeFormattedDatetime + " 00:00:00"
		return DateParser(&tryDate, format, true)
	}

	return nil
}

func DateTimeParserRaw(freeFormattedDatetime *string) *strfmt.DateTime {
	if freeFormattedDatetime == nil {
		return nil
	}

	if dob, err := strfmt.ParseDateTime(*freeFormattedDatetime); err == nil {
		return &dob
	}
	return nil
}

func DateParserRaw(freeFormattedDatetime *string) *strfmt.Date {
	dateTimeParsed := DateTimeParserRaw(freeFormattedDatetime)
	if dateTimeParsed != nil {
		tgl := strfmt.Date(*dateTimeParsed)
		return &tgl
	}
	return nil
}
