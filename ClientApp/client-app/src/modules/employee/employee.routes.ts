import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/role-guard/role.guard";
import { OrdersPageComponent } from "../shared/pages/orders-page/orders-page.component";
import { ReviewsPageComponent } from "./pages/reviews-page/reviews-page.component";

export const EmployeeRoutes: Routes = [
    {
        path: "orders",
        pathMatch: "full",
        component: OrdersPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "EMPLOYEE" },
    },
    {
        path: "reviews",
        pathMatch: "full",
        component: ReviewsPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "EMPLOYEE" },
    }
];