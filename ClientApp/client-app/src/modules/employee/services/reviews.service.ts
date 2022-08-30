import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { ReviewDTO } from "src/modules/shared/models/ReviewDTO";
import { ReviewDTOMessage } from "src/modules/shared/models/ReviewDTOMessage";
import { ReviewsPageable } from "src/modules/shared/models/ReviewsPageable";

@Injectable({
    providedIn: 'root'
})
export class ReviewsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    getReviewsOfRestaurant(idRestaurant: number, pageNum: number, pageSize: number)
    : Observable<HttpResponse<ReviewsPageable>> {

        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: {
            restaurantId: idRestaurant,
            size: pageSize,
            page: pageNum
        }};

        return this.http.get<HttpResponse<ReviewsPageable>>(environment.url + "/api/reviews/getReviewsOfRestaurant", queryParams);
    }

    reportReview(reviewDTO: ReviewDTO): Observable<HttpResponse<ReviewDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.put<HttpResponse<ReviewDTOMessage>>(environment.url + "/api/reviews/reportReview", reviewDTO, queryParams); 
    }
}