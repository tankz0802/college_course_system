import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-student-list',
  templateUrl: './student-list.component.html',
  styleUrls: ['./student-list.component.less']
})
export class StudentListComponent implements OnInit {

  studentList:any[];
  editLine:number;
  cid:string;
  editGrade:number;
  editGPA:number;
  constructor(private http:HttpClient, private route:ActivatedRoute) { }

  ngOnInit(): void {
    this.editLine = -1;
    this.studentList = [];
    this.cid=this.route.snapshot.queryParams["cid"]
    this.getCourseStudentGradeList()
  }

  getCourseStudentGradeList() {
    this.http.get('api/course_student_grade_list/'+this.cid).toPromise().then((data)=>{
      this.studentList = data["data"];
    }).catch((err)=>{
      alert(err.error)
    })
  }

  edit(i:number) {
    this.editLine = i;
    this.editGrade = this.studentList[i].grade;
    this.editGPA = this.studentList[i].gpa;
  }

  cancelEdit() {
    this.editLine = -1;
    this.editGrade = 0;
    this.editGPA = 0;
  }

  confirmEdit() {
    let grade = {
      sid: this.studentList[this.editLine].id,
      cid: this.cid,
      grade: this.editGrade,
      gpa: this.editGPA,
    }
    this.http.post('api/update_grade', grade).toPromise().then((data)=>{
      this.studentList = data["data"];
      this.editGPA = 0;
      this.editGrade = 0;
      this.editLine = -1;
      alert("修改成功");
    }).catch((err)=>{
      alert(err.error)
    })
  }

}
