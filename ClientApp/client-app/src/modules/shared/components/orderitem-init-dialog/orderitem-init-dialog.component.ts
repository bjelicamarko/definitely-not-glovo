import { Component, Inject, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ArticleDTO } from '../../models/ArticleDTO';
import { OrderItemDTO } from '../../models/OrderItemDTO';
import { ConformationDialogComponent } from '../conformation-dialog/conformation-dialog.component';

@Component({
  selector: 'app-orderitem-init-dialog',
  templateUrl: './orderitem-init-dialog.component.html',
  styleUrls: ['./orderitem-init-dialog.component.scss']
})
export class OrderitemInitDialogComponent implements OnInit {
  
  orderItem: OrderItemDTO = {
    Id: 0,
    IdOrder: 0,
    IdArticle: 0,
    ArticleName: '',
    CurrentPrice: 0,
    Quantity: 0,
    TotalPrice: 0
  }
  
  constructor(public dialogRef: MatDialogRef<OrderitemInitDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: ArticleDTO,
    public dialog: MatDialog) {
      dialogRef.beforeClosed().subscribe(() => dialogRef.close(this.orderItem));
  }

  ngOnInit(): void {

  }

  decrementQuantity() {
    if (this.orderItem.Quantity > 0) {
      this.orderItem.Quantity = this.orderItem.Quantity - 1
      this.orderItem.TotalPrice = this.orderItem.TotalPrice - this.data.Price
    }
  }

  incrementQuanitity() {
    this.orderItem.Quantity = this.orderItem.Quantity + 1
    this.orderItem.TotalPrice = this.orderItem.TotalPrice + this.data.Price
  }

  addItemInOrder() {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Ordering item",
        body: "You want to order " + this.data.ArticleName + "?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        this.orderItem.IdArticle = this.data.Id
        this.orderItem.ArticleName = this.data.ArticleName
        this.orderItem.CurrentPrice = this.data.Price
        this.dialogRef.close()
      }
    })
  }
}
