import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { OrderDTOMessage } from "../models/OrderDTOMessage";
import { ReviewDTO } from "../models/ReviewDTO";
import { ReviewDTOMessage } from "../models/ReviewDTOMessage";
import { ReviewsPageable } from "../models/ReviewsPageable";

@Injectable({
    providedIn: 'root'
})
export class ReviewsUtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    findReviewByOrder(id: number): Observable<HttpResponse<ReviewDTOMessage>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
        }

        return this.http.get<HttpResponse<ReviewDTOMessage>>(environment.url + "/api/reviews/findReviewByOrder/" + id, queryParams);
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

    averageRatingOfRestaurant(idRestaurant: number) : Observable<HttpResponse<number>>{
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<number>>(environment.url + "/api/reviews/averageRatingOfRestaurant/" + idRestaurant, queryParams);
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