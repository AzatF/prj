package incident

import (
	"encoding/json"
	"net/http"
	"project/config"
	"project/internal/model"
	"project/pkg/logging"
	"sort"
)

var (
	incidentInfo []model.IncidentDataModel
)

func CheckIncidentInfo(cfg *config.Config, logger *logging.Logger) ([]model.IncidentDataModel, error) {

	resp, err := http.Get("http://" + cfg.IncidentHost + ":" + cfg.IncidentPort + "/accendent")
	if err != nil {
		logger.Errorf("Status code incident: %v", resp.StatusCode)
		logger.Fatal(err)
	}
	logger.Infof("Status code incident: %v", resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		err = json.NewDecoder(resp.Body).Decode(&incidentInfo)
		if err != nil {
			return nil, err
		}

		sort.SliceStable(incidentInfo, func(i, j int) bool {
			return incidentInfo[i].Status < incidentInfo[j].Status
		})

	}

	return incidentInfo, nil
}
