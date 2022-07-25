import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { UserDTO } from 'src/modules/shared/models/UserDTO';
import { UserDTOMessage } from 'src/modules/shared/models/UserDTOMessage';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
import { UsersUtilsService } from 'src/modules/shared/services/users-utils.service';
import { UsersService } from '../../services/users.service';

@Component({
  selector: 'app-create-user-page',
  templateUrl: './create-user-page.component.html',
  styleUrls: ['./create-user-page.component.scss']
})
export class CreateUserPageComponent implements OnInit {
  roles: string[] = [
    'APPUSER', 'DELIVERER', 'EMPLOYEE'
  ];

  user: UserDTO = {
    Id: 0,
    Email: '',
    Password: '',
    FirstName: '',
    LastName: '',
    Contact: '',
    Role: '',
    Banned: false,
    Image: 'assets/user.jpg',
    ImagePath: 'assets/user.jpg',
    Changed: false
  }
  
  public selectedFile: File | undefined

  public userIdFromRoute: number

  constructor(private usersService: UsersService,
    private usersUtilsService: UsersUtilsService,
    private snackBarService: SnackBarService,
    private router: Router,
    private route: ActivatedRoute
  ) {
    this.userIdFromRoute = 0
  }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.userIdFromRoute = Number(routeParams.get('userId'));

    if (this.userIdFromRoute !== 0) {
      this.usersUtilsService.findUserById(this.userIdFromRoute)
      .subscribe((response) => {
        var temp = response.body as UserDTOMessage;
        this.user = temp.UserDTO;
        this.snackBarService.openSnackBar(temp.Message);
      })
    }
  }

  createUser() {
    if (this.user.Email && this.user.Password && this.user.FirstName &&
      this.user.LastName && this.user.Contact && this.user.Role && 
      this.user.Image && this.user.ImagePath !== 'assets/user.jpg') {
        let reader = new FileReader();
        reader.readAsDataURL(this.selectedFile!);
        reader.onload = () => {
          this.user.Image = reader.result;
          this.usersService.createUser(this.user)
          .subscribe((response) => {
            var temp = response.body as UserDTOMessage;
            this.snackBarService.openSnackBar(temp.Message);
            this.user = temp.UserDTO;
            this.router.navigate(["/app/main/admin/users"]);
          })
        };
        reader.onerror = function (error) {
          console.log('Error: ', error);
        };
    }
  }

  updateUser() {
    if (this.user.FirstName &&
      this.user.LastName && this.user.Contact && 
      this.user.Image && this.user.ImagePath !== 'assets/user.jpg') {
      if (this.user.Changed) {
        let reader = new FileReader();
        reader.readAsDataURL(this.selectedFile!);
        reader.onload = () => {
          this.user.Image = reader.result;
          this.usersUtilsService.updateUser(this.user)
          .subscribe((response) => {
            var temp = response.body as UserDTOMessage;
            this.snackBarService.openSnackBar(temp.Message);
            this.user = temp.UserDTO;
            this.router.navigate(["/app/main/admin/users"]);
          })
        };
        reader.onerror = function (error) {
          console.log('Error: ', error);
        };
      } else {
        this.usersUtilsService.updateUser(this.user)
        .subscribe((response) => {
          var temp = response.body as UserDTOMessage;
          this.snackBarService.openSnackBar(temp.Message);
          this.user = temp.UserDTO;
          this.router.navigate(["/app/main/admin/users"]);
        })
      }
    }
  }

  onFileChanged(event: any) {
    this.selectedFile = event.target.files[0]
    this.user.ImagePath = "images/" + this.selectedFile?.name as string;
    this.user.Changed = true
  }

}
