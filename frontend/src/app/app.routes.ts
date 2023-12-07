import { Routes } from '@angular/router';
import { TestPageComponent } from './test-page/test-page.component';
import { HomepageComponent } from './homepage/homepage.component';

export const routes: Routes = [
    { path: '', redirectTo: '/home', pathMatch: 'full' },
    { path: "test-page", component: TestPageComponent },
    { path: "homepage", component: HomepageComponent },

];
