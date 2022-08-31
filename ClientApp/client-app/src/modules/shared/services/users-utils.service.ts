import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { UserDTO } from "../models/UserDTO";
import { UserDTOMessage } from "../models/UserDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class UsersUtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    findUserById(id: number): Observable<HttpResponse<UserDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<UserDTOMessage>>(environment.url + "/api/users/findUserById/" + id, queryParams);
    }

    updateUser(userDTO: UserDTO): Observable<HttpResponse<UserDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.put<HttpResponse<UserDTOMessage>>(environment.url + "/api/users/updateUser", userDTO, queryParams); 
    }
}