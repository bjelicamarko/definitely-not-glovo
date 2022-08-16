import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { Report } from "../models/Report";

@Injectable({
    providedIn: 'root'
})
export class ReportsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    getReports():  Observable<HttpResponse<Report>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<Report>>("not-glovo/api/reports/getReports", queryParams);
    }
}