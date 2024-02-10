import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  username = ""
  password = ""

  user = {
    username: "",
    password: ""
  }
  constructor(private router: Router) { }
  session: any;

  login(user: string, pass: string) {

    alert(user + pass)
    //CALL API
    //this.session =
    if (true) {
      localStorage.setItem('session', JSON.stringify(this.session))
      this.router.navigate(['/account']);
    }
    return this.user;
  }


}
