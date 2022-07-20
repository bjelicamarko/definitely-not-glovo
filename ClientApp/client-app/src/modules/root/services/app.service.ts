import { HttpClient, HttpHeaders, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { UserDTO } from "src/modules/shared/models/UserDTO";

@Injectable({
    providedIn: "root",
})
export class AppService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {}

    register(userDTO: UserDTO): Observable<HttpResponse<string>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        responseType: "text"
        };

        return this.http.post<HttpResponse<string>>("not-glovo/api/users/register", 
        userDTO, queryParams);
    }
}