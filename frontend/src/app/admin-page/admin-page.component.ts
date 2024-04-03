import { Component } from '@angular/core';
import { CommonModule, JsonPipe } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-admin-page',
  standalone: true,
  imports: [CommonModule, FormsModule],
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
    product_name: "",
    price: 0,
    description: "",
    token: localStorage.getItem('session')
  }
  constructor() { }

  async submit(name: string, price: string, description: string) {
    this.product.product_name = name
    this.product.price = parseFloat(price)
    this.product.description = description

    this.post_str = JSON.stringify(this.product)
    //POST
    const json = await fetch("http://localhost:8080/v1/post/product", {
      method: 'POST',
      body: this.post_str
    }).then((response) => response.json())

  }


}
