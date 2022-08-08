import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { ConformationDialogComponent } from 'src/modules/shared/components/conformation-dialog/conformation-dialog.component';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { ReviewDTO } from 'src/modules/shared/models/ReviewDTO';
import { ReviewDTOMessage } from 'src/modules/shared/models/ReviewDTOMessage';
import { ReviewsPageable } from 'src/modules/shared/models/ReviewsPageable';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { ReviewsService } from '../../services/reviews.service';

@Component({
  selector: 'app-reviews-page',
  templateUrl: './reviews-page.component.html',
  styleUrls: ['./reviews-page.component.scss']
})
export class ReviewsPageComponent implements OnInit {
  @ViewChild(PaginationComponent) pagination?: PaginationComponent;
  pageSize: number;
  currentPage: number;
  totalSize: number;
  reviews: ReviewDTO[]
  
  searchFormGroup: FormGroup;
  
  constructor(private fb: FormBuilder,
    private reviewsService: ReviewsService,
    private snackBarService: SnackBarService,
    public dialog: MatDialog) { 
    this.reviews = []
    this.pageSize = 5;
    this.currentPage = 1;
    this.totalSize = 1;
    this.searchFormGroup = this.fb.group({
      reported: [''],
    }); 
  }

  ngOnInit(): void {
    this.reviewsService.searchReviews(0, 0, "false", this.currentPage - 1, this.pageSize)
    .subscribe((response) => {
      var temp = response.body as ReviewsPageable;
      this.totalSize = temp.TotalElements;
      this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
      this.reviews = temp.Elements as ReviewDTO[];
    })

    this.onChanges();
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
    this.reviewsService.searchReviews(0, 0, this.searchFormGroup.value.reported, 
      this.currentPage, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as ReviewsPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (this.currentPage - 1).toString());
        this.reviews = temp.Elements as ReviewDTO[];
      })
  }

  onChanges(): void {
    this.searchFormGroup.valueChanges
    .subscribe(val => {
      this.reviewsService.searchReviews(0, 0, val.reported, 
        0, this.pageSize).subscribe((response) => {
          if (response.body != null) {
            var temp = response.body as ReviewsPageable;
            this.totalSize = Number(temp.TotalElements);
            this.setPagination((this.totalSize).toString(), (0).toString());
            this.reviews = temp.Elements as ReviewDTO[];
            if (this.pagination) {
              this.pagination.setActivePage(1);
            } 
          }
        })
    })
  }

  deleteReview(review: ReviewDTO) {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Removing Review",
        body: "You want to remove review?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        this.reviewsService.deleteReview(review.Id)
        .subscribe((response) => {
          var temp = response.body as ReviewDTOMessage
          this.snackBarService.openSnackBar(temp.Message)
          this.reset()
        })
      }
    })
  }
}
