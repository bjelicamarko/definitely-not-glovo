import { Routes } from "@angular/router";
import { ProfileInfoPageComponent } from "../shared/pages/profile-info-page/profile-info-page.component";
import { RestaurantsPageComponent } from "../shared/pages/restaurants-page/restaurants-page.component";
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
        path: "restaurant-info/:restaurantId",
        pathMatch: "full",
        component: CreateRestaurantPageComponent
    },
    {
        path: "restaurants",
        pathMatch: "full",
        component: RestaurantsPageComponent
    },
];