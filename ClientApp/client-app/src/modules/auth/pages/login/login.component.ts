import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Login } from 'src/modules/shared/models/Login';
import { AuthService } from '../../services/auth.service';
import { JwtHelperService } from "@auth0/angular-jwt";
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private snackBarService: SnackBarService

  ) { 
    this.form = this.fb.group({
      email: [null, Validators.required],
      password: [null, Validators.required],
    });
  }

  ngOnInit(): void {
  }

  submit() {
    const auth: Login = {
      Email: this.form.value.email,
      Password: this.form.value.password,
    };

    this.authService.login(auth)
    .subscribe((result: any) => {

      if (result !== null) {
        this.snackBarService.openSnackBar("Successful login!");

        const token = JSON.stringify(result);
        localStorage.setItem("user", token);

        const jwt: JwtHelperService = new JwtHelperService();
        const info = jwt.decodeToken(token);
        const role = info.role;

        if (role === "ADMIN") {
          this.router.navigate(["/app/main/admin/users"]);
        } else if (role === "APPUSER") {
          this.router.navigate(["app/main/appuser/restaurants"]);
        } else if (role === "DELIVERER") {
          this.router.navigate(["app/main/deliverer/orders"]);
        } else if (role === "EMPLOYEE") {
          this.router.navigate(["app/main/employee/orders"]);
        } else {
          this.router.navigate(["app/main"]);
        }
      }
      
    });
  }
}
