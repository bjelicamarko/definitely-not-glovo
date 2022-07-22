import { HttpClient, HttpHeaders, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ImageMessage } from "../models/ImageMessage";
import { UserDTO } from "../models/UserDTO";

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
}