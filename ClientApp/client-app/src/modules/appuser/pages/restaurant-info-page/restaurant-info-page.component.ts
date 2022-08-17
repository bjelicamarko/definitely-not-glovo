import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Feature, View } from 'ol';
import { Geometry, Point } from 'ol/geom';
import TileLayer from 'ol/layer/Tile';
import Map from 'ol/Map';
import { fromLonLat, transform } from 'ol/proj';
import OSM from 'ol/source/OSM';
import Icon from 'ol/style/Icon';
import Style from 'ol/style/Style';
import Vector from 'ol/layer/Vector';
import VectorTemp from 'ol/source/Vector';
import { RestaurantDTO } from 'src/modules/shared/models/RestaurantDTO';
import { RestaurantDTOMessage } from 'src/modules/shared/models/RestaurantDTOMessage';
import { RestaurantsUtilsService } from 'src/modules/shared/services/restaurants-utils';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { ArticlesPageComponent } from 'src/modules/shared/pages/articles-page/articles-page.component';
import { OrderDTO } from 'src/modules/shared/models/OrderDTO';
import { OrderItemDTO } from 'src/modules/shared/models/OrderItemDTO';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { MatDialog } from '@angular/material/dialog';
import { ConformationDialogComponent } from 'src/modules/shared/components/conformation-dialog/conformation-dialog.component';
import { HttpClient } from '@angular/common/http';
import { MapAddress } from 'src/modules/shared/models/MapAddress';
import * as moment from 'moment';
import { OrdersService } from '../../services/orders.service';
import { OrderDTOMessage } from 'src/modules/shared/models/OrderDTOMessage';

@Component({
  selector: 'app-restaurant-info-page',
  templateUrl: './restaurant-info-page.component.html',
  styleUrls: ['./restaurant-info-page.component.scss']
})
export class RestaurantInfoPageComponent implements OnInit {

  map!: Map;

  restaurant: RestaurantDTO = {
    Id: 0,
    RestaurantName: '',
    ContactPhone: '',
    Image: 'assets/restaurant.png',
    ImagePath: 'assets/restaurant.png',
    Country: '',
    City: '',
    Street: '',
    StreetNumber: '',
    Ptt: 0,
    DisplayName: '',
    Longitude: 0,
    Latitude: 0,
    Changed: false,
    Delivery: 0
  }
  
  newOrder: OrderDTO = {
    Id: 0,
    IdRestaurant: 0,
    IdAppUser: 0,
    IdEmployee: 0,
    IdDeliverer: 0,
    OrderStatus: '',
    TotalPrice: 0,
    Tip: 0,
    Note: '',
    DateTime: '',
    Country: '',
    City: '',
    Street: '',
    StreetNumber: '',
    Ptt: 0,
    DisplayName: '',
    Longitude: 0,
    Latitude: 0,
    OrderItemsDTO: [],
    Reviewed: false,
    RestaurantName: ''
  }
  
  idUser: number;
  

  public restaurantIdFromRoute: number
  @ViewChild(ArticlesPageComponent) articlesPageComponent: ArticlesPageComponent | undefined;

  constructor(private route: ActivatedRoute,
    private http: HttpClient,
    private restaurantsUtilsService: RestaurantsUtilsService,
    private snackBarService: SnackBarService,
    private authService: AuthService,
    private dialog: MatDialog,
    private ordersService: OrdersService,
    private router: Router) {
    this.restaurantIdFromRoute = 0
    this.idUser = 0
  }

  ngOnInit(): void {
    var info = this.authService.getInfo();
    this.idUser = info.Id;

    const routeParams = this.route.snapshot.paramMap;
    this.restaurantIdFromRoute = Number(routeParams.get('restaurantId'));

    var iconFeature: Feature<Geometry>
    iconFeature = new Feature();
    
    var iconFeature2: Feature<Geometry>
    iconFeature2 = new Feature();

    this.restaurantsUtilsService.findRestaurantById(this.restaurantIdFromRoute)
    .subscribe((response) => {
      var temp = response.body as RestaurantDTOMessage;
      this.restaurant = temp.RestaurantDTO;
      this.snackBarService.openSnackBar(temp.Message);

      // dodati odmah dostavu na ukupnu cijenu
      this.newOrder.TotalPrice = this.restaurant.Delivery;

      if (this.articlesPageComponent) {
        this.articlesPageComponent.setRestaurantName(this.restaurant.RestaurantName);
      }
  
      iconFeature = new Feature({
        geometry: new Point(fromLonLat([this.restaurant.Longitude, this.restaurant.Latitude])),
      });
      this.initMap(this.restaurant.Longitude, this.restaurant.Latitude, iconFeature, iconFeature2)
    })
  }

