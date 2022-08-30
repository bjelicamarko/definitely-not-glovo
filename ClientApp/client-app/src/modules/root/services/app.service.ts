import { HttpClient, HttpHeaders, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { ResponseMessage } from "src/modules/shared/models/ResponseMessage";
import { UserDTO } from "src/modules/shared/models/UserDTO";

@Injectable({
    providedIn: "root",
})
export class AppService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    register(userDTO: UserDTO): Observable<HttpResponse<ResponseMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.post<HttpResponse<ResponseMessage>>(environment.url + "/api/users/register", 
        userDTO, queryParams);
    }
}