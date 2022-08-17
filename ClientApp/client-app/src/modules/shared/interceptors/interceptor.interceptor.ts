import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

@Injectable()
export class Interceptor implements HttpInterceptor {
  intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const item = localStorage.getItem("user");
    if (item) {
      const decodedItem = JSON.parse(item);
      const cloned = req.clone({
        headers: req.headers.set('Authorization', 'Bearer ' + decodedItem),
      });

      return next.handle(cloned);
    } else {
      return next.handle(req);
    }
  }
}