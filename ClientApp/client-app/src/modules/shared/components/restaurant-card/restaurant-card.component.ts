import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { RestaurantDTO } from '../../models/RestaurantDTO';

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
    Latitude: 0
  }

  @Output() renderList: EventEmitter<any> = new EventEmitter();
  
  constructor() { }

  ngOnInit(): void {
  }

}
