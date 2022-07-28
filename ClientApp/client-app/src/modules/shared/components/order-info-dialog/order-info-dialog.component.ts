import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
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
    OrderItemsDTO: []
  }

  orderStatusDTO: OrderStatusDTO = {
    IdOrder: 0,
    OrderStatus: '',
    IdEmployee: 0,
    IdDeliverer: 0
  }

  constructor(public dialogRef: MatDialogRef<OrderInfoDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: IdRole,
    private ordersUtilsService: OrdersUtilsService,
    private snackBarService: SnackBarService) { 
      dialogRef.beforeClosed().subscribe(() => dialogRef.close(this.order));
    }

  ngOnInit(): void {

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
}
