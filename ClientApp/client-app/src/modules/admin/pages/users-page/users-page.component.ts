import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
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

  searchFormGroup: FormGroup;

  constructor(private fb: FormBuilder,
    private usersService: UsersService) { 
    this.users = [];
    this.pageSize = 5;
    this.currentPage = 1;
    this.totalSize = 1;
    this.searchFormGroup = this.fb.group({
      searchField: [''],
      userType: [''],
    }); 
  }

  ngOnInit(): void {
    this.usersService.findAllUsers(this.currentPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as UsersPageable;
        console.log(temp);
        this.totalSize = Number(temp.TotalElements);
        this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
        this.users = temp.Elements as UserDTO[];
    });

    this.onChanges();
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
    this.usersService.searchUsers(this.searchFormGroup.value.searchField, this.searchFormGroup.value.userType,
      newPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as UsersPageable;
        this.totalSize = Number(temp.TotalElements);
        this.setPagination((this.totalSize).toString(), (newPage - 1).toString());
        this.users = temp.Elements as UserDTO[];
    });

  }
  
  renderList() {
    this.usersService.findAllUsers(this.currentPage - 1, this.pageSize)
      .subscribe((response) => {
        var temp = response.body as UsersPageable;
        console.log(temp);
        this.totalSize = Number(temp.TotalElements);
        this.setPagination((this.totalSize).toString(), (this.currentPage-1).toString());
        this.users = temp.Elements as UserDTO[];
    });
  }

  onChanges(): void {
    this.searchFormGroup.valueChanges
    .subscribe(val => {
      if (!this.searchFormGroup.get('searchField')?.valid)
        return;
      
      this.usersService.searchUsers(val.searchField, val.userType, 
        0, this.pageSize).subscribe((response) => {
          if (response.body != null) {
            var temp = response.body as UsersPageable;
            this.totalSize = Number(temp.TotalElements);
            this.setPagination((this.totalSize).toString(), (0).toString());
            this.users = temp.Elements as UserDTO[];
            if (this.pagination) {
              this.pagination.setActivePage(1);
            } 
          }
        })
    })
  }
}
