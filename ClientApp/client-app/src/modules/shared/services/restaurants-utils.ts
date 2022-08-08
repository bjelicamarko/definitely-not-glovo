import { HttpHeaders, HttpClient, HttpParams, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { RestaurantDTOMessage } from "../models/RestaurantDTOMessage";
import { RestaurantsPageable } from "../models/RestaurantsPageable";

@Injectable({
    providedIn: 'root'
})
export class RestaurantsUtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    findAllRestaurants(page: number, size: number): Observable<HttpResponse<RestaurantsPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: new HttpParams()
            .set("page", String(page))
            .append("size", String(size))
        };

        return this.http.get<HttpResponse<RestaurantsPageable>>("not-glovo/api/restaurants/findAllRestaurants", queryParams);
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

        return this.http.get<HttpResponse<RestaurantsPageable>>("not-glovo/api/restaurants/searchRestaurants", queryParams);
    }

    findRestaurantById(id: number): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<RestaurantDTOMessage>>("not-glovo/api/restaurants/findRestaurantById/" + id, queryParams);
    }

    findRestaurantByName(name: string): Observable<HttpResponse<RestaurantDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<RestaurantDTOMessage>>("not-glovo/api/restaurants/findRestaurantByName/" + name, queryParams);
    }
}
