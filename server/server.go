package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project/app/model"
	result2 "project/app/result"
	"project/config"
	"project/pkg/logging"
	"text/template"
)

type list struct {
	cfg    *config.Config
	logger *logging.Logger
}

type Handlers interface {
	HomeServer(host, port string, logger *logging.Logger)
	handleConnection(w http.ResponseWriter, r *http.Request)
}

func NewServer(cfg *config.Config, logger *logging.Logger) (Handlers, error) {
	return &list{
		cfg:    cfg,
		logger: logger,
	}, nil
}

var resultT model.ResultT

func (l *list) HomeServer(host, port string, logger *logging.Logger) {

	logger.Infof("server listening oh %s:%s", host, port)
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api", l.handleConnection)

	http.Handle("/", rtr)
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./web/"))))
	logger.Fatal(http.ListenAndServe(host+":"+port, nil))

}

func (l *list) handleConnection(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	result, err := result2.GetResultData(l.cfg, l.logger)
	resultT.Status = true
	resultT.Error = ""
	if err != nil {
		l.logger.Error(err)
		resultT.Status = false
		resultT.Error = "Error on collect model"
	}

	resultT.Data = result

	resp, _ := json.Marshal(resultT)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//_, _ = w.Write(resp)

	t, err := template.ParseFiles("./web/status_page.html", "./web/main.js" /*, "./web/chart.min.js", "./web/main.css"*/)
	if err != nil {
		w.WriteHeader(404)
		l.logger.Error(err)
		return
	}
	err = t.ExecuteTemplate(w, "status_page.html", resp)
	if err != nil {
		l.logger.Error(err)
	}

}
