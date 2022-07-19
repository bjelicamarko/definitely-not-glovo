import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { MainPageComponent } from "./pages/main-page/main-page.component";
import { NotFoundPageComponent } from "./pages/not-found-page/not-found-page.component";
import { RootLayoutPageComponent } from "./pages/root-layout-page/root-layout-page.component";

const routes: Routes = [
  {
    path: "not-glovo",
    component: RootLayoutPageComponent,
    children: [
      {
        path: "auth",
        loadChildren: () =>
          import("./../auth/auth.module").then((m) => m.AuthModule),
      },
      {
        path: "main",
        component: MainPageComponent
      }
    ]
  },
  {
    path: "",
    redirectTo: "not-glovo/auth/login",
    pathMatch: "full",
  },
  {
    path: "**",
    component: NotFoundPageComponent
  },
];
  
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }