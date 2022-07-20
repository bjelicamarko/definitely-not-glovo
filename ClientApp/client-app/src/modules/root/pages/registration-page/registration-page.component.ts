import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ResponseMessage } from 'src/modules/shared/models/ResponseMessage';
import { UserDTO } from 'src/modules/shared/models/UserDTO';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';
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
    private appService: AppService,
    private snackBarService: SnackBarService,
    private router: Router,
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

    this.appService.register(userDTO)
    .subscribe((response) => {
      console.log(response.body)
      var message = response.body?.message as string
      this.snackBarService.openSnackBar(message)

      if (message === "registration succeeded") {
        this.router.navigate(["/app/auth/login"]);
      }

    })
  }
}
