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
	timeCachedResult() (model.ResultSetT, error)
	GetResultData() (model.ResultSetT, error)
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
	CacheResultOld model.ResultSetT
	stTimeEnd      = time.Now().Format(time.RFC3339)
)

func (r *CashResultData) GetResultData() (model.ResultSetT, error) {

	r.Mu.Lock()
	defer r.Mu.Unlock()

	if time.Now().Format(time.RFC3339) < stTimeEnd {
		r.logger.Info("result from cache!===")
		return CacheResultOld, nil
	} else {
		CacheResultNew, err := r.timeCachedResult()
		r.logger.Info("new result!===")
		stTimeEnd = time.Now().Add(30 * time.Second).Format(time.RFC3339)
		if err != nil {
			return CacheResultOld, err
		}
		return CacheResultNew, nil
	}

}

func (r *CashResultData) timeCachedResult() (model.ResultSetT, error) {

	CacheResult := model.ResultSetT{}

	smsInfo, err := sms.CheckSMSInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	sortedSMS, err := sms.SortSMSInfo(smsInfo, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	mmsInfo, err := mms.CheckMMSInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	sortedMMS, err := mms.SortMMSInfo(mmsInfo, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	voiceInfo, err := voice.CheckVoiceInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	emailInfo, err := email.CheckEmailInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	fast, slow, err := email.SortEmailInfo(emailInfo, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	billingInfo, err := billing.CheckBillingInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	supportInfo, err := support.CheckSupportInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Error(err)
	}

	sortedSupportInfo, err := support.SortSupportInfo(supportInfo)
	if err != nil {
		r.logger.Error(err)
	}

	incidentInfo, err := incident.CheckIncidentInfo(r.cfg, r.logger)
	if err != nil {
		r.logger.Error(err)
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

	return CacheResult, nil

}
