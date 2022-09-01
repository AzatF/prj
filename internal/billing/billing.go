package billing

import (
	"io/ioutil"
	"path"
	"project/config"
	"project/internal/model"
	"project/pkg/logging"
	"strconv"
)

func CheckBillingInfo(cfg *config.Config, logger *logging.Logger) (billingInfo model.BillingDataModel, err error) {

	var (
		n1, n2, n3, n4, n5, n6 int
	)

	file, err := ioutil.ReadFile(path.Join(cfg.DataPath, "billing.data"))
	if err != nil {
		logger.Error(err)
		return billingInfo, err
	}

	a := string(file)

	if len(a) == 6 {

		n1, err = strconv.Atoi(string([]rune(a)[0]))
		n2, err = strconv.Atoi(string([]rune(a)[1]))
		n3, err = strconv.Atoi(string([]rune(a)[2]))
		n4, err = strconv.Atoi(string([]rune(a)[3]))
		n5, err = strconv.Atoi(string([]rune(a)[4]))
		n6, err = strconv.Atoi(string([]rune(a)[5]))
		if err != nil {
			logger.Error(err)
		}
	}

	if n1 == 1 {
		billingInfo.CheckoutPage = true
	}
	if n2 == 1 {
		billingInfo.FraudControl = true
	}
	if n3 == 1 {
		billingInfo.Recurring = true
	}
	if n4 == 1 {
		billingInfo.Payout = true
	}
	if n5 == 1 {
		billingInfo.Purchase = true
	}
	if n6 == 1 {
		billingInfo.CreateCustomer = true
	}

	return billingInfo, nil
}
