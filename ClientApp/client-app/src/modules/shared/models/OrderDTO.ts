import { OrderItemDTO } from "./OrderItemDTO";

export interface OrderDTO {
    Id: number;
    IdRestaurant: number;
    IdAppUser: number;
    IdEmployee: number;
    IdDeliverer: number;
    OrderStatus: string;
    TotalPrice: number;
    Tip: number;
    Note: string;
    DateTime: string;
    OrderItemsDTO: OrderItemDTO[]
}