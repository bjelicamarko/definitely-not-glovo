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

    createRestaurant(restaurantDTO: RestaurantDTO): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<RestaurantDTOMessage>>("not-glovo/api/restaurants/createRestaurant", restaurantDTO, queryParams);
    }

    updateRestaurant(restaurantDTO: RestaurantDTO): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.put<HttpResponse<RestaurantDTOMessage>>("not-glovo/api/restaurants/updateRestaurant", restaurantDTO, queryParams);
    }

    deleteRestaurant(id: number): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.delete<HttpResponse<RestaurantDTOMessage>>("not-glovo/api/restaurants/deleteRestaurant/" + id, queryParams);
    }
}