import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ArticleDTO } from 'src/modules/shared/models/ArticleDTO';
import { ArticleDTOMessage } from 'src/modules/shared/models/ArticleDTOMessage';
import { RestaurantDTO } from 'src/modules/shared/models/RestaurantDTO';
import { RestaurantsPageable } from 'src/modules/shared/models/RestaurantsPageable';
import { ArticlesUtilsService } from 'src/modules/shared/services/articles-utils';
import { RestaurantsUtilsService } from 'src/modules/shared/services/restaurants-utils';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { ArticlesService } from '../../services/articles.service';

@Component({
  selector: 'app-create-article-page',
  templateUrl: './create-article-page.component.html',
  styleUrls: ['./create-article-page.component.scss']
})
export class CreateArticlePageComponent implements OnInit {
  article: ArticleDTO = {
    Id: 0,
    ArticleName: '',
    ArticleType: '',
    Price: 0,
    Description: '',
    RestaurantName: '',
    Image: 'assets/article.png',
    ImagePath: 'assets/article.png',
    Changed: false
  }

  public selectedFile: File | undefined

  public articleIdFromRoute: number

  restaurants: RestaurantDTO[];

  constructor(private articlesService: ArticlesService,
    private articlesUtilsService: ArticlesUtilsService,
    private snackBarService: SnackBarService,
    private router: Router,
    private route: ActivatedRoute,
    private restaurantsUtilsService: RestaurantsUtilsService
  ) { 
    this.articleIdFromRoute = 0
    this.restaurants = [];
  }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.articleIdFromRoute = Number(routeParams.get('articleId'));

    if (this.articleIdFromRoute !== 0) {
      this.articlesUtilsService.findArticleById(this.articleIdFromRoute)
      .subscribe((response) => {
        var temp = response.body as ArticleDTOMessage;
        this.article = temp.ArticleDTO;
        this.snackBarService.openSnackBar(temp.Message);
      })
    }

    if (this.articleIdFromRoute === 0) {
      this.restaurantsUtilsService.findAllRestaurants(0, 100)
      .subscribe((response) => {
        var temp = response.body as RestaurantsPageable;
        this.restaurants = temp.Elements as RestaurantDTO[];
      })
    }
  }

  createArticle() {
    if (this.article.ArticleName && this.article.ArticleType && 
      this.article.Price && this.article.Description && 
      this.article.RestaurantName && this.article.Image && 
      this.article.ImagePath !== 'assets/article.png') {
        let reader = new FileReader();
        reader.readAsDataURL(this.selectedFile!);
        reader.onload = () => {
          this.article.Image = reader.result;
          this.articlesService.createArticle(this.article)
          .subscribe((response) => {
            var temp = response.body as ArticleDTOMessage;
            this.snackBarService.openSnackBar(temp.Message);
            this.article = temp.ArticleDTO;
            this.router.navigate(["/app/main/admin/articles"]);
          })
        };
        reader.onerror = function (error) {
          console.log('Error: ', error);
        };
    }
  }

  updateArticle() {
    if (this.article.ArticleName && this.article.ArticleType && 
      this.article.Price && this.article.Description && 
      this.article.RestaurantName && this.article.Image && 
      this.article.ImagePath !== 'assets/article.png') {
      if (this.article.Changed) {
        let reader = new FileReader();
        reader.readAsDataURL(this.selectedFile!);
        reader.onload = () => {
          this.article.Image = reader.result;
          this.articlesService.updateArticle(this.article)
          .subscribe((response) => {
            var temp = response.body as ArticleDTOMessage;
            this.snackBarService.openSnackBar(temp.Message);
            this.article = temp.ArticleDTO;
            this.router.navigate(["/app/main/admin/articles"]);
          })
        };
        reader.onerror = function (error) {
          console.log('Error: ', error);
        };
      } else {
        this.articlesService.updateArticle(this.article)
        .subscribe((response) => {
          var temp = response.body as ArticleDTOMessage;
          this.snackBarService.openSnackBar(temp.Message);
          this.article = temp.ArticleDTO;
          this.router.navigate(["/app/main/admin/articles"]);
        })
      }
    }
  }

  onFileChanged(event: any) {
    this.selectedFile = event.target.files[0]
    this.article.ImagePath = "images/" + this.selectedFile?.name as string;
    this.article.Changed = true
  }
}
