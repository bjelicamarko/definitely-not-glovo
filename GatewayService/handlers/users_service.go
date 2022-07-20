package handlers

import (
	"GatewayService/utils"
	"net/http"
)

func Login(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	req, _ := http.NewRequest(http.MethodPost,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/login", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func Register(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	req, _ := http.NewRequest(http.MethodPost,
		utils.UsersServiceRoot.Next().Host+UsersServiceApi+"/register", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}
