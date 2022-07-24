import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';

import {MatSnackBarModule} from '@angular/material/snack-bar';
import { SnackBarService } from "./services/snack-bar.service";
import { MaterialExampleModule } from 'src/material.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { PaginationComponent } from './components/pagination/pagination.component';
import { UtilsService } from "./services/utils.service";
import { ConformationDialogComponent } from './components/conformation-dialog/conformation-dialog.component';
import { RestaurantsPageComponent } from './pages/restaurants-page/restaurants-page.component';
import { RestaurantCardComponent } from './components/restaurant-card/restaurant-card.component';
import { RestaurantsUtilsService } from "./services/restaurants-utils";
import { UsersUtilsService } from "./services/users-utils.service";
import { RestaurantInfoComponent } from './pages/restaurant-info/restaurant-info.component';
import { UserInfoComponent } from './pages/user-info/user-info.component';

@NgModule({
    declarations: [

    
    PaginationComponent,
            ConformationDialogComponent,
            RestaurantsPageComponent,
            RestaurantCardComponent,
            RestaurantInfoComponent,
            UserInfoComponent
  ],
    imports: [
        CommonModule,
        MaterialExampleModule,
        MatSnackBarModule,
        FormsModule,
        ReactiveFormsModule
    ],
    exports: [
        PaginationComponent
    ],
    providers: [
        SnackBarService,
        UtilsService,
        RestaurantsUtilsService,
        UsersUtilsService
    ]
})
export class SharedModule { }