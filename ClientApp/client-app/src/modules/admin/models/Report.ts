import { ArticlePriceQuantity } from "./ArticlePriceQuantity";

export interface Report {
    map_restaurants: Map<string, number>;
    map_articles: Map<string, ArticlePriceQuantity>;
    date: string;
}