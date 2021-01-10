import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-teach-course',
  templateUrl: './teach-course.component.html',
  styleUrls: ['./teach-course.component.less']
})
export class TeachCourseComponent implements OnInit {

  courseList:any[];
  constructor(private http:HttpClient,private cookieService:CookieService) { }

  ngOnInit(): void {
    this.courseList = [];
    this.http.get('api/teach_course_list/'+this.cookieService.get('_id')).toPromise().then((data)=>{
      this.courseList = data["data"];
    }).catch((err)=>{
      alert(err.error.msg);
    })
  }

}
