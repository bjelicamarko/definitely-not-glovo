import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ImageMessage } from "../models/ImageMessage";
import { RestaurantsPageable } from "../models/RestaurantsPageable";
import { UserDTO } from "../models/UserDTO";
import { UserDTOMessage } from "../models/UserDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class UtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}
    
    public getNoPages(totalItems: number, pageSize: number): number {
        return Math.ceil(totalItems / pageSize);
    }

    findUserById(id: number): Observable<HttpResponse<UserDTO>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<UserDTO>>("not-glovo/api/users/findUserById/" + id, queryParams);
    }

    saveImageUser(imageMessage: ImageMessage): Observable<HttpResponse<UserDTO>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<UserDTO>>("not-glovo/api/users/saveImageUser", imageMessage, queryParams);
    }

    updateUser(userDTO: UserDTO):  Observable<HttpResponse<UserDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.put<HttpResponse<UserDTOMessage>>("not-glovo/api/users/updateUser", userDTO, queryParams);
    }

    getRestaurants(page: number, size: number): Observable<HttpResponse<RestaurantsPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: new HttpParams()
            .set("page", String(page))
            .append("size", String(size))
        };

        return this.http.get<HttpResponse<RestaurantsPageable>>("not-glovo/api/restaurants/getRestaurants", queryParams);
    }
}