import { HttpHeaders, HttpClient, HttpParams, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { RestaurantDTOMessage } from "../models/RestaurantDTOMessage";
import { RestaurantsPageable } from "../models/RestaurantsPageable";

@Injectable({
    providedIn: 'root'
})
export class RestaurantsUtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    findAllRestaurants(page: number, size: number): Observable<HttpResponse<RestaurantsPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: new HttpParams()
            .set("page", String(page))
            .append("size", String(size))
        };

        return this.http.get<HttpResponse<RestaurantsPageable>>(environment.url + "/api/restaurants/findAllRestaurants", queryParams);
    }

    searchRestaurants(searchFieldVal: string, pageNum: number, pageSize: number): Observable<HttpResponse<RestaurantsPageable>> {
        if (!searchFieldVal)
            searchFieldVal = ''
        
        let queryParams = {};
        queryParams = {
            headers: this.headers,
            observe: "response",
            params: {
                searchField: searchFieldVal,
                size: pageSize,
                page: pageNum
        }};

        return this.http.get<HttpResponse<RestaurantsPageable>>(environment.url + "/api/restaurants/searchRestaurants", queryParams);
    }

    findRestaurantById(id: number): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<RestaurantDTOMessage>>(environment.url + "/api/restaurants/findRestaurantById/" + id, queryParams);
    }

    findRestaurantByName(name: string): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<RestaurantDTOMessage>>(environment.url + "/api/restaurants/findRestaurantByName/" + name, queryParams);
    }
}
