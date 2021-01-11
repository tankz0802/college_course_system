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
    if(this.cookieService.check('_id') && this.cookieService.check('_role') && this.cookieService.check('_name')) {

      this.router.navigateByUrl('home/'+this.cookieService.get('_select'));
    }
  }

  login() {
    this.http.post('api/login', this.loginForm).toPromise().then((data)=>{
      console.log(111)
      this.cookieService.set('_id', this.loginForm.id,);
      this.cookieService.set('_name', data["name"]);
      this.cookieService.set('_role', data["role"]);
      this.cookieService.set('_select', 'index');
      this.router.navigateByUrl("/home/index");
    }).catch((err)=>{
      alert(err.error)
    })
  }

}
