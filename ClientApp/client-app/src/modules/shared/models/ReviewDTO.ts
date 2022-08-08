export interface ReviewDTO {
    Id: number;
    Comment: string;
    Rating: number;
    InappropriateContent: boolean;
    DateTime: string;
    IdRestaurant: number;
    IdOrder: number;
    IdUser: number;
    EmailUser: string;
}