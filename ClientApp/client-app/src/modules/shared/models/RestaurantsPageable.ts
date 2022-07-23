import { RestaurantDTO } from "./RestaurantDTO";

export interface RestaurantsPageable {
    Elements: RestaurantDTO[];
    TotalElements: number;
}