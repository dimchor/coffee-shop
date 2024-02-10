import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

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

  user = {
    username: "",
    password: ""
  }
  constructor(private http: HttpClient) { };
  session: any;

  register(user: string, pass: string) {

    alert(user + pass)
    //CALL API
    //this.session =

    this.http.get('http://backend:8080/ping').
      subscribe((data) => {
        console.log(data)
        alert(data)
      });

  }



}
