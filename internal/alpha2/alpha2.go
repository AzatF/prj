package alpha2

import (
	"os"
	"path"
	"project/config"
	"project/internal/model"
	"strings"
)

func CountryCodeAlpha2(cfg *config.Config) (isoCodeRes []model.ISO3166, err error) {

	var ss []string
	var isoCode model.ISO3166

	codeAlpha2, err := os.ReadFile(path.Join(cfg.BasePath, "country_code.csv"))
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

func GetProviders(cfg *config.Config, name string) ([]string, error) {

	var s []string
	var result []string

	file, err := os.ReadFile(path.Join(cfg.BasePath, "providers.csv"))
	if err != nil {
		return nil, err
	}

	res := strings.Split(string(file), ";")

	if len(res) > 1 {
		for _, v := range res {
			s = strings.Split(v, ",")
			if strings.TrimSpace(s[0]) == name {
				for i, k := range s {
					if i >= 1 {
						result = append(result, strings.TrimSpace(k))
					}
				}
			}
		}
	}

	return result, nil
}
