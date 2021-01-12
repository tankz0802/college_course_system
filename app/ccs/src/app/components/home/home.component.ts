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
      this.selected = 'index';
      this.cookieService.set('_select', this.selected)
    }else{
      this.selected = this.cookieService.get('_select');
    }
    this.router.navigateByUrl('/home/'+this.selected)
  }

  index() {
    this.selected = 'index';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/index')
  }

  selectCourse() {
    this.selected = 'select_course';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/select_course')
  }

  assignCourse() {
    this.selected = 'assign_course';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/assign_course')
  }

  grade() {
    this.selected = 'grade';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/grade')
  }

  teachCourse() {
    this.selected = 'teach_course',
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/teach_course')
  }

  courseTable() {
    this.selected = 'course_table';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/course_table')
  }

  collegeInfo() {
    this.selected = 'college_info';
    this.cookieService.set('_select', this.selected)
    this.router.navigateByUrl('/home/college_info')
  }

  logout() {
    this.cookieService.delete('_select');
    this.cookieService.delete('_role');
    this.cookieService.delete('_id');
    this.cookieService.delete('_name');
    this.router.navigateByUrl('/login')
  }

}
