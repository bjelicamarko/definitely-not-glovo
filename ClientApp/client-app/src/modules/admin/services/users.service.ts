import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { UserDTO } from "src/modules/shared/models/UserDTO";
import { UserDTOMessage } from "src/modules/shared/models/UserDTOMessage";
import { UsersPageable } from "src/modules/shared/models/UsersPageable";

@Injectable({
    providedIn: 'root'
})
export class UsersService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    findAllUsers(page: number, size: number): Observable<HttpResponse<UsersPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: new HttpParams()
            .set("page", String(page))
            .append("size", String(size))
        };

        return this.http.get<HttpResponse<UsersPageable>>(environment.url + "/api/users/findAllUsers", queryParams);
    }

    searchUsers(searchFieldVal: string, userTypeVal: string, 
        pageNum: number, pageSize: number): Observable<HttpResponse<UsersPageable>> {
        
        if (!searchFieldVal)
            searchFieldVal = ''
        if (!userTypeVal || userTypeVal === 'all')
            userTypeVal = ''

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

        return this.http.get<HttpResponse<UsersPageable>>(environment.url + "/api/users/searchUsers", queryParams);
    }

    createUser(userDTO: UserDTO): Observable<HttpResponse<UserDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.post<HttpResponse<UserDTOMessage>>(environment.url + "/api/users/createUser", userDTO, queryParams); 
    }

    deleteUser(id: number): Observable<HttpResponse<UserDTOMessage>>{
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.delete<HttpResponse<UserDTOMessage>>(environment.url + "/api/users/deleteUser/" + id, queryParams); 
    }
    
    banUser(id: number): Observable<HttpResponse<UserDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.patch<HttpResponse<UserDTOMessage>>(environment.url + "/api/users/banUser/" + id, null, queryParams);        
    }

    unbanUser(id: number): Observable<HttpResponse<UserDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.patch<HttpResponse<UserDTOMessage>>(environment.url + "/api/users/unbanUser/" + id, null, queryParams);        
    }

}
