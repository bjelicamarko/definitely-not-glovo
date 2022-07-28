import { HttpHeaders, HttpClient, HttpResponse } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { OrderDTOMessage } from "../models/OrderDTOMessage";
import { OrdersPageable } from "../models/OrdersPageable"
import { OrderStatusDTO } from "../models/OrderStatusDTO";

@Injectable({
    providedIn: 'root'
})
export class OrdersUtilsService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });
    
    constructor(private http: HttpClient) {}

    searchOrders(userId: number, restaurantId: number, 
        orderStatus: string, priceFrom: number, priceTo: number,
        page: number, size: number)
    : Observable<HttpResponse<OrdersPageable>>  {
        let queryParams = {};

        if (!orderStatus || orderStatus === 'all')
            orderStatus = ''

        if (!priceFrom || priceFrom < 0)
            priceFrom = 0
        if (!priceTo || priceTo <= 0)
            priceTo = 10000

        queryParams = {
            headers: this.headers,
            observe: "response",
            params: {
                userId: userId,
                restaurantId: restaurantId,
                orderStatus: orderStatus,
                priceFrom: priceFrom,
                priceTo: priceTo,
                page: page,
                size: size
        }
        };

        return this.http.get<HttpResponse<OrdersPageable>>
        ("not-glovo/api/orders/searchOrders", queryParams);
    }

    findOrderById(id: number): Observable<HttpResponse<OrderDTOMessage>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
        }

        return this.http.get<HttpResponse<OrderDTOMessage>>
        ("not-glovo/api/orders/findOrderById/" + id, queryParams);
    }

    changeStatusOfOrder(orderStatusDTO: OrderStatusDTO): Observable<HttpResponse<OrderDTOMessage>> {
        let queryParams = {};

        queryParams = {
            headers: this.headers,
            observe: "response",
        }

        return this.http.put<HttpResponse<OrderDTOMessage>>
        ("not-glovo/api/orders/changeStatusOfOrder", orderStatusDTO,  queryParams);
    }
}