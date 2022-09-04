package email

import (
	"os"
	"path"
	"project/config"
	"project/internal/alpha2"
	"project/internal/model"
	"project/pkg/logging"
	"sort"
	"strconv"
	"strings"
)

func CheckEmailInfo(cfg *config.Config, logger *logging.Logger) ([]model.EmailDataModel, error) {

	var (
		emailInfo    model.EmailDataModel
		emailInfoSum []model.EmailDataModel
	)

	file, err := os.ReadFile(path.Join(cfg.DataPath, "email.data"))
	if err != nil {
		return nil, err
	}

	codeA2, err := alpha2.CountryCodeAlpha2()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	emailProviders := strings.Split(cfg.ProvidersEmail, " ")
	emailFile := strings.Split(string(file), "\n")

	for _, v := range emailFile {
		emailFileInfo := strings.Split(v, ";")

		if len(emailFileInfo) == 3 {

			for _, k := range emailProviders {

				for _, c := range codeA2 {

					if emailFileInfo[0] == c.Alpha2 && emailFileInfo[1] == k {

						emailInfo.Country = emailFileInfo[0]
						emailInfo.Provider = emailFileInfo[1]
						emailInfo.DeliveryTime, err = strconv.Atoi(emailFileInfo[2])
						if err != nil {
							logger.Error(err)
						}
						emailInfoSum = append(emailInfoSum, emailInfo)
					}
				}
			}
		}
	}

	sort.SliceStable(emailInfoSum, func(i, j int) bool {
		return emailInfoSum[i].DeliveryTime < emailInfoSum[j].DeliveryTime
	})

	return emailInfoSum, nil

}

func SortEmailInfo(emailInfo []model.EmailDataModel, logger *logging.Logger) (map[string][]model.EmailDataModel, map[string][]model.EmailDataModel, error) {

	var mapEmailInfo = map[string][]model.EmailDataModel{}
	var mapEmailInfoFast = map[string][]model.EmailDataModel{}
	var mapEmailInfoSlow = map[string][]model.EmailDataModel{}

	codeA2, err := alpha2.CountryCodeAlpha2()
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	for _, v := range emailInfo {
		for _, k := range codeA2 {
			if v.Country == k.Alpha2 {
				mapEmailInfo[v.Country] = append(mapEmailInfo[v.Country], v)
			}
		}
	}

	for _, v := range mapEmailInfo {
		sort.SliceStable(v, func(i, j int) bool {
			return v[i].DeliveryTime < v[j].DeliveryTime
		})
		for i, k := range v {

			if i <= 2 {
				mapEmailInfoFast[k.Country] = append(mapEmailInfoFast[k.Country], k)
			}

			if i >= len(v)-3 {
				mapEmailInfoSlow[k.Country] = append(mapEmailInfoSlow[k.Country], k)
			}
		}
	}

	return mapEmailInfoFast, mapEmailInfoSlow, nil

}
