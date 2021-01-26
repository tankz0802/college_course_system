import { CookieService } from 'ngx-cookie-service';
import { UserInfoService } from './../../user-info.service';
import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less']
})
export class LoginComponent implements OnInit {


  loginForm = {
    id: '',
    password: '',
  }

  constructor(private http: HttpClient, private router: Router, private cookieService:CookieService) { }
  

  ngOnInit(): void {
    
  }

  login() {
    console.log('login');
    this.http.post('api/login', this.loginForm).toPromise().then((data)=>{
      this.cookieService.set('_id', this.loginForm.id, 1, '/', null, false, "Lax");
      this.cookieService.set('_name', data["name"], 1 , '/', null, false, "Lax");
      this.cookieService.set('_role', data["role"], 1 , '/', null, false, "Lax");
      this.cookieService.set('_select', 'index', 1 , '/', null, false, "Lax");
      this.router.navigateByUrl("/home/index");
    }).catch((err)=>{
      alert(JSON.stringify(err.error))
    })
  }

}
