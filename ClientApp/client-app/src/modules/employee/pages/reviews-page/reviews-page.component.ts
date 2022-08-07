import { Component, OnInit, ViewChild } from '@angular/core';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { RestaurantDTOMessage } from 'src/modules/shared/models/RestaurantDTOMessage';
import { ReviewDTO } from 'src/modules/shared/models/ReviewDTO';
import { ReviewDTOMessage } from 'src/modules/shared/models/ReviewDTOMessage';
import { ReviewsPageable } from 'src/modules/shared/models/ReviewsPageable';
import { UserDTOMessage } from 'src/modules/shared/models/UserDTOMessage';
import { RestaurantsUtilsService } from 'src/modules/shared/services/restaurants-utils';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UsersUtilsService } from 'src/modules/shared/services/users-utils.service';
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
  
  idRestaurantForSearch: number;

  constructor(private reviewsService: ReviewsService,
    private authService: AuthService,
    private usersUtilsService: UsersUtilsService,
    private restaurantsUtilsService: RestaurantsUtilsService,
    private snackBarService: SnackBarService) { 
    this.reviews = []
    this.pageSize = 5;
    this.currentPage = 1;
    this.totalSize = 1;
    this.idRestaurantForSearch = 0;
  }

  ngOnInit(): void {
    var info = this.authService.getInfo();
    this.usersUtilsService.findUserById(info.Id)
    .subscribe((response) => {
      var temp = response.body as UserDTOMessage;
      this.restaurantsUtilsService.findRestaurantByName(temp.UserDTO.RestaurantName)
        .subscribe((response) => {
          var temp = response.body as RestaurantDTOMessage;
          this.idRestaurantForSearch = temp.RestaurantDTO.Id;
          this.reviewsService.getReviewsOfRestaurant(this.idRestaurantForSearch,
          this.currentPage - 1, this.pageSize)
          .subscribe((response) => {
            var temp = response.body as ReviewsPageable;
            this.totalSize = temp.TotalElements;
            this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
            this.reviews = temp.Elements as ReviewDTO[];
          })
      })
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
    this.reviewsService.getReviewsOfRestaurant(this.idRestaurantForSearch, 
      this.currentPage, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as ReviewsPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (this.currentPage - 1).toString());
        this.reviews = temp.Elements as ReviewDTO[];
      })
  }

  reportReview(review: ReviewDTO) {
    this.reviewsService.reportReview(review)
    .subscribe((response) => {
      var temp = response.body as ReviewDTOMessage
      this.snackBarService.openSnackBar(temp.Message)
      this.reset()
    })
  }
}
