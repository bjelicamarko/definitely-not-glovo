import { Component, Inject, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ReviewDTO } from '../../models/ReviewDTO';
import { ReviewsPageable } from '../../models/ReviewsPageable';
import { ReviewsUtilsService } from '../../services/reviews-utils';
import { PaginationComponent } from '../pagination/pagination.component';

@Component({
  selector: 'app-reviews-dialog',
  templateUrl: './reviews-dialog.component.html',
  styleUrls: ['./reviews-dialog.component.scss']
})
export class ReviewsDialogComponent implements OnInit {
  @ViewChild(PaginationComponent) pagination?: PaginationComponent;
  pageSize: number;
  currentPage: number;
  totalSize: number;
  reviews: ReviewDTO[]
  
  avg: number;

  constructor(public dialogRef: MatDialogRef<ReviewsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: number,
    private reviewsUtilsService: ReviewsUtilsService) { 
    this.reviews = []
    this.pageSize = 5;
    this.currentPage = 1;
    this.totalSize = 1;
    this.avg = 0;
  }

  ngOnInit(): void {
    this.reviewsUtilsService.getReviewsOfRestaurant(this.data,
      this.currentPage - 1, this.pageSize)
    .subscribe((response) => {
      var temp = response.body as ReviewsPageable;
      this.totalSize = temp.TotalElements;
      this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
      this.reviews = temp.Elements as ReviewDTO[];
    })

    this.reviewsUtilsService.averageRatingOfRestaurant(this.data)
    .subscribe((response) => {
      this.avg = response.body as number;
    })
  }

  setPagination(totalItemsHeader: string | null, currentPageHeader: string | null) {
    if (totalItemsHeader) {
      this.totalSize = parseInt(totalItemsHeader);
    }
    if (currentPageHeader) {
      this.currentPage = parseInt(currentPageHeader);
    }
  }

  changePage(newPage: number) {
    this.currentPage = newPage
    this.reset()
  }

  reset() {
    this.reviewsUtilsService.getReviewsOfRestaurant(this.data, 
      this.currentPage, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as ReviewsPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (this.currentPage - 1).toString());
        this.reviews = temp.Elements as ReviewDTO[];
      })
  }
}
