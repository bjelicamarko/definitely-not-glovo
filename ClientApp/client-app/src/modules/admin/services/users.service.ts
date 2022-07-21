import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
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
}
