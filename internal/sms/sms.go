package sms

import (
	"os"
	"path"
	"project/config"
	"project/internal/alpha2"
	"project/internal/model"
	"project/pkg/logging"
	"sort"
	"strings"
)

func CheckSMSInfo(cfg *config.Config, logger *logging.Logger) ([]model.SMSDataModel, error) {

	var smsSliceString model.SMSDataModel
	var smsSliceSum []model.SMSDataModel

	file, err := os.ReadFile(path.Join(cfg.DataPath, "sms.data"))
	if err != nil {
		return nil, err
	}

	codeA2, err := alpha2.CountryCodeAlpha2(cfg)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	prov, err := alpha2.GetProviders(cfg, "sms")
	if err != nil {
		logger.Errorf("error read file providers: %v", err)
	}

	smsFile := strings.Split(string(file), "\n")

	for _, v := range smsFile {

		smsDataString := strings.Split(v, ";")

		if len(smsDataString) == 4 {

			for _, k := range prov {

				for _, c := range codeA2 {

					if smsDataString[3] == k && smsDataString[0] == c.Alpha2 {

						smsSliceString.Country = c.Country
						smsSliceString.Bandwidth = smsDataString[1]
						smsSliceString.ResponseTime = smsDataString[2]
						smsSliceString.Provider = smsDataString[3]
						smsSliceSum = append(smsSliceSum, smsSliceString)
					}
				}
			}
		}

		if len(smsDataString) != 4 {
			logger.Warnf("broken line SMS: %s", v)
		}
	}

	sort.SliceStable(smsSliceSum, func(i, j int) bool {
		return smsSliceSum[i].Provider < smsSliceSum[j].Provider
	})

	return smsSliceSum, nil
}

func SortSMSInfo(smsInfo []model.SMSDataModel) ([]model.SMSDataModel, error) {

	var first []model.SMSDataModel
	for _, v := range smsInfo {
		first = append(first, v)
	}

	sort.SliceStable(first, func(i, j int) bool {
		return first[i].Country < first[j].Country
	})

	return first, nil

}
