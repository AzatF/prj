package sms

import (
	"io/ioutil"
	"log"
	"path"
	"project/app/alpha2"
	"project/app/model"
	"project/config"
	"project/pkg/logging"
	"sort"
	"strings"
)

var (
	smsSliceString model.SMSDataModel
	smsSliceSum    []model.SMSDataModel
	first          []model.SMSDataModel
)

func CheckSMSInfo(cfg *config.Config, logger *logging.Logger) ([]model.SMSDataModel, error) {

	file, err := ioutil.ReadFile(path.Join(cfg.DataPath, "sms.data"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	codeA2, err := alpha2.CountryCodeAlpha2()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	prov := strings.Split(cfg.Providers, " ")
	smsFile := strings.Split(string(file), "\n")

	for _, v := range smsFile {

		smsDataString := strings.Split(v, ";")

		if len(smsDataString) == 4 {

			for _, k := range prov {

				for _, c := range codeA2 {

					if smsDataString[3] == k && smsDataString[0] == c.Alpha2 {

						smsSliceString.Country = smsDataString[0]
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

func SortSMSInfo(smsInfo []model.SMSDataModel, logger *logging.Logger) (sorted []model.SMSDataModel, err error) {

	codeA2, err := alpha2.CountryCodeAlpha2()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, v := range smsInfo {

		for _, c := range codeA2 {

			if v.Country == c.Alpha2 {

				v.Country = c.Country
				first = append(first, v)
			}
		}
	}

	sort.SliceStable(first, func(i, j int) bool {
		return first[i].Country < first[j].Country
	})

	return first, nil

}
