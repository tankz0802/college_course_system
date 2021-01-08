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
  constructor(private route:ActivatedRoute, private http:HttpClient,private router:Router,private cookieService:CookieService) { }

  ngOnInit(): void {
    this.role = this.cookieService.get('_role');
  }

  index() {
    this.selected = '首页';
    this.router.navigateByUrl('/home/index')
  }

  selectCourse() {
    this.selected = '选课';
    this.router.navigateByUrl('/home/select_course')
  }

  assignCourse() {
    this.selected = '排课';
    this.router.navigateByUrl('/home/assign_course')
  }

  grade() {
    this.selected = '成绩';
    this.router.navigateByUrl('/home/grade')
  }

  teachCourse() {
    this.selected = '我的授课',
    this.router.navigateByUrl('/home/teach_course')
  }

  courseTable() {
    this.selected = '课程表';
    this.router.navigateByUrl('/home/course_table')
  }

  collegeInfo() {
    this.selected = '学院信息';
    this.router.navigateByUrl('/home/college_info')
  }

}
