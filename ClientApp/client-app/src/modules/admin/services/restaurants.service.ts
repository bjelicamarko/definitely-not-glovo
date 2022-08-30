import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { RestaurantDTO } from "src/modules/shared/models/RestaurantDTO";
import { RestaurantDTOMessage } from "src/modules/shared/models/RestaurantDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class RestaurantsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    createRestaurant(restaurantDTO: RestaurantDTO): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<RestaurantDTOMessage>>(environment.url + "/api/restaurants/createRestaurant", restaurantDTO, queryParams);
    }

    updateRestaurant(restaurantDTO: RestaurantDTO): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.put<HttpResponse<RestaurantDTOMessage>>(environment.url + "/api/restaurants/updateRestaurant", restaurantDTO, queryParams);
    }

    deleteRestaurant(id: number): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.delete<HttpResponse<RestaurantDTOMessage>>(environment.url + "/api/restaurants/deleteRestaurant/" + id, queryParams);
    }
}