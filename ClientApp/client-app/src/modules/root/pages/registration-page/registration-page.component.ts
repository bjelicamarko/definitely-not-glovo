import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UserDTO } from 'src/modules/shared/models/UserDTO';
import { AppService } from '../../services/app.service';

@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.scss']
})
export class RegistrationPageComponent implements OnInit {

  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private appService: AppService
  ) { 
    this.form = this.fb.group({
      email: [null, Validators.required],
      password: [null, Validators.required],
      firstName: [null, Validators.required],
      lastName: [null, Validators.required],
      contact: [null, Validators.required],
    });
  }

  ngOnInit(): void {
  }

  submit() {
    const userDTO: UserDTO = {
      Email: this.form.value.email,
      Password: this.form.value.password,
      FirstName: this.form.value.firstName,
      LastName: this.form.value.lastName,
      Contact: this.form.value.contact
    };

    console.log(userDTO)

    // this.appService.register(userDTO)
    // .subscribe((response) => {
    //   console.log(response.body);
    // })
  }
}
