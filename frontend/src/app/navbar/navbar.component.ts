import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterModule } from '@angular/router';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css'
})
export class NavbarComponent {
  @Input() currentTab = ""

  constructor(private router: Router) { }
  session = localStorage.getItem('session');

  async logout() {

    // const json = await fetch("http://localhost:8080/v1/post/logout_user", {
    //   method: 'POST',
    //   //body: this.post_str
    // }).then((response) => response.json())

    localStorage.removeItem('session');
    this.router.navigate(['/login']);
    //this.router.navigate(['/homepage']);
  }
}
