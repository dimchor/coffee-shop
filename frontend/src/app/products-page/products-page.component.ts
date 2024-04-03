import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductGridComponent } from '../product-grid/product-grid.component';
import { ProductFiltersComponent } from '../product-filters/product-filters.component';
import { NavbarComponent } from '../navbar/navbar.component';
import { ProductComponent } from '../product/product.component';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-products-page',
  standalone: true,
  imports: [CommonModule, NavbarComponent, ProductGridComponent, ProductFiltersComponent],
  templateUrl: './products-page.component.html',
  styleUrl: './products-page.component.css'
})
export class ProductsPageComponent {

  json_productList: any
  productList: any
  filteredList: ProductComponent[] = []
  tags: any
  pList = [
    new ProductComponent('Java gia olous', '22.80', "Blends", "1001", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia ligous', '47.50', "Blends", "1002", "../assets/java-logo.jpg"),
    new ProductComponent('Java gia kanenan', '89.20', "Blends", "1003", "../assets/java-logo.jpg"),
    new ProductComponent('See ', '0.99', "Tea", "1004", "../assets/C_Logo.png"),
    new ProductComponent('See sharp', '14.70', "Tea", "1005", "../assets/C_Logo.png"),
    new ProductComponent('See plus plus', '185.00', "Tea", "1006", "../assets/C_Logo.png"),
    new ProductComponent('Bython', '89.20', "Utilities", "1007", "../assets/java-logo.jpg")
  ]

  x = {
    name: "",
    price: 0,
    description: "",
  }

  receiveMSG($event: ProductComponent[]) {
    this.filteredList = $event
  }

  constructor(private http: HttpClient) {

    for (var p of this.pList) {
      this.x.name = p.pName
      this.x.price = parseFloat(p.pPrice)
      this.x.description = p.pTag

      this.http.post('http://localhost:8080/v1/post/product', JSON.stringify(this.x));
    }
    this.getPrdList()

  }


  async getPrdList() {
    // const response = await fetch('http://localhost:8080/v1/get/products', {
    //   method: 'GET'
    // }).then((response) => response.json())
    // //let text = await response.text();
    // return response;

    // this.http.get('http://localhost:8080/GetProducts').
    //   subscribe((json_productList) => {
    //     console.log(json_productList);
    //     this.json_productList = json_productList;
    //   });

    //this.json_productList = this.http.get('http://localhost:8080/GetProducts')

    //this.json_productList = this.getPrdList()


    const json = await fetch('http://localhost:8080/v1/get/products', {
      method: 'GET'
    }).then((response) => response.json())

    console.log(json);

    this.productList = json
    console.log(this.productList);

    // for (var obj of this.json_productList) {
    //   this.productList.append(JSON.parse(obj))
    // }


    this.tags = ["Blends", "Tea", "Utilities"]
    let i = 0
    for (var obj of json) {
      i = i + 1
      console.log(i);
      this.filteredList.push(new ProductComponent(obj.name, obj.price, obj.description, obj.id, "../assets/java-logo.jpg"))
    }

    return 0
  }

  // productList = [
  //   new ProductComponent('Java gia olous', '22.80', "Blends", "1001", "../assets/java-logo.jpg"),
  //   new ProductComponent('Java gia ligous', '47.50', "Blends", "1002", "../assets/java-logo.jpg"),
  //   new ProductComponent('Java gia kanenan', '89.20', "Blends", "1003", "../assets/java-logo.jpg"),
  //   new ProductComponent('See ', '0.99', "Tea", "1004", "../assets/C_Logo.png"),
  //   new ProductComponent('See sharp', '14.70', "Tea", "1005", "../assets/C_Logo.png"),
  //   new ProductComponent('See plus plus', '185.00', "Tea", "1006", "../assets/C_Logo.png"),
  //   new ProductComponent('Bython', '89.20', "Utilities", "1007", "../assets/java-logo.jpg")

  // ]

}
