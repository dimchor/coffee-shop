import { Routes } from '@angular/router';
//import { usrGuard } from './usr.guard';
import { HomepageComponent } from './homepage/homepage.component';
import { ProductsPageComponent } from './products-page/products-page.component';
import { AuthenticationPageComponent } from './authentication-page/authentication-page.component';
import { AuthGuardService } from './auth-guard.service';
import { AccountPageComponent } from './account-page/account-page.component';
import { AdminPageComponent } from './admin-page/admin-page.component';

export const routes: Routes = [
    { path: '', redirectTo: '/homepage', pathMatch: 'full' },
    { path: "products", component: ProductsPageComponent },
    { path: "homepage", component: HomepageComponent },
    { path: "login", component: AuthenticationPageComponent },
    { path: "admin", component: AdminPageComponent },
    { path: "account", component: AccountPageComponent, canActivate: [AuthGuardService] },

];
