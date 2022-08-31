import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { ReviewDTO } from "src/modules/shared/models/ReviewDTO";
import { ReviewDTOMessage } from "src/modules/shared/models/ReviewDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class ReviewsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    createReview(reviewDTO: ReviewDTO): Observable<HttpResponse<ReviewDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<ReviewDTOMessage>>(environment.url + "/api/reviews/createReview", reviewDTO, queryParams);
    }

}