import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { OrderDTOMessage } from "../models/OrderDTOMessage";
import { ReviewDTO } from "../models/ReviewDTO";
import { ReviewDTOMessage } from "../models/ReviewDTOMessage";
import { ReviewsPageable } from "../models/ReviewsPageable";

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

    averageRatingOfRestaurant(idRestaurant: number) : Observable<HttpResponse<number>>{
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<number>>("not-glovo/api/reviews/averageRatingOfRestaurant/" + idRestaurant, queryParams);
    }

    createReview(reviewDTO: ReviewDTO): Observable<HttpResponse<ReviewDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<ReviewDTOMessage>>("not-glovo/api/reviews/createReview", reviewDTO, queryParams);
    }

}