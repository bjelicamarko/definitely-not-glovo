export interface UserDTO {
    Id: number;
    Email: string;
    Password: string;
    FirstName: string;
    LastName: string;
    Contact: string;
    Role: string;
    Banned: boolean;
    Image: string | ArrayBuffer | null;
    ImagePath: string;
    Changed: boolean;
    RestaurantName: string;
}