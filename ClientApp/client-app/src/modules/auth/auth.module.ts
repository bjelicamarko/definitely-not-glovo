import { CommonModule } from "@angular/common";
import { NgModule } from "@angular/core";
import { ReactiveFormsModule } from "@angular/forms";
import { RouterModule } from "@angular/router";
import { AuthRoutes } from "./auth.routes";
import { LoginComponent } from "./pages/login/login.component";

@NgModule({
    declarations: [
      LoginComponent
    ],
    imports: [
      CommonModule,
      ReactiveFormsModule,
      RouterModule.forChild(AuthRoutes),
    ],
    providers: []
  })
  export class AuthModule { }