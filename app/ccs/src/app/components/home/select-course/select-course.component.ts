import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-select-course',
  templateUrl: './select-course.component.html',
  styleUrls: ['./select-course.component.less']
})
export class SelectCourseComponent implements OnInit {

  loading:boolean;
  courseList:any[];
  sid:Number;
  constructor(private http:HttpClient, private cookieService:CookieService) { }

  ngOnInit(): void {
    this.loading = true;
    this.courseList = [];
    this.sid = Number(this.cookieService.get('_id'));
    this.http.get('api/elective_course/'+this.sid).toPromise().then((data)=>{
      this.courseList = data["data"];
      this.loading = false
    }).catch((err)=>{
      alert(err.error)
    })
  }

  select(cid) {
    this.http.post('api/select_elective_course', {sid:this.sid, cid:cid}).toPromise().then((data)=>{
      this.courseList = data["data"];
    }).catch((err)=>{
      alert(JSON.stringify(err.error));
    })
  }

  cancel(cid) {
    let req = {'sid': this.sid, 'cid':cid}
    this.http.post('api/cancel_elective_course', req).toPromise().then((data)=>{
      this.courseList = data["data"];
    }).catch((err)=>{
      alert(JSON.stringify(err.error))
    })
  }

}
