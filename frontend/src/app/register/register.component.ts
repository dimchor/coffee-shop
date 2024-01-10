import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

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

  session: any;

  register(user: string, pass: string) {

    alert(user + pass)
    //CALL API
    //this.session =

  }



}
