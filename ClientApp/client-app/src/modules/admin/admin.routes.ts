import { Routes } from "@angular/router";
import { UsersPageComponent } from "./pages/users-page/users-page.component";

export const AdminRoutes: Routes = [
    {
        path: "users",
        pathMatch: "full",
        component: UsersPageComponent,
    }
];