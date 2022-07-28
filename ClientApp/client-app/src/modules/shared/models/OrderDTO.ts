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
    Country: string;
    City: string;
    Street: string;
    StreetNumber: string;
    Ptt: number;
    DisplayName: string;
    Longitude: number;
    Latitude: number;
    OrderItemsDTO: OrderItemDTO[];
}