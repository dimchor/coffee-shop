import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../navbar/navbar.component';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-account-page',
  standalone: true,
  imports: [CommonModule, NavbarComponent],
  templateUrl: './account-page.component.html',
  styleUrl: './account-page.component.css'
})
export class AccountPageComponent {
  username = "placeholder"
  state = 2
  orders = [{ ID: "128734", date: "22-2-2024", cost: "45.75", curStatus: "Completed" },
  { ID: "128734", date: "22-2-2024", cost: "45.75", curStatus: "Completed" },
  { ID: "128734", date: "22-2-2024", cost: "45.75", curStatus: "Completed" },
  { ID: "128734", date: "22-2-2024", cost: "45.75", curStatus: "Completed" },
  { ID: "128734", date: "22-2-2024", cost: "45.75", curStatus: "Completed" }
  ]
  cart = [{ name: "Sea Plus PLus", ID: "128734", cost: "45.75", quantity: "1" }]

  constructor(private http: HttpClient) {
    //API Call for account info
    //
    // this.http.get('http://backend:8080/ping').
    //   subscribe((data) => {
    //     console.log(data);
    //   });
  }

  getCartItems() {

  }

  changePass() {

  }

  viewOrders() {
    this.state = 1
  }

  viewCart() {
    this.state = 2
  }

}
