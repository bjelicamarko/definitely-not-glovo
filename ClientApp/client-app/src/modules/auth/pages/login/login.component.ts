import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Login } from 'src/modules/shared/models/Login';
import { AuthService } from '../../services/auth.service';
import { JwtHelperService } from "@auth0/angular-jwt";

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

    console.log(auth)
    
    this.authService.login(auth)
    .subscribe((result: any) => {
      const token = JSON.stringify(result);
      localStorage.setItem("user", token);

      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(token);
      console.log(token)
      console.log(info)
    });
  }
}