  initMap(longitude: number, latitude: number, iconFeature: Feature<Geometry>,
    iconFeature2: Feature<Geometry>) {
    iconFeature.setStyle(
      new Style({
        image: new Icon({
        color: 'rgba(255, 0, 0, .5)',
        crossOrigin: 'anonymous',
        src: 'assets/bigdot.png',
        scale: 0.15,
        }),
      })
    )

    iconFeature2.setStyle(
      new Style({
        image: new Icon({
        color: 'rgba(255, 0, 0, .5)',
        crossOrigin: 'anonymous',
        src: 'assets/mark.png',
        scale: 0.15,
        }),
      })
    )

    this.map = new Map({
      view: new View({
        center: fromLonLat([longitude, latitude]),
        zoom: 15,
      }),
      layers: [
        new TileLayer({
          source: new OSM(),
        }),
        new Vector({
					source: new VectorTemp({
					features: [iconFeature, iconFeature2]
					}),
				})
      ],
      target: 'ol-map'
    });

    this.map.on('singleclick', (evt) => {
      var coordinate = transform(evt.coordinate, 'EPSG:3857', 'EPSG:4326');
      this.newOrder.Longitude = coordinate[0];
      this.newOrder.Latitude = coordinate[1];
      this.http.get(`https://nominatim.openstreetmap.org/reverse?format=json&lon=${coordinate[0]}&lat=${coordinate[1]}`)
      .subscribe((response) => {
          console.log(response)
          var temp = response as MapAddress
          
          this.newOrder.City = temp.address.city;
          this.newOrder.Country = temp.address.country;
          this.newOrder.Street = temp.address.road;
          this.newOrder.StreetNumber = temp.address.house_number;
          this.newOrder.Ptt = Number(temp.address.postcode);
          this.newOrder.DisplayName = temp.display_name;

          iconFeature2.set('geometry', new Point(fromLonLat([coordinate[0], coordinate[1]])));
      })
    });
  }

  addItemToOrder(item: OrderItemDTO) {
    if (item.ArticleName) {
      if (this.newOrder.OrderItemsDTO.length === 0) { 
        this.newOrder.IdAppUser = this.idUser
        this.newOrder.IdRestaurant = this.restaurant.Id
        this.newOrder.OrderItemsDTO.push(item)
      } else {
        var indikator = true
        this.newOrder.OrderItemsDTO.forEach(function (value) {
          if (value.ArticleName === item.ArticleName) {
            value.Quantity = value.Quantity + item.Quantity
            value.TotalPrice = value.TotalPrice + item.TotalPrice
            indikator = false
          }
        }); 
  
        if (indikator)
          this.newOrder.OrderItemsDTO.push(item)
      }
      
      this.newOrder.TotalPrice = this.newOrder.TotalPrice + item.TotalPrice
    }
  }

  removeItemFromOrder(item: OrderItemDTO) {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Removing order item",
        body: "You want to remove " + item.ArticleName + "?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        const index =   this.newOrder.OrderItemsDTO.indexOf(item, 0);

        if (index > -1) {
          this.newOrder.OrderItemsDTO.splice(index, 1);
          this.newOrder.TotalPrice = this.newOrder.TotalPrice - item.TotalPrice
        }  
      }
    })
  }

  createOrder() {
    if (this.newOrder.Tip >= 0 && this.newOrder.Country && 
      this.newOrder.City && this.newOrder.Street && this.newOrder.StreetNumber &&
      this.newOrder.Ptt && this.newOrder.DisplayName) {

        this.dialog.open(ConformationDialogComponent, {
          data:
          {
            title: "Creating Order",
            body: "You want to order?"
          },
        }).afterClosed().subscribe(result => {
          if (result) {
            this.newOrder.RestaurantName = this.restaurant.RestaurantName
            this.newOrder.IdRestaurant = this.restaurant.Id
            this.newOrder.OrderStatus = 'ORDERED'
            this.newOrder.IdAppUser = this.idUser
            this.newOrder.DateTime = moment().format('DD.MM.YYYY. HH:mm')
            this.ordersService.createOrder(this.newOrder)
            .subscribe((response) => {
              var temp = response.body as OrderDTOMessage;
              this.snackBarService.openSnackBar(temp.Message)
              if (temp.Message === 'order successfully created') {
                this.resetOrder();
                this.router.navigate(["/app/main/appuser/orders"]);
              }
            })
          }
        })
    }
  }

  resetOrder() {
    this.newOrder = {
      Id: 0,
      IdRestaurant: 0,
      IdAppUser: 0,
      IdEmployee: 0,
      IdDeliverer: 0,
      OrderStatus: '',
      TotalPrice: 0,
      Tip: 0,
      Note: '',
      DateTime: '',
      Country: '',
      City: '',
      Street: '',
      StreetNumber: '',
      Ptt: 0,
      DisplayName: '',
      Longitude: 0,
      Latitude: 0,
      OrderItemsDTO: [],
      Reviewed: false,
      RestaurantName: ''
    }
    
  }
}
