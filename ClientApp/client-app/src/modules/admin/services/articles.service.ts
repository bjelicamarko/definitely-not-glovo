import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { ArticleDTO } from "src/modules/shared/models/ArticleDTO";
import { ArticleDTOMessage } from "src/modules/shared/models/ArticleDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class ArticlesService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    createArticle(articleDTO: ArticleDTO): Observable<HttpResponse<ArticleDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<ArticleDTOMessage>>("not-glovo/api/articles/createArticle", articleDTO, queryParams);
    } 

    updateArticle(articleDTO: ArticleDTO): Observable<HttpResponse<ArticleDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.put<HttpResponse<ArticleDTOMessage>>("not-glovo/api/articles/updateArticle", articleDTO, queryParams);
    } 

    deleteArticle(id: number): Observable<HttpResponse<ArticleDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.delete<HttpResponse<ArticleDTOMessage>>("not-glovo/api/articles/deleteArticle/" + id, queryParams);
    } 
}