import { CookieService } from 'ngx-cookie-service';
import { UserInfoService } from './user-info.service';
import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class RoleGuard implements CanActivate {

  constructor(private cookieService:CookieService) {}

  canActivate(): boolean  {
    let role = this.cookieService.get('_role');
    if(role === "" || role == undefined) {
      return false;
    }
    return true;
  }
  
}
