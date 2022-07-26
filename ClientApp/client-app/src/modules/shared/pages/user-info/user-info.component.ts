import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { UserDTO } from '../../models/UserDTO';
import { UserDTOMessage } from '../../models/UserDTOMessage';
import { SnackBarService } from '../../services/snack-bar.service';
import { UsersUtilsService } from '../../services/users-utils.service';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.scss']
})
export class UserInfoComponent implements OnInit {

  user: UserDTO = {
    Id: 0,
    Email: '',
    Password: '',
    FirstName: '',
    LastName: '',
    Contact: '',
    Role: '',
    Banned: false,
    Image: null,
    ImagePath: '',
    Changed: false,
    RestaurantName: ''
  }
  
  public selectedFile: File | undefined

  constructor(
    private authService: AuthService,
    private usersUtilsService: UsersUtilsService,
    private snackBarService: SnackBarService
  ) { }

  ngOnInit(): void {
    var info = this.authService.getInfo()
    this.usersUtilsService.findUserById(info.Id)
    .subscribe((response) => {
      var temp = response.body as UserDTOMessage;
      this.user = temp.UserDTO;
      this.snackBarService.openSnackBar(temp.Message);
    })

  }

  updateUser() {
    if (this.user.FirstName && this.user.LastName && this.user.Contact
      && this.user.Image) {
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
