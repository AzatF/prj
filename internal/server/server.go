package server

import (
	"encoding/json"
	"log"
	"net/http"
	"project/config"
	"project/internal/model"
	"project/internal/result"
	"project/pkg/logging"
)

type list struct {
	cfg    *config.Config
	logger *logging.Logger
	Cache  result.CashResultData
}

type Handlers interface {
	HomeServer()
	handleConnection(w http.ResponseWriter, r *http.Request)
}

func NewServer(cfg *config.Config, logger *logging.Logger) (Handlers, error) {
	return &list{
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (l *list) HomeServer() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api", l.handleConnection)
}

func (l *list) handleConnection(w http.ResponseWriter, r *http.Request) {

	resultT := model.ResultT{}
	a, err := result.NewResult(l.logger, l.cfg)
	if err != nil {
		l.logger.Error(err)
	}

	resultT.Data = a.GetResultData()

	if result.CollectDataError {
		resultT.Status = false
		resultT.Error = "error on collect data"
	} else {
		resultT.Status = true
		resultT.Error = ""
	}

	resp, err := json.MarshalIndent(resultT, "", " ")
	if err != nil {
		log.Printf("json marshal error: %v", err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, _ = w.Write(resp)

}
