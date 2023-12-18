import { Routes } from '@angular/router';
import { TestPageComponent } from './test-page/test-page.component';
import { HomepageComponent } from './homepage/homepage.component';
import { ProductsPageComponent } from './products-page/products-page.component';

export const routes: Routes = [
    { path: '', redirectTo: '/homepage', pathMatch: 'full' },
    { path: "products", component: ProductsPageComponent },
    { path: "homepage", component: HomepageComponent },

];
