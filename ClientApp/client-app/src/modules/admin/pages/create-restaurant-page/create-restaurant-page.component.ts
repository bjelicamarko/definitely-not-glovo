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
import { Geometry, Point } from 'ol/geom';
import Vector from 'ol/layer/Vector';
import VectorTemp from 'ol/source/Vector';
import { RestaurantsService } from '../../services/restaurants.service';
import { RestaurantDTOMessage } from 'src/modules/shared/models/RestaurantDTOMessage';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { ActivatedRoute, Router } from '@angular/router';
import { RestaurantsUtilsService } from 'src/modules/shared/services/restaurants-utils';

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
    Image: 'assets/restaurant.png',
    ImagePath: 'assets/restaurant.png',
    Country: '',
    DisplayName: '',
    Changed: false,
    Delivery: 0
  }
  
  public selectedFile: File | undefined

  public restaurantIdFromRoute: number

  constructor(private http: HttpClient,
    private restaurantsService: RestaurantsService,
    private restaurantsUtilsService: RestaurantsUtilsService,
    private snackBarService: SnackBarService,
    private router: Router,
    private route: ActivatedRoute) { 
      this.restaurantIdFromRoute = 0
    }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.restaurantIdFromRoute = Number(routeParams.get('restaurantId'));
  
    var iconFeature: Feature<Geometry>
    iconFeature = new Feature();

    var longitude = 18.3501358
    var latitude = 42.7060377

    if (this.restaurantIdFromRoute !== 0) {
      this.restaurantsUtilsService.findRestaurantById(this.restaurantIdFromRoute)
      .subscribe((response) => {
        var temp = response.body as RestaurantDTOMessage;
        this.restaurant = temp.RestaurantDTO;
        this.snackBarService.openSnackBar(temp.Message)

        iconFeature = new Feature({
          geometry: new Point(fromLonLat([this.restaurant.Longitude, this.restaurant.Latitude])),
        });
        this.initMap(this.restaurant.Longitude, this.restaurant.Latitude, iconFeature)
      })
    } else {
      this.initMap(longitude, latitude, iconFeature)
    }
    
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

    this.map.on('singleclick', (evt) => {
      var coordinate = transform(evt.coordinate, 'EPSG:3857', 'EPSG:4326');
      this.restaurant.Longitude = coordinate[0];
      this.restaurant.Latitude = coordinate[1];
      this.http.get(`https://nominatim.openstreetmap.org/reverse?format=json&lon=${coordinate[0]}&lat=${coordinate[1]}`)
      .subscribe((response) => {
          console.log(response)
          var temp = response as MapAddress
          
          this.restaurant.City = temp.address.city;
          this.restaurant.Country = temp.address.country;
          this.restaurant.Street = temp.address.road;
          this.restaurant.StreetNumber = temp.address.house_number;
          this.restaurant.Ptt = Number(temp.address.postcode);
          this.restaurant.DisplayName = temp.display_name;

          iconFeature.set('geometry', new Point(fromLonLat([coordinate[0], coordinate[1]])));
      })
    });
  }


  createRestaurant() {
    if (this.restaurant.RestaurantName && this.restaurant.ContactPhone && 
      this.restaurant.City && this.restaurant.Street && this.restaurant.StreetNumber &&
      this.restaurant.Image && this.restaurant.Country && 
      this.restaurant.ImagePath !== 'assets/restaurant.png' && this.restaurant.Delivery >= 0) {
        let reader = new FileReader();
        reader.readAsDataURL(this.selectedFile!);
        reader.onload = () => {
          this.restaurant.Image = reader.result;
          this.restaurantsService.createRestaurant(this.restaurant)
          .subscribe((response) => {
            var temp = response.body as RestaurantDTOMessage;
            this.snackBarService.openSnackBar(temp.Message);
            this.restaurant = temp.RestaurantDTO;
            this.router.navigate(["/app/main/admin/restaurants"]);
          })
        };
        reader.onerror = function (error) {
         console.log('Error: ', error);
        };
    }
  }

  updateRestaurant() {
    if (this.restaurant.RestaurantName && this.restaurant.ContactPhone && 
      this.restaurant.City && this.restaurant.Street && this.restaurant.StreetNumber &&
      this.restaurant.Image && this.restaurant.Country && 
      this.restaurant.ImagePath !== 'assets/restaurant.png' && this.restaurant.Delivery >= 0) {
      
      if (this.restaurant.Changed) {
        let reader = new FileReader();
        reader.readAsDataURL(this.selectedFile!);
        reader.onload = () => {
          this.restaurant.Image = reader.result;
          this.restaurantsService.updateRestaurant(this.restaurant)
          .subscribe((response) => {
            var temp = response.body as RestaurantDTOMessage;
            this.snackBarService.openSnackBar(temp.Message);
            this.restaurant = temp.RestaurantDTO;
            this.router.navigate(["/app/main/admin/restaurants"]);
          })
        };
        reader.onerror = function (error) {
         console.log('Error: ', error);
        };
      } else {
        this.restaurantsService.updateRestaurant(this.restaurant)
        .subscribe((response) => {
          var temp = response.body as RestaurantDTOMessage;
          this.snackBarService.openSnackBar(temp.Message);
          this.restaurant = temp.RestaurantDTO;
          this.router.navigate(["/app/main/admin/restaurants"]);
        })
      }
    }
  }

  onFileChanged(event: any) {
    this.selectedFile = event.target.files[0]
    this.restaurant.ImagePath = "images/" + this.selectedFile?.name as string;
    this.restaurant.Changed = true
  }
  
}
