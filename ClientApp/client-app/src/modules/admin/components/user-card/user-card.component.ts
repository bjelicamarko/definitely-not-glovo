import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ConformationDialogComponent } from 'src/modules/shared/components/conformation-dialog/conformation-dialog.component';
import { UserDTO } from 'src/modules/shared/models/UserDTO';
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
    Image: ''
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
          var message = response.body?.message as string
          this.snackBarService.openSnackBar(message)
          if (message === "user successfully banned") {
            this.user.Banned = true
          }
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
          var message = response.body?.message as string
          this.snackBarService.openSnackBar(message)
          if (message === "user successfully unbanned") {
            this.user.Banned = false
          }
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
          var message = response.body?.message as string
          this.snackBarService.openSnackBar(message)
          this.renderList.emit(null);
        })
      }
    })
  }

  profileInfo(id: number): void {
    this.router.navigate(["/app/main/admin/profile-info/" + id]);
  }
}
