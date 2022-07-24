import { Routes } from "@angular/router";
import { RestaurantsPageComponent } from "../shared/pages/restaurants-page/restaurants-page.component";
import { CreateRestaurantPageComponent } from "./pages/create-restaurant-page/create-restaurant-page.component";
import { CreateUserPageComponent } from "./pages/create-user-page/create-user-page.component";
import { UsersPageComponent } from "./pages/users-page/users-page.component";

export const AdminRoutes: Routes = [
    {
        path: "users",
        pathMatch: "full",
        component: UsersPageComponent,
    },
    {
        path: "user-info/:userId",
        pathMatch: "full",
        component: CreateUserPageComponent,
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