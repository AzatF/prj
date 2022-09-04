package alpha2

import (
	"os"
	"project/internal/model"
	"strings"
)

var (
	ss      []string
	isoCode model.ISO3166
)

func CountryCodeAlpha2() (isoCodeRes []model.ISO3166, err error) {

	codeAlpha2, err := os.ReadFile("./data/country_code.csv")
	if err != nil {
		return nil, err
	}

	codeA2 := strings.Split(string(codeAlpha2), ";")

	if len(codeA2) > 0 {
		for _, v := range codeA2 {
			ss = strings.Split(v, ",")
			if len(ss) == 4 {
				isoCode.Country = strings.TrimSpace(ss[0])
				isoCode.Alpha2 = strings.TrimSpace(ss[1])
				isoCode.Alpha3 = strings.TrimSpace(ss[2])
				isoCode.Code = strings.TrimSpace(ss[3])
				isoCodeRes = append(isoCodeRes, isoCode)
			}
		}
	}

	return isoCodeRes, nil
}
