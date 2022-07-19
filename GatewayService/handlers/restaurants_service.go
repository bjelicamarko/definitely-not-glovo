package handlers

import (
	"GatewayService/utils"
	"net/http"
)

func FindAllRestaurants(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	response, err := http.Get(
		utils.RestaurantsServiceRoot.Next().Host + RestaurantsServiceApi + "/getRestaurants")

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}
