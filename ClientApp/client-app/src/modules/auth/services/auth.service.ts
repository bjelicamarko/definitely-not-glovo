import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Token } from "@angular/compiler";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { Login } from "src/modules/shared/models/Login";


@Injectable({
    providedIn: "root",
})
export class AuthService {
    private headers = new HttpHeaders({ "Content-Type": "application/json" });

    constructor(private http: HttpClient) {}

    login(auth: Login): Observable<Token> {
        return this.http.post<Token>("not-glovo/api/users/login", auth, {
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
    
}