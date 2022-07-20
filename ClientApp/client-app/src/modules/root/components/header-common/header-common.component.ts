import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/modules/auth/services/auth.service';

@Component({
  selector: 'app-header-common',
  templateUrl: './header-common.component.html',
  styleUrls: ['./header-common.component.scss']
})
export class HeaderCommonComponent implements OnInit {

  constructor(
    private authService: AuthService
  ) { }

  ngOnInit(): void {
  }

  logout(): void {
    this.authService.logout();
  }
}
