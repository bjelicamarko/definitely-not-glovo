import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
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

    searchReviews(restaurantId: number, userId: number, inappropriate: string, 
        pageNum: number, pageSize: number)  : Observable<HttpResponse<ReviewsPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: {
            restaurantId: restaurantId,
            userId: userId,
            inappropriate: inappropriate,
            size: pageSize,
            page: pageNum
        }};

        return this.http.get<HttpResponse<ReviewsPageable>>(environment.url + "/api/reviews/searchReviews", queryParams);
    }

    deleteReview(id: number): Observable<HttpResponse<ReviewDTOMessage>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
        }

        return this.http.delete<HttpResponse<ReviewDTOMessage>>(environment.url + "/api/reviews/deleteReview/" + id, queryParams);
    }
}