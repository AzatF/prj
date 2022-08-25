package support

import (
	"encoding/json"
	"math"
	"net/http"
	"project/app/model"
	"project/config"
	"project/pkg/logging"
)

var (
	sumTicket int
	workLoad  int
)

func CheckSupportInfo(cfg *config.Config, logger *logging.Logger) (supportInfo []model.SupportDataModel, err error) {

	resp, err := http.Get("http://" + cfg.SupportHost + ":" + cfg.SupportPort + "/support")
	if err != nil {
		logger.Errorf("Status code support: %v", resp.StatusCode)
		return nil, err
	}

	logger.Infof("Status code support: %v", resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		err = json.NewDecoder(resp.Body).Decode(&supportInfo)
		if err != nil {
			return nil, err
		}

		return supportInfo, nil
	}

	return
}

func SortSupportInfo(supportInfo []model.SupportDataModel) (sortedSupportInfo []int, err error) {

	medianTime := 60.0 / 18.0 / 1.0
	for _, v := range supportInfo {
		sumTicket += v.ActiveTickets
	}

	waitTime := float64(sumTicket) * medianTime

	if int(sumTicket) >= 9 && int(sumTicket) < 16 {
		workLoad = 2
	}
	if int(sumTicket) < 9 {
		workLoad = 1
	}
	if int(sumTicket) >= 16 {
		workLoad = 3
	}

	sortedSupportInfo = append(sortedSupportInfo, workLoad)
	sortedSupportInfo = append(sortedSupportInfo, int(math.Round(waitTime)))

	return
}
