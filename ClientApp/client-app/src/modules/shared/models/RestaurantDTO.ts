export interface RestaurantDTO {
    Id: number;
    RestaurantName: string;
    ContactPhone: string;
    Image: string | ArrayBuffer | null;
    ImagePath: string;
    Country: string;
    City: string;
    Street: string;
    StreetNumber: string;
    Ptt: number;
    DisplayName: string;
    Longitude: number;
    Latitude: number;
    Changed: boolean;
}