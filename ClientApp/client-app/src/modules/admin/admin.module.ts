import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { AdminRoutes } from './admin.routes';
import { FormsModule, ReactiveFormsModule } from "@angular/forms";

import { MaterialExampleModule } from 'src/material.module';
import { UsersPageComponent } from './pages/users-page/users-page.component';
import { SharedModule } from "../shared/shared.module";

@NgModule({
   declarations: [
    UsersPageComponent
  ],
    imports: [
        CommonModule,
        RouterModule.forChild(AdminRoutes),
        MaterialExampleModule,
        FormsModule,
        ReactiveFormsModule,
        SharedModule
    ]
})
export class AdminModule { }