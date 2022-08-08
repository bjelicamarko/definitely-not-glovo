import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { ReviewDTO } from '../../models/ReviewDTO';

@Component({
  selector: 'app-reviews-view',
  templateUrl: './reviews-view.component.html',
  styleUrls: ['./reviews-view.component.scss']
})
export class ReviewsViewComponent implements OnInit {
  @Input() reviews: ReviewDTO[] = []

  @Output() sendReviewToParent = new EventEmitter<ReviewDTO>();
  
  role: string;
  constructor(private authService: AuthService) {
    this.role = ""
  }

  ngOnInit(): void {
    var info = this.authService.getInfo();
    this.role = info.role;
  }

  reportReview(review: ReviewDTO) {
    console.log(review);
    this.sendReviewToParent.emit(review)
  }

  deleteReview(review: ReviewDTO) {
    this.sendReviewToParent.emit(review)
  }
}
