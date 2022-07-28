import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/role-guard/role.guard";
import { OrdersPageComponent } from "../shared/pages/orders-page/orders-page.component";
import { RestaurantsPageComponent } from "../shared/pages/restaurants-page/restaurants-page.component";
import { RestaurantInfoPageComponent } from "./pages/restaurant-info-page/restaurant-info-page.component";

export const AppUserRoutes: Routes = [
    {
        path: "restaurants",
        pathMatch: "full",
        component: RestaurantsPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "APPUSER" },
    },
    {
        path: "restaurant-info/:restaurantId",
        pathMatch: "full",
        component: RestaurantInfoPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "APPUSER" },
    },
    {
        path: "orders",
        pathMatch: "full",
        component: OrdersPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "APPUSER" },
    }
];