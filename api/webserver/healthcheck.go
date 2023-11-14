package webserver

import (
	"fmt"
	"net/http"
)

func HealthCheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "OK")
}
