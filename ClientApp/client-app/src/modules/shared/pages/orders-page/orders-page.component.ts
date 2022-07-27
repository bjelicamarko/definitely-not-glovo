import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { PaginationComponent } from '../../components/pagination/pagination.component';
import { OrderDTO } from '../../models/OrderDTO';
import { OrdersPageable } from '../../models/OrdersPageable';
import { OrdersUtilsService } from '../../services/orders-utils';

@Component({
  selector: 'app-orders-page',
  templateUrl: './orders-page.component.html',
  styleUrls: ['./orders-page.component.scss']
})
export class OrdersPageComponent implements OnInit {
  @ViewChild(PaginationComponent) pagination?: PaginationComponent;
  pageSize: number;
  currentPage: number;
  totalSize: number;
  orders: OrderDTO[];

  searchFormGroup: FormGroup;

  idUser: number;
  userRole: string;

  delivererStatuses: string[] = ['all', 'READY', 'TAKEN', 'DELIVERED'];
  employeeStatuses: string[] = ['all', 'ORDERED', 'ACCEPTED', 'READY', 'TAKEN', 'DELIVERED'];
  appuserStatuses: string[] = ['all', 'ORDERED', 'ACCEPTED', 'READY', 'TAKEN', 'DELIVERED', 'CANCELLED'];

  constructor(private fb: FormBuilder,
    private authService: AuthService,
    private ordersUtilsService: OrdersUtilsService) { 
    this.orders = []
    this.pageSize = 5;
    this.currentPage = 1;
    this.totalSize = 1;
    this.idUser = 0;
    this.userRole = '';
    this.searchFormGroup = this.fb.group({
      orderStatus: [''],
      priceFrom: [''],
      priceTo: [''],
    }); 
  }

  ngOnInit(): void {
    var info = this.authService.getInfo();
    this.idUser = info.Id;
    this.userRole = info.role;
    this.ordersUtilsService.searchOrders(this.idUser, 
      0, 'all', 0, 0, this.currentPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as OrdersPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
        this.orders = temp.Elements as OrderDTO[];
        console.log(this.orders);
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
    this.ordersUtilsService.searchOrders(this.idUser, 0, 
      this.searchFormGroup.value.orderStatus,
      this.searchFormGroup.value.priceFrom, this.searchFormGroup.value.priceTo, 
      newPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as OrdersPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (newPage - 1).toString());
        this.orders = temp.Elements as OrderDTO[];
      })
  }

  onChanges(): void {
    this.searchFormGroup.valueChanges
    .subscribe(val => {
      if (!this.searchFormGroup.get('orderStatus')?.valid ||
      !this.searchFormGroup.get('priceFrom')?.valid ||
      !this.searchFormGroup.get('priceTo')?.valid)
        return;
      
      this.ordersUtilsService.searchOrders
      (this.idUser, 0, val.orderStatus, val.priceFrom, val.priceTo,
        0, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as OrdersPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (0).toString());
        this.orders = temp.Elements as OrderDTO[];
        if (this.pagination) {
          this.pagination.setActivePage(1);
        }
      })   
    })
  }
}
