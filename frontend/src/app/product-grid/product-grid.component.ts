import { Component, Input, Output, EventEmitter, OnInit, SimpleChanges, OnChanges } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductComponent } from '../product/product.component';
import { Router } from '@angular/router';
import { CartService } from '../cart.service';

@Component({
  selector: 'app-product-grid',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './product-grid.component.html',
  styleUrl: './product-grid.component.css'
})
export class ProductGridComponent implements OnChanges {
  private router: Router = new Router;
  ngOnChanges(changes: SimpleChanges): void {
    this.router.navigateByUrl('/products')
  }

  @Input() list: ProductComponent[] = [];
  @Input() filteredList: ProductComponent[] = [];

  constructor(public CartService: CartService) {
  }

  setCart(pr: ProductComponent): void {
    this.CartService.pushList(pr)
  }
}
