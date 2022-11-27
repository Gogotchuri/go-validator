package indiaValidators

import (
	"regexp"
)

const (
	indiaGSTINString   = "^[0-9]{2}[A-Z]{5}[0-9]{4}[A-Z]{1}[1-9A-Z]{1}[Z]{1}[0-9A-Z]{1}$"
	indiaTransINString = "^[0-9]{2}[0-9A-Z]{13}$"
	indiaPANString     = "^[a-zA-Z]{5}[0-9]{4}[a-zA-Z]{1}$"
)

var (
	IndiaGSTINRegex   = regexp.MustCompile(indiaGSTINString)
	IndiaTransINRegex = regexp.MustCompile(indiaTransINString)
	IndiaPANRegex     = regexp.MustCompile(indiaPANString)
)
