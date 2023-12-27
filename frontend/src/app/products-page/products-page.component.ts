import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductGridComponent } from '../product-grid/product-grid.component';
import { ProductFiltersComponent } from '../product-filters/product-filters.component';
import { NavbarComponent } from '../navbar/navbar.component';
import { ProductComponent } from '../product/product.component';

@Component({
  selector: 'app-products-page',
  standalone: true,
  imports: [CommonModule, NavbarComponent, ProductGridComponent, ProductFiltersComponent],
  templateUrl: './products-page.component.html',
  styleUrl: './products-page.component.css'
})
export class ProductsPageComponent {

  receiveMSG($event: ProductComponent[]) {
    this.filteredList = $event
  }

  productList = [
    new ProductComponent('Java gia olous', '22.80', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia ligous', '47.50', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia kanenan', '89.20', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('See ', '0.99', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See sharp', '14.70', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See plus plus', '185.00', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('Bython', '89.20', "Utilities", "../assets/java-logo.jpg"),


    new ProductComponent('Java gia olous', '22.80', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia ligous', '47.50', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia kanenan', '89.20', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('See ', '0.99', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See sharp', '14.70', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See plus plus', '185.00', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('Bython', '89.20', "Utilities", "../assets/java-logo.jpg"), new ProductComponent('Java gia olous', '22.80', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia ligous', '47.50', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia kanenan', '89.20', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('See ', '0.99', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See sharp', '14.70', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See plus plus', '185.00', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('Bython', '89.20', "Utilities", "../assets/java-logo.jpg"), new ProductComponent('Java gia olous', '22.80', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia ligous', '47.50', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia kanenan', '89.20', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('See ', '0.99', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See sharp', '14.70', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See plus plus', '185.00', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('Bython', '89.20', "Utilities", "../assets/java-logo.jpg"), new ProductComponent('Java gia olous', '22.80', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia ligous', '47.50', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia kanenan', '89.20', "Blends", "../assets/java-logo.jpg"),
    new ProductComponent('See ', '0.99', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See sharp', '14.70', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('See plus plus', '185.00', "Tea", "../assets/C_Logo.png"),
    new ProductComponent('Bython', '89.20', "Utilities", "../assets/java-logo.jpg"),

  ]
  tags = ["Blends", "Tea", "Utilities"]

  filteredList = this.productList
}
