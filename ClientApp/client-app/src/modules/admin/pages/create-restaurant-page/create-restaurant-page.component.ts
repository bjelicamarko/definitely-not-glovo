import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Feature, View } from 'ol';
import TileLayer from 'ol/layer/Tile';
import Map from 'ol/Map';
import { fromLonLat, transform } from 'ol/proj';
import OSM from 'ol/source/OSM';
import Style from 'ol/style/Style';
import Icon from 'ol/style/Icon';
import { MapAddress } from 'src/modules/shared/models/MapAddress';
import { RestaurantDTO } from 'src/modules/shared/models/RestaurantDTO';
import { Point } from 'ol/geom';
import Vector from 'ol/layer/Vector';
import VectorTemp from 'ol/source/Vector';

@Component({
  selector: 'app-create-restaurant-page',
  templateUrl: './create-restaurant-page.component.html',
  styleUrls: ['./create-restaurant-page.component.scss']
})
export class CreateRestaurantPageComponent implements OnInit {

  map!: Map;

  restaurant: RestaurantDTO = {
    Id: 0,
    RestaurantName: '',
    ContactPhone: '',
    City: '',
    Street: '',
    StreetNumber: '',
    Ptt: 0,
    Longitude: 0,
    Latitude: 0,
    Image: '',
    ImagePath: '',
    Country: '',
    DisplayName: ''
  }
  
  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    var iconFeature = new Feature();

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
        center: fromLonLat([18.3501358, 42.7060377]),
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

    this.map.on('singleclick', (evt) => {
      var coordinate = transform(evt.coordinate, 'EPSG:3857', 'EPSG:4326');
      this.restaurant.Longitude = coordinate[0];
      this.restaurant.Latitude = coordinate[1];
      this.http.get(`http://nominatim.openstreetmap.org/reverse?format=json&lon=${coordinate[0]}&lat=${coordinate[1]}`)
      .subscribe((response) => {
          console.log(response)
          var temp = response as MapAddress
          
          this.restaurant.City = temp.address.town;
          this.restaurant.Country = temp.address.country;
          this.restaurant.Street = temp.address.road;
          this.restaurant.StreetNumber = temp.address.house_number;
          this.restaurant.Ptt = Number(temp.address.postcode);
          this.restaurant.DisplayName = temp.display_name;

          iconFeature.set('geometry', new Point(fromLonLat([coordinate[0], coordinate[1]])));
      })
  });
  }

}
