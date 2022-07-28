import { ArticleDTO } from "./ArticleDTO";

export interface ArticlesPageable {
    Elements: ArticleDTO[];
    TotalElements: number;
}