import { Component, EventEmitter, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductComponent } from '../product/product.component';
import { ProductsPageComponent } from '../products-page/products-page.component';
import { FormsModule } from '@angular/forms';
import { Output } from '@angular/core';

@Component({
    selector: 'app-product-filters',
    standalone: true,
    templateUrl: './product-filters.component.html',
    styleUrl: './product-filters.component.css',
    imports: [CommonModule, ProductsPageComponent, FormsModule]
})

export class ProductFiltersComponent {

    @Input() list: any[] = [];
    @Input() tags: string[] = []
    @Output() msg = new EventEmitter();

    res: ProductComponent[] = this.list

    searchText = "";
    selected_tag = "";
    radio = false;

    filteredArr(list: ProductComponent[], arg1: string) {
        this.res = []
        if (!this.list) {
            return [];
        }

        if (!arg1) {
            if (this.radio == true) {
                this.list.forEach((p, index) => {

                    if (p.pTag === this.selected_tag) {
                        this.res.push(p)
                    }
                });
            }
            else {
                this.list.forEach((p, index) => {
                    this.res.push(p)
                });
            }
            this.msg.emit(this.res)
            return this.res;
        }

        let i, j: string
        if (this.radio == false) {
            this.list.forEach((p, index) => {

                i = p.pName.toLocaleLowerCase();

                if (i.includes(arg1)) {
                    this.res.push(p)
                }
            });
        }
        else {

            this.list.forEach((p, index) => {

                i = p.pName.toLocaleLowerCase();
                j = p.pTag;

                if (i.includes(arg1) && j === this.selected_tag) {
                    this.res.push(p)
                }

            });
        }
        this.msg.emit(this.res)
        return this.res;
    }

    select_tag() {
        this.selected_tag = this.selected_tag;
        this.radio = true;
    }

    clear_filters() {
        this.searchText = "";
        this.selected_tag = "";
        this.radio = false;

        this.res = []
        this.list.forEach((p, index) => {
            this.res.push(p)
        });
        this.msg.emit(this.res)
    }
}
