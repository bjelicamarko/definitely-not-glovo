import { Component, OnInit } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.scss']
})
export class MainPageComponent implements OnInit {

  public role: string;

  constructor() { 
    this.role = "";
  }

  ngOnInit(): void {
    const item = localStorage.getItem("user");
    if (item) {
      const jwt: JwtHelperService = new JwtHelperService();
      this.role = jwt.decodeToken(item).role;
      console.log(this.role)
    }
  }

  checkRole() {
    const item = localStorage.getItem("user");
    if (item) {
      const jwt: JwtHelperService = new JwtHelperService();
      this.role = jwt.decodeToken(item).role;
      console.log(this.role)
    }
  }
}
