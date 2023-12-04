import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../navbar/navbar.component';
import { ProductGridComponent } from "../product-grid/product-grid.component";

@Component({
    selector: 'app-test-page',
    standalone: true,
    templateUrl: './test-page.component.html',
    styleUrl: './test-page.component.css',
    imports: [CommonModule, NavbarComponent, ProductGridComponent]
})
export class TestPageComponent {

}
