import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { RestaurantDTO } from "src/modules/shared/models/RestaurantDTO";
import { RestaurantDTOMessage } from "src/modules/shared/models/RestaurantDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class RestaurantsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    saveRestaurant(restaurantDTO: RestaurantDTO): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<RestaurantDTOMessage>>("not-glovo/api/restaurants/saveRestaurant", restaurantDTO, queryParams);
    }
}