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
import { UserInfoComponent } from './pages/user-info/user-info.component';
import { ArticlesUtilsService } from "./services/articles-utils";
import { ArticleCardComponent } from './components/article-card/article-card.component';
import { ArticlesPageComponent } from './pages/articles-page/articles-page.component';
import { HTTP_INTERCEPTORS } from "@angular/common/http";
import { Interceptor } from "./interceptors/interceptor.interceptor";
import { OrdersPageComponent } from './pages/orders-page/orders-page.component';
import { OrderInfoDialogComponent } from './components/order-info-dialog/order-info-dialog.component';
import { OrderitemInitDialogComponent } from './components/orderitem-init-dialog/orderitem-init-dialog.component';

@NgModule({
    declarations: [

    
    PaginationComponent,
            ConformationDialogComponent,
            RestaurantsPageComponent,
            RestaurantCardComponent,
            UserInfoComponent,
            ArticleCardComponent,
            ArticlesPageComponent,
            OrdersPageComponent,
            OrderInfoDialogComponent,
            OrderitemInitDialogComponent
  ],
    imports: [
        CommonModule,
        MaterialExampleModule,
        MatSnackBarModule,
        FormsModule,
        ReactiveFormsModule
    ],
    exports: [
        PaginationComponent,
        ArticlesPageComponent
    ],
    providers: [
        SnackBarService,
        UtilsService,
        RestaurantsUtilsService,
        UsersUtilsService,
        ArticlesUtilsService,
        { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
    ]
})
export class SharedModule { }