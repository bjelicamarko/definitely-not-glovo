import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { AppUserRoutes } from "./appuser.routes";

import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { MaterialExampleModule } from 'src/material.module';
import { SharedModule } from "../shared/shared.module";
import { RestaurantInfoPageComponent } from './pages/restaurant-info-page/restaurant-info-page.component';

@NgModule({
    declarations: [

    
    RestaurantInfoPageComponent
  ],
    imports: [
        CommonModule,
        RouterModule.forChild(AppUserRoutes),
        MaterialExampleModule,
        FormsModule,
        ReactiveFormsModule,
        SharedModule
    ]
})
export class AppUserModule {}