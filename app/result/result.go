package result

import (
	"project/app/billing"
	"project/app/email"
	"project/app/incident"
	"project/app/mms"
	"project/app/model"
	"project/app/sms"
	"project/app/support"
	"project/app/voice"
	"project/config"
	"project/pkg/logging"
)

func GetResultData(cfg *config.Config, logger *logging.Logger) (model.ResultSetT, error) {

	smsInfo, err := sms.CheckSMSInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}

	sortedSMS, err := sms.SortSMSInfo(smsInfo, logger)
	if err != nil {
		logger.Error(err)
	}

	mmsInfo, err := mms.CheckMMSInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}

	sortedMMS, err := mms.SortMMSInfo(mmsInfo, logger)
	if err != nil {
		logger.Error(err)
	}

	voiceInfo, err := voice.CheckVoiceInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}

	emailInfo, err := email.CheckEmailInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}

	fast, slow, err := email.SortEmailInfo(emailInfo, logger)
	if err != nil {
		logger.Error(err)
	}

	billingInfo, err := billing.CheckBillingInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}

	supportInfo, err := support.CheckSupportInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}

	sortedSupportInfo, err := support.SortSupportInfo(supportInfo)
	if err != nil {
		logger.Error(err)
	}

	incidentInfo, err := incident.CheckIncidentInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}

	var resultT model.ResultSetT

	resultT.SMS = append(resultT.SMS, smsInfo)
	resultT.SMS = append(resultT.SMS, sortedSMS)
	resultT.MMS = append(resultT.MMS, mmsInfo)
	resultT.MMS = append(resultT.MMS, sortedMMS)
	resultT.VoiceCall = append(resultT.VoiceCall, voiceInfo...)

	resultT.Email = make(map[string][][]model.EmailDataModel)
	for i, v := range fast {
		resultT.Email[i] = append(resultT.Email[i], v)
	}
	for i, v := range slow {
		resultT.Email[i] = append(resultT.Email[i], v)
	}

	resultT.Billing = billingInfo
	resultT.Support = append(resultT.Support, sortedSupportInfo...)
	resultT.Incident = append(resultT.Incident, incidentInfo...)

	return resultT, nil

}
