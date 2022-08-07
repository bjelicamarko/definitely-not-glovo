import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ReviewDTO } from "src/modules/shared/models/ReviewDTO";
import { ReviewDTOMessage } from "src/modules/shared/models/ReviewDTOMessage";
import { ReviewsPageable } from "src/modules/shared/models/ReviewsPageable";

@Injectable({
    providedIn: 'root'
})
export class ReviewsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {}

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

        return this.http.get<HttpResponse<ReviewsPageable>>("not-glovo/api/reviews/getReviewsOfRestaurant", queryParams);
    }

    reportReview(reviewDTO: ReviewDTO): Observable<HttpResponse<ReviewDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        };

        return this.http.put<HttpResponse<ReviewDTOMessage>>("not-glovo/api/reviews/reportReview", reviewDTO, queryParams); 
    }
}