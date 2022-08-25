package voice

import (
	"io/ioutil"
	"path"
	"project/app/alpha2"
	"project/app/model"
	"project/config"
	"project/pkg/logging"
	"strconv"
	"strings"
)

var (
	voiceInfo    model.VoiceDataModel
	voiceInfoSum []model.VoiceDataModel
	sum          float64
)

func CheckVoiceInfo(cfg *config.Config, logger *logging.Logger) ([]model.VoiceDataModel, error) {

	file, err := ioutil.ReadFile(path.Join(cfg.DataPath, "voice.data"))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	codeA2, err := alpha2.CountryCodeAlpha2()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	voiceProviders := strings.Split(cfg.ProvidersVoice, " ")
	voiceFile := strings.Split(string(file), "\n")

	for _, v := range voiceFile {
		voiceFileInfo := strings.Split(v, ";")

		if len(voiceFileInfo) == 8 {

			for _, k := range voiceProviders {

				for _, c := range codeA2 {

					if voiceFileInfo[3] == k && voiceFileInfo[0] == c.Alpha2 {

						voiceInfo.Country = voiceFileInfo[0]
						voiceInfo.Bandwidth = voiceFileInfo[1]
						voiceInfo.ResponseTime = voiceFileInfo[2]
						voiceInfo.Provider = voiceFileInfo[3]
						sum, err = strconv.ParseFloat(voiceFileInfo[4], 32)
						if err != nil {
							logger.Error(err)
						}

						voiceInfo.ConnectionStability = float32(sum)
						voiceInfo.TTFB, err = strconv.Atoi(voiceFileInfo[5])
						if err != nil {
							logger.Error(err)
						}

						voiceInfo.VoicePurity, err = strconv.Atoi(voiceFileInfo[6])
						if err != nil {
							logger.Error(err)
						}

						voiceInfo.MedianOfCallsTime, err = strconv.Atoi(voiceFileInfo[7])
						if err != nil {
							logger.Error(err)
						}

						voiceInfoSum = append(voiceInfoSum, voiceInfo)
					}
				}
			}
		}

		if len(voiceFileInfo) != 8 {
			logger.Warnf("broken line VoiceCall: %s", v)
		}
	}

	return voiceInfoSum, nil

}
