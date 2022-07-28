import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { MaterialExampleModule } from 'src/material.module';
import { SharedModule } from "../shared/shared.module";
import { DelivererRoutes } from "./deliverer.routes";

@NgModule({
    declarations: [

    ],
    imports: [
        CommonModule,
        RouterModule.forChild(DelivererRoutes),
        MaterialExampleModule,
        FormsModule,
        ReactiveFormsModule,
        SharedModule
    ]
})
export class DelivererModule {}