import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { RestaurantsService } from 'src/modules/admin/services/restaurants.service';
import { RestaurantDTO } from '../../models/RestaurantDTO';
import { RestaurantDTOMessage } from '../../models/RestaurantDTOMessage';
import { SnackBarService } from '../../services/snack-bar.service';
import { ConformationDialogComponent } from '../conformation-dialog/conformation-dialog.component';

@Component({
  selector: 'app-restaurant-card',
  templateUrl: './restaurant-card.component.html',
  styleUrls: ['./restaurant-card.component.scss']
})
export class RestaurantCardComponent implements OnInit {
  @Input() restaurant: RestaurantDTO = {
    Id: 0,
    RestaurantName: '',
    ContactPhone: '',
    Image: null,
    ImagePath: '',
    Country: '',
    City: '',
    Street: '',
    StreetNumber: '',
    Ptt: 0,
    DisplayName: '',
    Longitude: 0,
    Latitude: 0,
    Changed: false
  }

  @Output() renderList: EventEmitter<any> = new EventEmitter();
  
  constructor(public dialog: MatDialog,
    private router: Router,
    private restaurantsService: RestaurantsService,
    private snackBarService: SnackBarService) { }

  ngOnInit(): void {
  }

  restaurantInfo(id: number): void {
    // ovdje ce ici info stranica
    //this.router.navigate(["/app/main/admin/restaurant-info/" + id]);
  }

  updateRestaurant(id: number) {
    this.router.navigate(["/app/main/admin/restaurant-info/" + id]);
  }

  deleteRestaurant(id: number) {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Deleting restaurant",
        body: "You want to delete " + this.restaurant.RestaurantName + " restaurant?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        this.restaurantsService.deleteRestaurant(id)
        .subscribe((response) => {
          var temp = response.body as RestaurantDTOMessage;
          this.snackBarService.openSnackBar(temp.Message);
          this.renderList.emit(null);
        })
      }
    })
  }
}
