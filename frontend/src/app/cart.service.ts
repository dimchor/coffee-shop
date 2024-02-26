import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { ProductComponent } from './product/product.component';

@Injectable({
  providedIn: 'root'
})
export class CartService {

  cartList: ProductComponent[] = [];

  getList() {
    return this.cartList;
  }

  pushList(pr: ProductComponent) {
    this.cartList.push(pr)
    console.log(this.cartList[0])
  }
}
