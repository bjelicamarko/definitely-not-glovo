import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ReviewDTO } from '../../models/ReviewDTO';

@Component({
  selector: 'app-create-review-dialog',
  templateUrl: './create-review-dialog.component.html',
  styleUrls: ['./create-review-dialog.component.scss']
})
export class CreateReviewDialogComponent implements OnInit {

  comment: string;
  rating: number;

  constructor(public dialogRef: MatDialogRef<CreateReviewDialogComponent>,
  @Inject(MAT_DIALOG_DATA) public data: ReviewDTO, 
  ) { 
    this.comment = ""
    this.rating = 0
  }

  ngOnInit(): void {
  }

  reviewIt(): void {
    if (this.rating >= 0 && this.rating <= 10 && this.comment) {
      this.data.Comment = this.comment;
      this.data.Rating = this.rating;
      this.dialogRef.close();
    }
  }
}
