import { Component, Inject, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { OrderDTO } from '../../models/OrderDTO';
import Map from 'ol/Map';
import { Feature, View } from 'ol';
import { Geometry, Point } from 'ol/geom';
import { fromLonLat } from 'ol/proj';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';
import Icon from 'ol/style/Icon';
import Style from 'ol/style/Style';
import Vector from 'ol/layer/Vector';
import VectorTemp from 'ol/source/Vector';
import { OrdersUtilsService } from '../../services/orders-utils';
import { OrderDTOMessage } from '../../models/OrderDTOMessage';
import { SnackBarService } from '../../services/snack-bar.service';
import { OrderStatusDTO } from '../../models/OrderStatusDTO';
import { ReviewDTO } from '../../models/ReviewDTO';
import { ReviewsUtilsService } from '../../services/reviews-utils';
import { ReviewDTOMessage } from '../../models/ReviewDTOMessage';
import { CreateReviewDialogComponent } from '../create-review-dialog/create-review-dialog.component';
import { AuthService } from 'src/modules/auth/services/auth.service';
import * as moment from 'moment';

export interface IdRole {
  Id: number;
  Role: string;
  IdUser: number;
}

@Component({
  selector: 'app-order-info-dialog',
  templateUrl: './order-info-dialog.component.html',
  styleUrls: ['./order-info-dialog.component.scss']
})
export class OrderInfoDialogComponent implements OnInit {

  map!: Map;

  order: OrderDTO = {
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
    Reviewed: false
  }

  orderStatusDTO: OrderStatusDTO = {
    IdOrder: 0,
    OrderStatus: '',
    IdEmployee: 0,
    IdDeliverer: 0
  }

  review: ReviewDTO = {
    Id: 0,
    Comment: '',
    Rating: 0,
    InappropriateContent: false,
    DateTime: '',
    IdRestaurant: 0,
    IdOrder: 0,
    IdUser: 0,
    EmailUser: ''
  }
  
  info: any = {}

  constructor(public dialogRef: MatDialogRef<OrderInfoDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: IdRole,
    private ordersUtilsService: OrdersUtilsService,
    private snackBarService: SnackBarService,
    private reviewsUtilsService: ReviewsUtilsService,
    private authService: AuthService,
    public dialog: MatDialog) { 
      dialogRef.beforeClosed().subscribe(() => dialogRef.close(this.order));
    }

  ngOnInit(): void {
    this.info = this.authService.getInfo();

    var iconFeature: Feature<Geometry>
    iconFeature = new Feature();

    this.ordersUtilsService.findOrderById(this.data.Id)
    .subscribe((response) => {
      var temp = response.body as OrderDTOMessage;
      this.snackBarService.openSnackBar(temp.Message);
      this.order = temp.OrderDTO;
      iconFeature = new Feature({
        geometry: new Point(fromLonLat([this.order.Longitude, this.order.Latitude])),
      });
      this.initMap(this.order.Longitude, this.order.Latitude, iconFeature)
      if (this.order.Reviewed) {
          this.reviewsUtilsService.findReviewByOrder(this.order.Id)
          .subscribe((response) => {
            var temp = response.body as ReviewDTOMessage;
            this.review = temp.ReviewDTO;
          })
      }
    })
  }

  initMap(longitude: number, latitude: number, iconFeature: Feature<Geometry>) {
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
					features: [iconFeature]
					}),
				})
      ],
      target: 'ol-map'
    });
  }
  
  changeStatus(status: string) {
    this.orderStatusDTO.IdOrder = this.order.Id;
    this.orderStatusDTO.OrderStatus = status;
    if (status === 'ACCEPTED' || status === 'READY')
      this.orderStatusDTO.IdEmployee = this.data.IdUser;
    if (status === 'TAKEN' || status === 'DELIVERED')
      this.orderStatusDTO.IdDeliverer = this.data.IdUser;

    this.ordersUtilsService.changeStatusOfOrder(this.orderStatusDTO)
    .subscribe((response) => {
      var temp = response.body as OrderDTOMessage;
      this.snackBarService.openSnackBar(temp.Message);
      this.order.OrderStatus = temp.OrderDTO.OrderStatus;
      if (temp.OrderDTO.IdDeliverer) 
        this.order.IdDeliverer = temp.OrderDTO.IdDeliverer;
      if (temp.OrderDTO.IdEmployee)
        this.order.IdEmployee = temp.OrderDTO.IdEmployee;
    })
  }

  reviewIt() {
    const ref = this.dialog.open(CreateReviewDialogComponent, {
      data: this.review,
      autoFocus: false,
      maxHeight: '90vh',
      width: '50%',
    });

    ref.afterClosed().subscribe(
      () => {
        if (this.review.Rating >= 0 && this.review.Rating <= 10 && this.review.Comment) {
          this.review.DateTime = moment().format('DD.MM.YYYY. HH:mm')
          this.review.IdRestaurant = this.order.IdRestaurant
          this.review.IdOrder = this.order.Id
          this.review.IdUser = this.order.IdAppUser
          this.review.EmailUser = this.info.email;
          console.log(this.review)
          this.reviewsUtilsService.createReview(this.review)
          .subscribe((response) => {
            var temp = response.body as ReviewDTOMessage;
            if (temp.Message === 'review successfully created') {
              this.ordersUtilsService.reviewOrder(this.order.Id)
              .subscribe((response) => {
                var temp = response.body as OrderDTOMessage;
                if (temp.Message === 'order successfully reviewed') {
                  this.order.Reviewed = true;
                }
              })
            }
          })
        }
      }  
    )
  }
}
