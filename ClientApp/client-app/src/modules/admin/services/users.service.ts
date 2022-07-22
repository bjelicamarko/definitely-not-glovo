import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ResponseMessage } from "src/modules/shared/models/ResponseMessage";
import { UsersPageable } from "src/modules/shared/models/UsersPageable";

@Injectable({
    providedIn: 'root'
})
export class UsersService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {}

    getUsers(page: number, size: number): Observable<HttpResponse<UsersPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: new HttpParams()
            .set("page", String(page))
            .append("size", String(size))
        };

        return this.http.get<HttpResponse<UsersPageable>>("not-glovo/api/users/getUsers", queryParams);
    }

    searchUsers(searchFieldVal: string, userTypeVal: string, 
        pageNum: number, pageSize: number): Observable<HttpResponse<UsersPageable>> {

        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: {
            searchField: searchFieldVal,
            userType: userTypeVal,
            size: pageSize,
            page: pageNum
        }};

        return this.http.get<HttpResponse<UsersPageable>>("not-glovo/api/users/searchUsers", queryParams);
    }

    banUser(id: number): Observable<HttpResponse<ResponseMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.patch<HttpResponse<ResponseMessage>>("not-glovo/api/users/banUser/" + id, null, queryParams);        
    }

    unbanUser(id: number): Observable<HttpResponse<ResponseMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.patch<HttpResponse<ResponseMessage>>("not-glovo/api/users/unbanUser/" + id, null, queryParams);        
    }

    deleteUser(id: number): Observable<HttpResponse<ResponseMessage>>{
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.delete<HttpResponse<ResponseMessage>>("not-glovo/api/users/deleteUser/" + id, queryParams); 
    }
}
