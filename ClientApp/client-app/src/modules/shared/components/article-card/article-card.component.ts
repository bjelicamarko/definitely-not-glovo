import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ArticlesService } from 'src/modules/admin/services/articles.service';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { ArticleDTO } from '../../models/ArticleDTO';
import { ArticleDTOMessage } from '../../models/ArticleDTOMessage';
import { OrderDTO } from '../../models/OrderDTO';
import { OrderItemDTO } from '../../models/OrderItemDTO';
import { ArticlesUtilsService } from '../../services/articles-utils';
import { SnackBarService } from '../../services/snack-bar.service';
import { ConformationDialogComponent } from '../conformation-dialog/conformation-dialog.component';
import { OrderitemInitDialogComponent } from '../orderitem-init-dialog/orderitem-init-dialog.component';

@Component({
  selector: 'app-article-card',
  templateUrl: './article-card.component.html',
  styleUrls: ['./article-card.component.scss']
})
export class ArticleCardComponent implements OnInit {

  @Input() article: ArticleDTO = {
    Id: 0,
    ArticleName: '',
    ArticleType: '',
    Price: 0,
    Description: '',
    RestaurantName: '',
    Image: null,
    ImagePath: '',
    Changed: false
  }
  
  @Output() renderList: EventEmitter<any> = new EventEmitter();
  @Output() sendItemToArticlesPage: EventEmitter<OrderItemDTO> = new EventEmitter();

  role: string = "";

  constructor(public dialog: MatDialog,
    private router: Router,
    private articlesService: ArticlesService,
    private authService: AuthService,
    private snackBarService: SnackBarService) { }

  ngOnInit(): void {
    var info = this.authService.getInfo();
    this.role = info.role;
  }

  updateArticle(id: number) {
    this.router.navigate(["/app/main/admin/article-info/" + id]);
  }

  orderArticle(any: ArticleDTO) {
    this.dialog.open(OrderitemInitDialogComponent, {
      data: any
    }).afterClosed().subscribe(result => {
      this.sendItemToArticlesPage.emit(result as OrderItemDTO);
    })
  }
  
  deleteArticle(id: number) {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Deleting article",
        body: "You want to remove " + this.article.ArticleName + "?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        this.articlesService.deleteArticle(id)
        .subscribe((response) => {
          var temp = response.body as ArticleDTOMessage;
          this.snackBarService.openSnackBar(temp.Message);
          this.renderList.emit(null);
        })
      }
    })
  }
}
