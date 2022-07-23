import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';
import { ConformationDialogComponent } from '../../components/conformation-dialog/conformation-dialog.component';
import { ImageMessage } from '../../models/ImageMessage';
import { UserDTO } from '../../models/UserDTO';
import { UserDTOMessage } from '../../models/UserDTOMessage';
import { SnackBarService } from '../../services/snack-bar.service';
import { UtilsService } from '../../services/utils.service';


@Component({
  selector: 'app-profile-info-page',
  templateUrl: './profile-info-page.component.html',
  styleUrls: ['./profile-info-page.component.scss']
})
export class ProfileInfoPageComponent implements OnInit {

  public userIdFromRoute: number

  public  user: UserDTO = {
    Id: 0,
    Email: '',
    FirstName: '',
    LastName: '',
    Contact: '',
    Role: '',
    Banned: false,
    Image: ''
  }

  public selectedFile: File | undefined
  
  constructor(public dialog: MatDialog,
    private route: ActivatedRoute,
    private utilsService: UtilsService,
    private snackBarService: SnackBarService) { 
      this.userIdFromRoute = 0
    }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.userIdFromRoute = Number(routeParams.get('userId'));
  
    if (this.userIdFromRoute !== 0) {
      this.utilsService.findUserById(this.userIdFromRoute)
      .subscribe(response => {
        this.user = response.body as UserDTO;
        console.log(this.user);
      })
    } else { // znaci kreiramo novog korisnika 
      console.log(this.userIdFromRoute)
    }
  }

  onFileChanged(event: any) {
    this.selectedFile = event.target.files[0]
  }

  onUpload() {
    let reader = new FileReader();
    reader.readAsDataURL(this.selectedFile!);
    reader.onload = () => {
      //console.log(reader.result);
      var temp: ImageMessage = {
        Image: reader.result,
        Path: this.selectedFile?.name as string,
        Id: this.user.Id
      } 
      this.utilsService.saveImageUser(temp)
      .subscribe((response) => {
        this.user = response.body as UserDTO;
      })

    };
    reader.onerror = function (error) {
     console.log('Error: ', error);
    };
  }

  updateProfile() {
    if (this.user.FirstName && this.user.LastName && this.user.Contact) {
      this.dialog.open(ConformationDialogComponent, {
        data:
        {
          title: "Updating Profile",
          body: "You want to update profile ?"
        },
      }).afterClosed().subscribe(result => {
        if (result) {
          this.utilsService.updateUser(this.user)
          .subscribe(response => {
            var temp = response.body as UserDTOMessage;
            this.user = temp.UserDTO;
            this.snackBarService.openSnackBar(temp.Message);
          })
        }
      })
    }
  }
}
