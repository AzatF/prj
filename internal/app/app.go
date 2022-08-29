package app

import (
	"net/http"
	"project/pkg/logging"
)

func StartServer(host, port string, logger *logging.Logger) {

	logger.Infof("server listening oh %s:%s", host, port)
	logger.Fatal(http.ListenAndServe(host+":"+port, nil))

}
