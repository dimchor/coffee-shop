import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../navbar/navbar.component';
import { HttpClient } from '@angular/common/http';
import { CartService } from '../cart.service';
import { ProductComponent } from '../product/product.component';
import { Router } from '@angular/router';

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
  //orders = [{ ID: "128734", date: "22-2-2024", cost: "45.75", curStatus: "Completed" }]
  orders = [{ ID: "", date: "", cost: "", curStatus: "" }]
  //cart = [{ name: "", ID: "", cost: "" }]
  cart = [new ProductComponent('Java gia olous', '22.80', "Blends", "1001", "../assets/java-logo.jpg")]
  constructor(private http: HttpClient, public CartService: CartService, private router: Router) {


    //API Call for account info
    //
    this.http.get('http://localhost:8080/ping').
      subscribe((data) => {
        console.log(data);
      });

    this.removeProduct(0)
    this.removeOrder(0)
    this.cart = this.CartService.getList()
  }

  async ping() {
    let response = await fetch('http://localhost:8080/ping', {
      method: 'GET'
    });
    let text = await response.text();
    return text;
  }

  changePass() {
    alert(this.ping());
  }


  removeProduct(index: number) {
    //this.cart = this.cart.filter(item => item !== this.cart[index]);
    this.cart.splice(index, 1);
  }

  removeOrder(index: number) {
    //this.orders = this.orders.filter(item => item !== this.orders[index]);
    this.orders.splice(index, 1);
  }

  checkout() {

  }

  viewOrders() {
    this.state = 1
  }

  viewCart() {
    this.state = 2
  }

  admin() {
    this.router.navigate(['/admin']);
  }
}
