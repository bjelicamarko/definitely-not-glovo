import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ReviewDTOMessage } from "../models/ReviewDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class ReviewsUtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    findReviewByOrder(id: number): Observable<HttpResponse<ReviewDTOMessage>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
        }

        return this.http.get<HttpResponse<ReviewDTOMessage>>("not-glovo/api/reviews/findReviewByOrder/" + id, queryParams);
    }
}