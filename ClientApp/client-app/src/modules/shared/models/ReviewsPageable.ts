import { ReviewDTO } from "./ReviewDTO";

export interface ReviewsPageable {
    Elements: ReviewDTO[];
    TotalElements: number;
}