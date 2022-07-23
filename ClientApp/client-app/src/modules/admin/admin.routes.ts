import { Routes } from "@angular/router";
import { ProfileInfoPageComponent } from "../shared/pages/profile-info-page/profile-info-page.component";
import { CreateRestaurantPageComponent } from "./pages/create-restaurant-page/create-restaurant-page.component";
import { UsersPageComponent } from "./pages/users-page/users-page.component";

export const AdminRoutes: Routes = [
    {
        path: "users",
        pathMatch: "full",
        component: UsersPageComponent,
    },
    {
        path: "profile-info/:userId",
        pathMatch: "full",
        component: ProfileInfoPageComponent
    },
    {
        path: "createUser",
        pathMatch: "full",
        component: ProfileInfoPageComponent
    },
    {
        path: "createRestaurant",
        pathMatch: "full",
        component: CreateRestaurantPageComponent
    }
];