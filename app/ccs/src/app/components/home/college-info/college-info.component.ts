import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-college-info',
  templateUrl: './college-info.component.html',
  styleUrls: ['./college-info.component.less']
})
export class CollegeInfoComponent implements OnInit {

  studentList:any[];
  teacherList:any[];
  courseList:any[];
  classList:any[];
  select:number;
  classIndex:number;
  constructor(private http:HttpClient) { }

  ngOnInit(): void {
    this.select = -1;
  }

  getStudentList() {
    this.http.get('api/student_list').toPromise().then((data)=>{
      this.studentList = data["data"];
      this.select = 1;
    }).catch((err)=>{
      alert(err.error)
    })
  }

  getTeacherList() {
    this.http.get('api/teacher_list').toPromise().then((data)=>{
      this.teacherList = data["data"];
      this.select = 2;

    }).catch((err)=>{
      alert(err.error)
    })
  }

  getCourseList() {
    this.http.get('api/course_list').toPromise().then((data)=>{
      this.courseList = data["data"];
      this.select = 3;
    }).catch((err)=>{
      alert(err.error)
    })
  }

  getClassList() {
    this.http.get('api/class_info_list').toPromise().then((data)=>{
      this.classList = data["data"];
      this.select = 4;
    }).catch((err)=>{
      alert(err.error)
    })
  }

  updateClassIndex(i:number) {
    console.log(i)
    if(i == this.classIndex) {
      this.classIndex = -1;
      return
    }
    this.classIndex = i;
  }

}
