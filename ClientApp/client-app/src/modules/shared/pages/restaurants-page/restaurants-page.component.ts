import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { PaginationComponent } from '../../components/pagination/pagination.component';
import { RestaurantDTO } from '../../models/RestaurantDTO';
import { RestaurantsPageable } from '../../models/RestaurantsPageable';
import { UtilsService } from '../../services/utils.service';

@Component({
  selector: 'app-restaurants-page',
  templateUrl: './restaurants-page.component.html',
  styleUrls: ['./restaurants-page.component.scss']
})
export class RestaurantsPageComponent implements OnInit {
  @ViewChild(PaginationComponent) pagination?: PaginationComponent;
  pageSize: number;
  currentPage: number;
  totalSize: number;
  restaurants: RestaurantDTO[];

  searchFormGroup: FormGroup;
  
  constructor(private fb: FormBuilder,
    private utilsService: UtilsService) {
      this.restaurants = []
      this.pageSize = 5;
      this.currentPage = 1;
      this.totalSize = 1;
      this.searchFormGroup = this.fb.group({
        searchField: [''],
      }); 
  }

  ngOnInit(): void {
    this.utilsService.getRestaurants(this.currentPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as RestaurantsPageable;
        console.log(temp);
        this.totalSize = Number(temp.TotalElements);
        this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
        this.restaurants = temp.Elements as RestaurantDTO[];
    });

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
  }

  renderList() {
  }

  onChanges(): void {
  }
}
