import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Feature, View } from 'ol';
import { Geometry, Point } from 'ol/geom';
import TileLayer from 'ol/layer/Tile';
import Map from 'ol/Map';
import { fromLonLat } from 'ol/proj';
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
    Changed: false
  }
  
  public restaurantIdFromRoute: number
  @ViewChild(ArticlesPageComponent) articlesPageComponent: ArticlesPageComponent | undefined;

  constructor(private route: ActivatedRoute,
    private restaurantsUtilsService: RestaurantsUtilsService,
    private snackBarService: SnackBarService) {
    this.restaurantIdFromRoute = 0
  }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.restaurantIdFromRoute = Number(routeParams.get('restaurantId'));

    var iconFeature: Feature<Geometry>
    iconFeature = new Feature();
    
    this.restaurantsUtilsService.findRestaurantById(this.restaurantIdFromRoute)
    .subscribe((response) => {
      var temp = response.body as RestaurantDTOMessage;
      this.restaurant = temp.RestaurantDTO;
      this.snackBarService.openSnackBar(temp.Message);

      if (this.articlesPageComponent) {
        this.articlesPageComponent.setRestaurantName(this.restaurant.RestaurantName);
      }
  
      iconFeature = new Feature({
        geometry: new Point(fromLonLat([this.restaurant.Longitude, this.restaurant.Latitude])),
      });
      this.initMap(this.restaurant.Longitude, this.restaurant.Latitude, iconFeature)
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
}
