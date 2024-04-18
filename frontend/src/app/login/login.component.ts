import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';


type userResponse = {
  username: string;
  password: string;
};

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
  status_code: any

  post_str = ""
  constructor(private router: Router, private http: HttpClient) { }
  session: any;

  async login(user: string, pass: string) {

    this.post_str = JSON.stringify({ "username": user, "password": pass })

    const json = await fetch("http://localhost:8080/v1/post/login_user", {
      method: 'POST',
      body: this.post_str
    }).then((response) => { this.status_code = response; return response.json() })
    // const status = await fetch("http://localhost:8080/v1/post/login_user", {
    //   method: 'POST',
    //   body: this.post_str
    // }).then((response) => response)
    // alert(pass)
    // alert(json)
    if (this.status_code.status === 200) {
      localStorage.setItem('session', JSON.stringify(json))
      this.router.navigate(['/account']);
    }
    else {
      alert("Incorrect Username or Password")
    }

    return null
  }


}
