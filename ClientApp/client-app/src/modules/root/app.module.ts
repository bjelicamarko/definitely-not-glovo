import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { AuthModule } from "../auth/auth.module";
import { NotFoundPageComponent } from './pages/not-found-page/not-found-page.component';
import { RootLayoutPageComponent } from './pages/root-layout-page/root-layout-page.component';
import { MainPageComponent } from './pages/main-page/main-page.component';
import { RegistrationPageComponent } from './pages/registration-page/registration-page.component';

import { ReactiveFormsModule } from "@angular/forms";
import { SharedModule } from "../shared/shared.module";
import { HeaderCommonComponent } from './components/header-common/header-common.component';
import { HeaderAdminComponent } from './components/header-admin/header-admin.component';
import { HeaderAppuserComponent } from './components/header-appuser/header-appuser.component';
import { HeaderEmployeeComponent } from './components/header-employee/header-employee.component';
import { HeaderDelivererComponent } from './components/header-deliverer/header-deliverer.component';

@NgModule({
    declarations: [
        AppComponent,
        NotFoundPageComponent,
        RootLayoutPageComponent,
        MainPageComponent,
        RegistrationPageComponent,
        HeaderCommonComponent,
        HeaderAdminComponent,
        HeaderAppuserComponent,
        HeaderEmployeeComponent,
        HeaderDelivererComponent,
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        HttpClientModule,
        ReactiveFormsModule,
        AuthModule,
        SharedModule,
    ],
    providers: [],
    bootstrap: [AppComponent]
  })
  export class AppModule { }