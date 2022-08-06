import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { OrderInfoDialogComponent } from '../../components/order-info-dialog/order-info-dialog.component';
import { PaginationComponent } from '../../components/pagination/pagination.component';
import { OrderDTO } from '../../models/OrderDTO';
import { OrdersPageable } from '../../models/OrdersPageable';
import { RestaurantDTOMessage } from '../../models/RestaurantDTOMessage';
import { UserDTOMessage } from '../../models/UserDTOMessage';
import { OrdersUtilsService } from '../../services/orders-utils';
import { RestaurantsUtilsService } from '../../services/restaurants-utils';
import { UsersUtilsService } from '../../services/users-utils.service';

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

  idUserForSearch: number;
  idRestaurantForSearch: number;

  constructor(private fb: FormBuilder,
    private authService: AuthService,
    private ordersUtilsService: OrdersUtilsService,
    private usersUtilsService: UsersUtilsService,
    private restaurantsUtilsService: RestaurantsUtilsService,
    public dialog: MatDialog) { 
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

    this.idUserForSearch = 0
    this.idRestaurantForSearch = 0
  }

  ngOnInit(): void {
    var info = this.authService.getInfo();
    this.idUser = info.Id;
    this.userRole = info.role;
    this.idUserForSearch = this.idUser
    if (this.userRole === 'APPUSER') {
      this.getInitOrders();
    } else {
      this.usersUtilsService.findUserById(this.idUser)
      .subscribe((response) => {
        var temp = response.body as UserDTOMessage;
        this.restaurantsUtilsService.findRestaurantByName(temp.UserDTO.RestaurantName)
        .subscribe((response) => {
          var temp = response.body as RestaurantDTOMessage;
          this.idRestaurantForSearch = temp.RestaurantDTO.Id;
          this.getInitOrders();
        })
      })
    }
    
    this.onChanges();
  }

  getInitOrders() {
    this.ordersUtilsService.searchOrders(this.userRole, this.idUserForSearch, 
      this.idRestaurantForSearch, 'all', 0, 0, this.currentPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as OrdersPageable;
        this.totalSize = temp.TotalElements;
        this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
        this.orders = temp.Elements as OrderDTO[];
        console.log(this.orders);
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
    this.ordersUtilsService.searchOrders(this.userRole, this.idUserForSearch, this.idRestaurantForSearch, 
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
      (this.userRole, this.idUserForSearch, this.idRestaurantForSearch, val.orderStatus, val.priceFrom, val.priceTo,
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

  onUpdate(order: OrderDTO) {
    const dialogRef = this.dialog.open(OrderInfoDialogComponent, {
      data: {Id: order.Id, Role: this.userRole, IdUser: this.idUser},
      autoFocus: false,
      maxHeight: '90vh',
      width: '100%',
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result)
      order.OrderStatus = result.OrderStatus;
      order.IdDeliverer = result.IdDeliverer;
      order.IdEmployee = result.IdEmployee;
    });
  }
}
