import { HttpHeaders, HttpClient, HttpParams, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ArticleDTOMessage } from "../models/ArticleDTOMessage";
import { ArticlesPageable } from "../models/ArticlesPageable";

@Injectable({
    providedIn: 'root'
})
export class ArticlesUtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    findAllArticles(page: number, size: number): Observable<HttpResponse<ArticlesPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: new HttpParams()
            .set("page", String(page))
            .append("size", String(size))
        };

        return this.http.get<HttpResponse<ArticlesPageable>>("not-glovo/api/articles/findAllArticles", queryParams);
    }

    findAllArticlesFromRestaurant(restaurantName: string, page: number, size: number): 
    Observable<HttpResponse<ArticlesPageable>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        params: new HttpParams()
            .set("restaurantName",restaurantName)
            .append("page", String(page))
            .append("size", String(size))
        };

        return this.http.get<HttpResponse<ArticlesPageable>>("not-glovo/api/articles/findAllArticlesFromRestaurant", queryParams);
    }

    searchArticles(restaurantName: string, searchField: string, articleType: string, 
        priceFrom: number, priceTo: number, page: number, size: number):
    Observable<HttpResponse<ArticlesPageable>> {
        let queryParams = {};

        if (!restaurantName)
            restaurantName = ''
        if (!searchField)
            searchField = ''
        if (!articleType || articleType === 'all')
            articleType = ''
        
        if (!priceFrom || priceFrom < 0)
            priceFrom = 0
        if (!priceTo || priceTo <= 0)
            priceTo = 10000
        
        queryParams = {
            headers: this.headers,
            observe: "response",
            params: {
                restaurantName: restaurantName,
                searchField: searchField,
                articleType: articleType,
                priceFrom: priceFrom,
                priceTo: priceTo,
                page: page,
                size: size,
        }};

        return this.http.get<HttpResponse<ArticlesPageable>>("not-glovo/api/articles/searchArticles", queryParams);
    }

    findArticleById(id: number): Observable<HttpResponse<ArticleDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.get<HttpResponse<ArticleDTOMessage>>("not-glovo/api/articles/findArticleById/" + id, queryParams);
    }
}