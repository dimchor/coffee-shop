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
  hasAdmin = false
  account: any
  token: any
  url = ""
  //orders = [{ ID: "128734", date: "22-2-2024", cost: "45.75", curStatus: "Completed" }]
  orders = [{ ID: "", date: "", cost: "", curStatus: "" }]
  //cart = [{ name: "", ID: "", cost: "" }]
  cart = [new ProductComponent('Java gia olous', '22.80', "Blends", "1001", "../assets/java-logo.jpg")]
  constructor(private http: HttpClient, public CartService: CartService, private router: Router) {

    this.account = this.getAccount()

    this.removeProduct(0)
    this.removeOrder(0)
    this.cart = this.CartService.getList()
  }

  async getAccount() {
    this.token = localStorage.getItem("session")
    this.token = this.token.replace(/(["])/g, "")
    this.url = this.url.concat("http://localhost:8080/v1/get/has_admin_rights/", this.token)
    const json = await fetch(this.url, {
      method: 'GET'
    }).then((response) => response.json())

    if (json.message == true) {
      this.hasAdmin = true
    }
    else if (json.message == false) {
      this.hasAdmin = false
    }

    //get account details

  }

  changePass() {
    alert("No");
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
    if (this.cart.length > 0) {
      let dateTime = new Date()
      var total = 0.0
      for (var p of this.cart) {
        total = total + parseFloat(p.pPrice)
      }
      this.orders.push({ ID: Math.floor(Math.random() * 99999).toString(), date: dateTime.toDateString(), cost: total.toString(), curStatus: "Pending" })
      this.cart = []
      alert("Order Submitted Successfully")
    }
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
