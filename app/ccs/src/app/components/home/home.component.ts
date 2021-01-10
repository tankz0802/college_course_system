import { UserInfoService } from './../../user-info.service';
import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap, Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.less']
})
export class HomeComponent implements OnInit {

  popFlag = true;
  selected:string;
  selectList:string[];
  role:string;
  showSelectFLag:boolean;
  name:string;
  constructor(private router:Router, private cookieService:CookieService) { }

  ngOnInit(): void {
    this.showSelectFLag = false;
    this.role = this.cookieService.get('_role');
    this.name = this.cookieService.get('_name');
    if(this.cookieService.get('_select') == undefined) {
      this.selected = '首页';
      this.cookieService.set('_select', this.selected)
    }else{
      this.selected = this.cookieService.get('_select');
    }
  }

  index() {
    this.selected = '首页';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/index')
  }

  selectCourse() {
    this.selected = '选课';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/select_course')
  }

  assignCourse() {
    this.selected = '排课';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/assign_course')
  }

  grade() {
    this.selected = '成绩';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/grade')
  }

  teachCourse() {
    this.selected = '我的授课',
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/teach_course')
  }

  courseTable() {
    this.selected = '课程表';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/course_table')
  }

  collegeInfo() {
    this.selected = '学院信息';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/college_info')
  }

  logout() {
    this.cookieService.delete('_id');
    this.cookieService.delete('_role');
    this.cookieService.delete('_name');
    this.cookieService.delete('_delete');
    this.router.navigateByUrl('/login')
  }

}
