package handlers

import (
	"GatewayService/utils"
	"net/http"
)

func FindAllRestaurants(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.RestaurantsServiceRoot.Next().Host + RestaurantsServiceApi + "/getRestaurants?page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}
