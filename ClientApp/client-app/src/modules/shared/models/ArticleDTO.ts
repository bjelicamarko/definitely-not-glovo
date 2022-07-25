export interface ArticleDTO {
    Id: number;
    ArticleName: string;
    ArticleType: string;
    Price: number;
    Description: string;
    RestaurantName: string;
    Image: string | ArrayBuffer | null;
    ImagePath: string;
    Changed: boolean;
}