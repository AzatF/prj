package test

import (
	"project/app/billing"
	"project/config"
	"project/pkg/logging"
)

var d int

func TestingApp(cfg *config.Config, logger *logging.Logger) {

	//------------SMS---------------------------------------------

	//smsDataInfo, _ := sms.CheckSMSInfo(cfg, logger)
	//for _, v := range smsDataInfo {
	//	d++
	//	logger.Infof("from sms: %d: %s\n", d, v)
	//}
	//d = 0
	//sortedSMS, _ := sms.SortSMSInfo(smsDataInfo, logger)
	//for _, v := range sortedSMS {
	//	d++
	//	logger.Infof("from sms: %d: %s\n", d, v)
	//}
	//d = 0

	//-------------MMS---------------------------------------------

	//mmsInfo, err := mms.CheckMMSInfo(cfg, logger)
	//if err != nil {
	//	logger.Error(err)
	//}
	//for _, v := range mmsInfo {
	//	d++
	//	logger.Infof("from mms: %d: %s\n", d, v)
	//}
	//d = 0

	//------------VOICE-------------------------------------------------

	//voiceInfo, err := voice.CheckVoiceInfo(cfg, logger)
	//if err != nil {
	//	logger.Error(err)
	//}
	//for _, v := range voiceInfo {
	//	d++
	//	logger.Infof("from voice: %d: %v\n", d, v)
	//}
	//d = 0

	//-------------EMAIL------------------------------------------------

	//emailInfo, err := email.CheckEmailInfo(cfg, logger)
	//if err != nil {
	//	logger.Error(err)
	//}
	//
	//fast, slow, err := email.SortEmailInfo(emailInfo, logger)
	//if err != nil {
	//	logger.Error(err)
	//}
	//
	//for i, s := range fast {
	//	logger.Infof("email info from func fast:%s - %v", i, s)
	//}
	//
	//for i, s := range slow {
	//	logger.Infof("email info from func slow:%s - %v", i, s)
	//}

	//for i, v := range emailInfo {
	//	d++
	//	logger.Infof("from email: %d: %v", i, v)
	//}
	//d = 0

	//============BILLING--------------------------------------------------

	billingInfo, err := billing.CheckBillingInfo(cfg, logger)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("billing:\nCheckoutPage: %v,\nFraudControl %v,\nRecurring %v,\nPayout %v,\nPayout %v,\nCreateCustomer %v", billingInfo.CheckoutPage, billingInfo.FraudControl,
		billingInfo.Recurring, billingInfo.Payout, billingInfo.Purchase, billingInfo.CreateCustomer)

	//------------SUPPORT----------------------------------------------------

	//supportInfo, err := support.CheckSupportInfo(cfg, logger)
	//if err != nil {
	//	logger.Error(err)
	//}
	//
	//for _, v := range supportInfo {
	//	logger.Infof("support info: %v\n", v)
	//}
	//
	//sortedSupportInfo, err := support.SortSupportInfo(supportInfo)
	//if err != nil {
	//	logger.Error(err)
	//}
	//logger.Infof("sorted support %d : %d", sortedSupportInfo[0], sortedSupportInfo[1])

	//------------INCIDENT-------------------------------------------------

	//incidentInfo, err := incident.CheckIncidentInfo(cfg, logger)
	//if err != nil {
	//	logger.Error(err)
	//}
	//
	//for _, v := range incidentInfo {
	//	logger.Infof("incident info: %v\n", v)
	//}

	//----------RESULT T------------------------------------------------

	//a, err := result.GetResultData(cfg, logger)
	//if err != nil {
	//	logger.Error(err)
	//}
	//for _, v := range a.SMS {
	//	for _, k := range v {
	//		logger.Infof("result sms info: %v\n", k)
	//	}
	//}
	//for _, v := range a.MMS {
	//	for _, k := range v {
	//		logger.Infof("result mms info: %v\n", k)
	//	}
	//}
	//for _, v := range a.VoiceCall {
	//	logger.Infof("result voiceCall info: %v\n", v)
	//}
	//
	//for _, v := range a.Email {
	//	logger.Infof("result email info: %v\n", v)
	//}
	//
	//logger.Infof("result billing info: %v\n", a.Billing)
	//logger.Infof("result support info: %v\n", a.Support)
	//
	//for _, v := range a.Incident {
	//	logger.Infof("result incident info: %v\n", v)
	//}

}
