import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

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
  post_str = ""
  //response: any
  constructor(private router: Router, private http: HttpClient) { }
  session: any;

  login(user: string, pass: string) {

    this.post_str = JSON.stringify({ "username": user, "password": pass })
    //this.response = this.http.post("http://localhost:8080/v1/post/login_user", this.post_str)
    this.http.post("http://localhost:8080/v1/post/login_user", this.post_str).
      subscribe((response) => {
        console.log(response);
      });

    //console.log(this.response)

    if (true) {//(response.username === user && response.pass === pass) {
      localStorage.setItem('session', JSON.stringify(this.session))
      this.router.navigate(['/account']);
    }
    return this.user;
  }


}
