import { Component, EventEmitter, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SearchFilterPipe } from '../search-filter.pipe';
import { ProductComponent } from '../product/product.component';
import { ProductsPageComponent } from '../products-page/products-page.component';
import { FormsModule } from '@angular/forms';
import { Output } from '@angular/core';

@Component({
    selector: 'app-product-filters',
    standalone: true,
    templateUrl: './product-filters.component.html',
    styleUrl: './product-filters.component.css',
    imports: [CommonModule, SearchFilterPipe, ProductsPageComponent, FormsModule]
})

export class ProductFiltersComponent {

    @Input() list: any[] = [];
    @Output() msg = new EventEmitter();

    filteredArr(list: ProductComponent[], arg1: string) {
        this.res = []
        if (!this.list) {
            return [];
        }
        if (!arg1) {
            alert("no string");
            return this.list;
        }

        let i: string
        this.list.forEach((p, index) => {

            i = p.pName.toLocaleLowerCase();

            if (i.includes(arg1)) {
                this.res.push(p)
            }
        });
        alert("pass");
        this.msg.emit(this.res)
        return this.res;
    }

    // filterPipe = new SearchFilterPipe();
    res: ProductComponent[] = this.list
    searchText = "";


}
