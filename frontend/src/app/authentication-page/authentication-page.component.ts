import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../navbar/navbar.component';
import { LoginComponent } from '../login/login.component';
import { RegisterComponent } from '../register/register.component';

@Component({
  selector: 'app-authentication-page',
  standalone: true,
  imports: [CommonModule, NavbarComponent, LoginComponent, RegisterComponent],
  templateUrl: './authentication-page.component.html',
  styleUrl: './authentication-page.component.css'
})
export class AuthenticationPageComponent {
  state = true;
  btn = "Register"
  str = "New User?"
  changeState() {

    if (this.state) {
      this.btn = "Login"
      this.str = "Existing User?"
    }
    else {
      this.btn = "Register"
      this.str = "New User?"
    }
    this.state = !this.state;
  }

}
