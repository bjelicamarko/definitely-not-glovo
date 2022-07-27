import { OrderDTO } from "./OrderDTO";

export interface OrdersPageable {
    Elements: OrderDTO[];
    TotalElements: number;
}