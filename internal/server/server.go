package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/config"
	"project/internal/model"
	"project/internal/result"
	"project/pkg/logging"
	"text/template"
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

var resultT model.ResultT

func (l *list) HomeServer() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/api", l.handleConnection)

	http.Handle("/", rtr)
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./web/"))))

}

func (l *list) handleConnection(w http.ResponseWriter, r *http.Request) {

	a, err := result.NewResult(l.logger, l.cfg)
	if err != nil {
		l.logger.Error(err)
	}
	resultNew, err := a.GetResultData()
	resultT.Status = true
	resultT.Error = ""
	if err != nil {
		l.logger.Error(err)
		resultT.Status = false
		resultT.Error = "Error on collect data"
	}

	resultT.Data = resultNew
	//resp, _ := json.Marshal(resultT)
	resp, _ := json.MarshalIndent(resultT, "", " ")

	//_, _ = w.Write(resp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	t, err := template.ParseFiles("./web/status_page.html")
	if err != nil {
		w.WriteHeader(404)
		l.logger.Error(err)
		return
	}
	err = t.ExecuteTemplate(w, "status_page.html", resp)
	if err != nil {
		l.logger.Error(err)
	}
	//_, _ = w.Write(resp)

}
