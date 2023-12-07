import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductComponent } from '../product/product.component';

@Component({
  selector: 'app-product-grid',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './product-grid.component.html',
  styleUrl: './product-grid.component.css'
})
export class ProductGridComponent {
  productList = [
    new ProductComponent('Java gia olous', '22.80', "../assets/java-logo.jpg"),
    new ProductComponent('Java gia ligous', '47.50', "../assets/java-logo.jpg"),
    new ProductComponent('Java gia kanenan', '89.20', "../assets/java-logo.jpg")
  ]

}
