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

@NgModule({
    declarations: [
        AppComponent,
        NotFoundPageComponent,
        RootLayoutPageComponent,
        MainPageComponent,
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        HttpClientModule,
        AuthModule,
    ],
    providers: [],
    bootstrap: [AppComponent]
  })
  export class AppModule { }