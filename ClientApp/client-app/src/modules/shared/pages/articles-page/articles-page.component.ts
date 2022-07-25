import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { PaginationComponent } from '../../components/pagination/pagination.component';
import { ArticleDTO } from '../../models/ArticleDTO';
import { ArticlesPageable } from '../../models/ArticlesPageable';
import { ArticlesUtilsService } from '../../services/articles-utils';

@Component({
  selector: 'app-articles-page',
  templateUrl: './articles-page.component.html',
  styleUrls: ['./articles-page.component.scss']
})
export class ArticlesPageComponent implements OnInit {
  @ViewChild(PaginationComponent) pagination?: PaginationComponent;
  pageSize: number;
  currentPage: number;
  totalSize: number;
  articles: ArticleDTO[];

  searchFormGroup: FormGroup;
  
  constructor(private fb: FormBuilder,
    private articlesUtilsService: ArticlesUtilsService) {
      this.articles = []
      this.pageSize = 5;
      this.currentPage = 1;
      this.totalSize = 1;
      this.searchFormGroup = this.fb.group({
        restaurantName: [''],
        searchField: [''],
        articleType: [''],
        priceFrom: [''],
        priceTo: ['']
      }); 
  }

  ngOnInit(): void {
    this.articlesUtilsService.findAllArticles(this.currentPage - 1, this.pageSize)
    .subscribe((response) => {
      var temp = response.body as ArticlesPageable;
      this.totalSize = Number(temp.TotalElements)
      this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
      this.articles = temp.Elements as ArticleDTO[];
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
    this.articlesUtilsService.searchArticles(this.searchFormGroup.value.restaurantName, 
      this.searchFormGroup.value.searchField, this.searchFormGroup.value.articleType, 
      this.searchFormGroup.value.priceFrom, this.searchFormGroup.value.priceTo, newPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as ArticlesPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (newPage - 1).toString());
        this.articles = temp.Elements as ArticleDTO[];
      })
  }

  renderList() {
    this.articlesUtilsService.findAllArticles(this.currentPage - 1, this.pageSize)
    .subscribe((response) => {
      var temp = response.body as ArticlesPageable;
      this.totalSize = Number(temp.TotalElements)
      this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
      this.articles = temp.Elements as ArticleDTO[];
    })
  }

  onChanges(): void {
    this.searchFormGroup.valueChanges
    .subscribe(val => {
      if (!this.searchFormGroup.get('restaurantName')?.valid ||
      !this.searchFormGroup.get('searchField')?.valid ||
      !this.searchFormGroup.get('articleType')?.valid ||
      !this.searchFormGroup.get('priceFrom')?.valid ||
      !this.searchFormGroup.get('priceTo')?.valid)
        return;

      this.articlesUtilsService.searchArticles(val.restaurantName, 
        val.searchField, val.articleType, val.priceFrom, val.priceTo, 0, this.pageSize)
        .subscribe((response) => {
          if (response.body != null) {
            var temp = response.body as ArticlesPageable;
            this.totalSize = temp.TotalElements;
            this.setPagination((this.totalSize).toString(), (0).toString());
            this.articles = temp.Elements as ArticleDTO[];
            if (this.pagination) {
              this.pagination.setActivePage(1);
            }
          }
        })
    })
  }
}
