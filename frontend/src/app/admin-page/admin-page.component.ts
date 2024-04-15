import { Component } from '@angular/core';
import { CommonModule, JsonPipe } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavbarComponent } from '../navbar/navbar.component';

@Component({
  selector: 'app-admin-page',
  standalone: true,
  imports: [CommonModule, FormsModule, NavbarComponent],
  templateUrl: './admin-page.component.html',
  styleUrl: './admin-page.component.css'
})
export class AdminPageComponent {

  post_str = ""
  tmp = {
    product_name: "",
    price: "",
    description: ""
    //token: localStorage.getItem('session')
  }
  product = {
    name: "",
    price: 0,
    description: "",
    token: localStorage.getItem('session')
  }
  constructor() { }

  async submit(product_name: string, price: string, description: string) {
    this.product.name = product_name
    this.product.price = parseFloat(price)
    this.product.description = description
    if (this.product.token) {
      this.product.token = this.product.token.replace(/(["])/g, "")
    }
    this.post_str = JSON.stringify(this.product)
    alert(this.post_str)
    //POST
    const json = await fetch("http://localhost:8080/v1/post/product", {
      method: 'POST',
      body: this.post_str
    }).then((response) => response.json())

  }


}
