import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

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

  session: any;

  login(user: string, pass: string) {

    alert(user + pass)
    //CALL API
    //this.user =
    //localStorage.setItem('session', JSON.stringify(this.session))

    return this.user;
  }


}
