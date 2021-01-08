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
  constructor(private http:HttpClient) { }

  ngOnInit(): void {
    this.loading = true;
    this.http.get('api/elective_course/1806100187').toPromise().then((data)=>{
      this.courseList = data["data"];
      this.loading = false
      console.log(this.courseList);
    }).catch((err)=>{
      alert(err.error.msg)
    })
  }

  select(cid) {
    this.http.post('api/select_elective_course', {sid:1806100187, cid:cid}).toPromise().then((data)=>{
      this.courseList = data["data"];
    }).catch((err)=>{
      alert(err.error.msg);
    })
  }

  cancel(cid) {
    this.http.post('api/cancel_elective_course', {sid: 1806100187, cid:cid}).toPromise().then((data)=>{
      this.courseList = data["data"];
    }).catch((err)=>{
      alert(err.error.msg)
    })
  }

}
