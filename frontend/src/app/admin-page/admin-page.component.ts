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

  product_json: any
  product = {
    product_name: "",
    price: "",
    description: ""
  }

  constructor() { }

  submit(name: string, price: string, description: string) {
    this.product.product_name = name
    this.product.price = price
    this.product.description = description

    this.product_json = JSON.stringify(this.product)
    //POST


  }


}
