import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {

  username = ""
  password = ""
  first_name = ""
  last_name = ""
  address = ""

  user = {
    username: "",
    password: "",
    first_name: "",
    last_name: "",
    address: ""
  }

  post_str = ""

  constructor(private router: Router, private http: HttpClient) { };
  session: any;

  async register(user: string, pass: string, fn: string, ln: string, addr: string) {

    alert(user + pass)

    this.post_str = JSON.stringify({
      "username": user,
      "password": pass,
      "first_name": fn,
      "last_name": ln,
      "address": addr
    })

    const json = await fetch("http://localhost:8080/v1/post/new_user", {
      method: 'POST',
      body: this.post_str
    }).then((response) => response.json())

    if (json != null) {
      localStorage.setItem('session', JSON.stringify(json.token))
      this.router.navigate(['/account']);
    }
  }



}
