import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ConformationDialogComponent } from 'src/modules/shared/components/conformation-dialog/conformation-dialog.component';
import { UserDTO } from 'src/modules/shared/models/UserDTO';
import { UserDTOMessage } from 'src/modules/shared/models/UserDTOMessage';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UsersService } from '../../services/users.service';

@Component({
  selector: 'app-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.scss']
})
export class UserCardComponent implements OnInit {

  @Input() user: UserDTO = {
    Id: 0,
    Email: '',
    FirstName: '',
    LastName: '',
    Contact: '',
    Role: '',
    Banned: false,
    Image: '',
    Password: '',
    ImagePath: '',
    Changed: false,
    RestaurantName: ''
  }

  @Output() renderList: EventEmitter<any> = new EventEmitter();
  
  constructor(public dialog: MatDialog,
    private usersService: UsersService,
    private snackBarService: SnackBarService,
    private router: Router) { }

  ngOnInit(): void {
  }

  banUser(id: number): void {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Banning User",
        body: "You want to ban " + this.user.FirstName + " " + this.user.LastName + " ?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        this.usersService.banUser(id)
        .subscribe((response) => {
          var temp = response.body as UserDTOMessage;
          this.snackBarService.openSnackBar(temp.Message)
          // if (message === "user successfully banned") {
          //   this.user.Banned = true
          // }
          this.user = temp.UserDTO;
        })
      }
    })
  }

  unbanUser(id: number): void {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Unbanning User",
        body: "You want to unban " + this.user.FirstName + " " + this.user.LastName + " ?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        this.usersService.unbanUser(id)
        .subscribe((response) => {
          var temp = response.body as UserDTOMessage;
          this.snackBarService.openSnackBar(temp.Message)
          // if (message === "user successfully unbanned") {
          //   this.user.Banned = false
          // }
          this.user = temp.UserDTO;
        })
      }
    })
  }

  deleteUser(id: number): void {
    this.dialog.open(ConformationDialogComponent, {
      data:
      {
        title: "Deleting User",
        body: "You want to remove " + this.user.FirstName + " " + this.user.LastName + " ?"
      },
    }).afterClosed().subscribe(result => {
      if (result) {
        this.usersService.deleteUser(id)
        .subscribe((response) => {
          var temp = response.body as UserDTOMessage;
          this.snackBarService.openSnackBar(temp.Message)
          this.renderList.emit(null);
        })
      }
    })
  }

  updateUser(id: number): void {
    //this.router.navigate(["/app/main/admin/profile-info/" + id]);
    this.router.navigate(["/app/main/admin/user-info/" + id]);
  }
}
