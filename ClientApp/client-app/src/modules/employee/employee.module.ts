import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { MaterialExampleModule } from 'src/material.module';
import { SharedModule } from "../shared/shared.module";
import { EmployeeRoutes } from "./employee.routes";
import { ReviewsPageComponent } from './pages/reviews-page/reviews-page.component';

@NgModule({
    declarations: [
    
    ReviewsPageComponent
  ],
    imports: [
        CommonModule,
        RouterModule.forChild(EmployeeRoutes),
        MaterialExampleModule,
        FormsModule,
        ReactiveFormsModule,
        SharedModule
    ]
})
export class EmployeeModule {}