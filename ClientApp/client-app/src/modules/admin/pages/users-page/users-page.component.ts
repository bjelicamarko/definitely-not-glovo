import { Component, OnInit, ViewChild } from '@angular/core';
import { PaginationComponent } from 'src/modules/shared/components/pagination/pagination.component';
import { UserDTO } from 'src/modules/shared/models/UserDTO';
import { UsersPageable } from 'src/modules/shared/models/UsersPageable';
import { UsersService } from '../../services/users.service';

@Component({
  selector: 'app-users-page',
  templateUrl: './users-page.component.html',
  styleUrls: ['./users-page.component.scss']
})
export class UsersPageComponent implements OnInit {
  @ViewChild(PaginationComponent) pagination?: PaginationComponent;
  pageSize: number;
  currentPage: number;
  totalSize: number;
  users: UserDTO[];

  constructor(private usersService: UsersService) { 
    this.users = [];
    this.pageSize = 2;
    this.currentPage = 1;
    this.totalSize = 1;
  }

  ngOnInit(): void {
    this.usersService.getUsers(this.currentPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as UsersPageable;
        console.log(temp);
        this.totalSize = Number(temp.TotalElements);
        this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
        this.users = temp.Elements as UserDTO[];
    });
  }

  setPagination(totalItemsHeader: string | null, currentPageHeader: string | null) {
    if (totalItemsHeader) {
      this.totalSize = parseInt(totalItemsHeader);
    }
    if (currentPageHeader) {
      this.currentPage = parseInt(currentPageHeader);
    }
  }

  changePage(newPage: number) {
    this.usersService.getUsers(newPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as UsersPageable;
        console.log(temp);
        this.totalSize = Number(temp.TotalElements);
        this.setPagination((this.totalSize).toString(), (newPage - 1).toString());
        this.users = temp.Elements as UserDTO[];
    });

  }
  
}
