import { Routes } from '@angular/router';
//import { usrGuard } from './usr.guard';
import { HomepageComponent } from './homepage/homepage.component';
import { ProductsPageComponent } from './products-page/products-page.component';
import { AuthenticationPageComponent } from './authentication-page/authentication-page.component';
import { AuthGuardService } from './auth-guard.service';

export const routes: Routes = [
    { path: '', redirectTo: '/homepage', pathMatch: 'full' },
    { path: "products", component: ProductsPageComponent, canActivate: [AuthGuardService] },
    { path: "homepage", component: HomepageComponent },
    { path: "login", component: AuthenticationPageComponent }

];
