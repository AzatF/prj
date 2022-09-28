package incident

import (
	"encoding/json"
	"net/http"
	"project/config"
	"project/internal/model"
	"project/pkg/logging"
	"sort"
)

func CheckIncidentInfo(cfg *config.Config, logger *logging.Logger) ([]model.IncidentDataModel, error) {

	var incidentInfo []model.IncidentDataModel

	resp, err := http.Get(cfg.IncidentHost + ":" + cfg.IncidentPort)
	if err != nil {
		return nil, err
	}

	logger.Infof("Status code incident: %v", resp.StatusCode)

	if resp.StatusCode == 200 {

		err = json.NewDecoder(resp.Body).Decode(&incidentInfo)
		if err != nil {
			return nil, err
		}

		sort.SliceStable(incidentInfo, func(i, j int) bool {
			return incidentInfo[i].Status < incidentInfo[j].Status
		})

	}
	defer resp.Body.Close()

	return incidentInfo, nil
}
