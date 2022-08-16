import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/role-guard/role.guard";
import { ArticlesPageComponent } from "../shared/pages/articles-page/articles-page.component";
import { RestaurantsPageComponent } from "../shared/pages/restaurants-page/restaurants-page.component";
import { CreateArticlePageComponent } from "./pages/create-article-page/create-article-page.component";
import { CreateRestaurantPageComponent } from "./pages/create-restaurant-page/create-restaurant-page.component";
import { CreateUserPageComponent } from "./pages/create-user-page/create-user-page.component";
import { ReportsPageComponent } from "./pages/reports-page/reports-page.component";
import { ReviewsPageComponent } from "./pages/reviews-page/reviews-page.component";
import { UsersPageComponent } from "./pages/users-page/users-page.component";

export const AdminRoutes: Routes = [
    {
        path: "users",
        pathMatch: "full",
        component: UsersPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    },
    {
        path: "user-info/:userId",
        pathMatch: "full",
        component: CreateUserPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    },
    {
        path: "restaurants",
        pathMatch: "full",
        component: RestaurantsPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    },
    {
        path: "restaurant-info/:restaurantId",
        pathMatch: "full",
        component: CreateRestaurantPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    },
    {
        path: "articles",
        pathMatch: "full",
        component: ArticlesPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    },
    {
        path: "article-info/:articleId",
        pathMatch: "full",
        component: CreateArticlePageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    },
    {
        path: "reviews",
        pathMatch: "full",
        component: ReviewsPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    },
    {
        path: "reports",
        pathMatch: "full",
        component: ReportsPageComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN" },
    }
];