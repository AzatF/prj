package result

import (
	"project/config"
	"project/internal/billing"
	"project/internal/email"
	"project/internal/incident"
	"project/internal/mms"
	"project/internal/model"
	"project/internal/sms"
	"project/internal/support"
	"project/internal/voice"
	"project/pkg/logging"
	"sync"
	"time"
)

type CashResultData struct {
	cfg    *config.Config
	Mu     *sync.Mutex
	logger *logging.Logger
}

type MakerRes interface {
	timeCachedResult() model.ResultSetT
	GetResultData() model.ResultSetT
}

func NewResult(logger *logging.Logger, cfg *config.Config) (MakerRes, error) {

	mu := sync.Mutex{}

	return &CashResultData{
		cfg:    cfg,
		Mu:     &mu,
		logger: logger,
	}, nil
}

var (
	CacheResultOld   model.ResultSetT
	stTimeEnd        = time.Now().Format(time.RFC3339)
	CollectDataError = false
)

func (r *CashResultData) GetResultData() model.ResultSetT {

	r.Mu.Lock()
	defer r.Mu.Unlock()

	if time.Now().Format(time.RFC3339) < stTimeEnd && !CollectDataError {
		r.logger.Info("result from cache!===")
		return CacheResultOld
	} else {
		CacheResultNew := r.timeCachedResult()
		r.logger.Info("new result!===")
		stTimeEnd = time.Now().Add(30 * time.Second).Format(time.RFC3339)
		return CacheResultNew
	}

}

func (r *CashResultData) timeCachedResult() model.ResultSetT {

	CacheResult := model.ResultSetT{}

	smsInfo, err := sms.CheckSMSInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Warn("Error collect SMS data")
		CollectDataError = true
	}

	sortedSMS, err := sms.SortSMSInfo(smsInfo)
	if err != nil {
		r.logger.Warn("Error sort SMS data")
		CollectDataError = true
	}

	mmsInfo, err := mms.CheckMMSInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Warn("Error collect MMS data")
		CollectDataError = true
	}

	sortedMMS, err := mms.SortMMSInfo(mmsInfo)
	if err != nil {
		r.logger.Warn("Error sort MMS data")
		CollectDataError = true
	}

	voiceInfo, err := voice.CheckVoiceInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Warn("Error collect VoiceCall data")
		CollectDataError = true
	}

	emailInfo, err := email.CheckEmailInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Warn("Error collect Email data")
		CollectDataError = true
	}

	fast, slow, err := email.SortEmailInfo(emailInfo, r.logger, r.cfg)
	if err != nil {
		r.logger.Warn("Error sort Email data")
		CollectDataError = true
	}

	billingInfo, err := billing.CheckBillingInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Warn("Error collect Billing data")
		CollectDataError = true
	}

	supportInfo, err := support.CheckSupportInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Warn("Error collect Support data")
		CollectDataError = true
	}

	sortedSupportInfo, err := support.SortSupportInfo(supportInfo)
	if err != nil {
		r.logger.Warn("Error sort Support data")
		CollectDataError = true
	}

	incidentInfo, err := incident.CheckIncidentInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Warn("Error collect Incident data")
		CollectDataError = true
	}

	CacheResult.SMS = append(CacheResult.SMS, smsInfo)
	CacheResult.SMS = append(CacheResult.SMS, sortedSMS)
	CacheResult.MMS = append(CacheResult.MMS, mmsInfo)
	CacheResult.MMS = append(CacheResult.MMS, sortedMMS)
	CacheResult.VoiceCall = append(CacheResult.VoiceCall, voiceInfo...)
	CacheResult.Email = make(map[string][][]model.EmailDataModel)
	for i, v := range fast {
		CacheResult.Email[i] = append(CacheResult.Email[i], v)
	}
	for i, v := range slow {
		CacheResult.Email[i] = append(CacheResult.Email[i], v)
	}
	CacheResult.Billing = billingInfo
	CacheResult.Support = append(CacheResult.Support, sortedSupportInfo...)
	CacheResult.Incident = append(CacheResult.Incident, incidentInfo...)

	CacheResultOld = CacheResult

	if err == nil {
		CollectDataError = false
	}

	return CacheResult

}
