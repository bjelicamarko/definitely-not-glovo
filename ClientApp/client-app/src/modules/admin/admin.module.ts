import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { AdminRoutes } from './admin.routes';
import { FormsModule, ReactiveFormsModule } from "@angular/forms";

import { MaterialExampleModule } from 'src/material.module';
import { UsersPageComponent } from './pages/users-page/users-page.component';
import { SharedModule } from "../shared/shared.module";
import { UserCardComponent } from './components/user-card/user-card.component';
import { CreateUserPageComponent } from './pages/create-user-page/create-user-page.component';
import { CreateRestaurantPageComponent } from './pages/create-restaurant-page/create-restaurant-page.component';
import { CreateArticlePageComponent } from './pages/create-article-page/create-article-page.component';
import { ReviewsPageComponent } from './pages/reviews-page/reviews-page.component';
import { ReportsPageComponent } from './pages/reports-page/reports-page.component';
import { PieChartComponent } from './components/pie-chart/pie-chart.component';
import { DoughnutChartComponent } from './components/doughnut-chart/doughnut-chart.component';

import { NgChartsModule } from 'ng2-charts';

@NgModule({
   declarations: [
    UsersPageComponent,
    UserCardComponent,
    CreateUserPageComponent,
    CreateRestaurantPageComponent,
    CreateArticlePageComponent,
    ReviewsPageComponent,
    ReportsPageComponent,
    PieChartComponent,
    DoughnutChartComponent,
  ],
    imports: [
        CommonModule,
        RouterModule.forChild(AdminRoutes),
        MaterialExampleModule,
        FormsModule,
        ReactiveFormsModule,
        SharedModule,
        NgChartsModule
    ]
})
export class AdminModule { }