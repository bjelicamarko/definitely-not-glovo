import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { OrderDTO } from "src/modules/shared/models/OrderDTO";
import { OrderDTOMessage } from "src/modules/shared/models/OrderDTOMessage";

@Injectable({
    providedIn: 'root'
})
export class OrdersService{
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    createOrder(orderDTO: OrderDTO): Observable<HttpResponse<OrderDTOMessage>> {
        let queryParams = {};

        queryParams = {
        headers: this.headers,
        observe: "response",
        }

        return this.http.post<HttpResponse<OrderDTOMessage>>("not-glovo/api/orders/createOrder", orderDTO, queryParams);
    }
}