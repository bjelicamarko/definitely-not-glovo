import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserCardComponent } from 'src/modules/admin/components/user-card/user-card.component';
import { ImageMessage } from '../../models/ImageMessage';
import { UserDTO } from '../../models/UserDTO';
import { UtilsService } from '../../services/utils.service';


@Component({
  selector: 'app-profile-info-page',
  templateUrl: './profile-info-page.component.html',
  styleUrls: ['./profile-info-page.component.scss']
})
export class ProfileInfoPageComponent implements OnInit {

  
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
  
  constructor(private route: ActivatedRoute,
    private utilsService: UtilsService) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const userIdFromRoute = Number(routeParams.get('userId'));
  
    this.utilsService.findUserById(userIdFromRoute)
    .subscribe(response => {
      this.user = response.body as UserDTO;
      console.log(this.user);
    })
  }

  onFileChanged(event: any) {
    this.selectedFile = event.target.files[0]
  }

  onUpload() {
    let reader = new FileReader();
    // let utilsService = this.utilsService
    // let path = this.selectedFile?.name as string
    // let id = this.user.Id
    reader.readAsDataURL(this.selectedFile!);
    reader.onload = () => {
     //me.modelvalue = reader.result;
      console.log(reader.result);
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
}
