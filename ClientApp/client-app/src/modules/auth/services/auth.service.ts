import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Token } from "@angular/compiler";
import { Injectable } from "@angular/core";
import { JwtHelperService } from "@auth0/angular-jwt";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment.prod";
import { Login } from "src/modules/shared/models/Login";


@Injectable({
    providedIn: "root",
})
export class AuthService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {
        this.headers.set('Access-Control-Allow-Origin', '*');
    }

    login(auth: Login): Observable<Token> {
        return this.http.post<Token>(environment.url + "/api/users/login", auth, {
          headers: this.headers,
          responseType: "json",
        });
    }
    
    logout(): void {
        localStorage.removeItem("user");
    }

    isLoggedIn(): boolean {
        if (!localStorage.getItem("user")) {
          return false;
        }
        return true;
    }

    getInfo(): any {
        const token = localStorage.getItem("user");

        const jwt: JwtHelperService = new JwtHelperService();
        const info = jwt.decodeToken(token!);
        return info
    }
    
}