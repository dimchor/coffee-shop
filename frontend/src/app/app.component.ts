import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { RouterModule } from '@angular/router';
import { HomepageComponent } from './homepage/homepage.component';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { JwtHelperService, JwtModule, JwtModuleOptions } from '@auth0/angular-jwt';
import { HttpClientModule } from '@angular/common/http';


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, FormsModule, RouterOutlet,
    HomepageComponent, RouterLink, RouterLinkActive,
    RouterModule, HttpClientModule, JwtModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'frontend';
}
