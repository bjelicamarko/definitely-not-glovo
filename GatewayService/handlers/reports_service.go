package handlers

import (
	"GatewayService/utils"
	"net/http"
)

func GetReports(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	response, err := http.Get(
		utils.ReportsServiceRoot.Next().Host + ReportsServiceApi +
			"/getReports")

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}
