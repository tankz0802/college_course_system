import { CookieService } from 'ngx-cookie-service';
import { ParamMap } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import * as xlsx from 'xlsx';

@Component({
  selector: 'app-course-table',
  templateUrl: './course-table.component.html',
  styleUrls: ['./course-table.component.less']
})
export class CourseTableComponent implements OnInit {
  @ViewChild('epltable', { static: false }) epltable: ElementRef;
  
  courseTable = {
    'title': '',
    'point112': '',
    'point134': '',
    'point156': '',
    'point178': '',
    'point19':'',
    'point110': '',
    'point111': '',
    'point212': '',
    'point234': '',
    'point256': '',
    'point278': '',
    'point29':'',
    'point210': '',
    'point211': '',
    'point312': '',
    'point334': '',
    'point356': '',
    'point378': '',
    'point39':'',
    'point310': '',
    'point311': '',
    'point412': '',
    'point434': '',
    'point456': '',
    'point478': '',
    'point49':'',
    'point410': '',
    'point411': '',
    'point512': '',
    'point534': '',
    'point556': '',
    'point578': '',
    'point59':'',
    'point510': '',
    'point511': '',
    'point612': '',
    'point634': '',
    'point656': '',
    'point678': '',
    'point69':'',
    'point610': '',
    'point611': '',
    'point712': '',
    'point734': '',
    'point756': '',
    'point778': '',
    'point79':'',
    'point710': '',
    'point711': '',
  }
  courseList:any[];

  constructor(private http:HttpClient, private cookieService:CookieService) { }

  ngOnInit(): void {
    if(this.cookieService.get('_role') === 'student') {
      this.getStudentCourseTable();
    }else{
      this.getTeacherCourseTable();
    }
  }

  getStudentCourseTable() {
    this.http.get('api/student_course_table/' + this.cookieService.get('_id')).toPromise().then((data)=>{
      console.log(data["data"])
      this.courseList = data["data"];
      this.courseTable['title'] = this.cookieService.get('_name')+'的课表'
      this.courseList.forEach((val, index, array)=>{
        console.log(val);
        if(val.section_start < 9) {
          let point = 'point' + val.week_day + val.section_start+val.section_end;
          this.courseTable[point] += val.cname+'/'+val.week_start+'-'+val.week_end+'周/'+val.cid+'/'+val.category;
          console.log(this.courseTable[point]);
        }else{
          for(let i=val.section_start;i<=val.section_end;i++) {
            let point = 'point' + val.week_day + i;
            this.courseTable[point] += val.cname+'/'+val.week_start+'-'+val.week_end+'周/'+val.cid+'/'+val.category+'\n';
            console.log(this.courseTable[point])
          }
        }
      })
    }).catch((err)=>{
      alert(err.error.msg)
    })
  }

  getTeacherCourseTable() {
    this.http.get('api/teacher_course_table/' + this.cookieService.get('_id')).toPromise().then((data)=>{
      console.log(data["data"])
      this.courseList = data["data"];
      this.courseTable['title'] = this.cookieService.get('_name')+'老师的课表'
      this.courseList.forEach((val, index, array)=>{
        console.log(val);
        if(val.section_start < 9) {
          let point = 'point' + val.week_day + val.section_start+val.section_end;
          this.courseTable[point] += val.cname+'/'+val.week_start+'-'+val.week_end+'周/'+val.cid+'/'+val.category;
          console.log(this.courseTable[point]);
        }else{
          for(let i=val.section_start;i<=val.section_end;i++) {
            let point = 'point' + val.week_day + i;
            this.courseTable[point] += val.cname+'/'+val.week_start+'-'+val.week_end+'周/'+val.cid+'/'+val.category+'\n';
            console.log(this.courseTable[point])
          }
        }
      })
    }).catch((err)=>{
        alert(err.error.msg)
    })
  }

  exportToExcel() {
    const ws: xlsx.WorkSheet =   
    xlsx.utils.table_to_sheet(this.epltable.nativeElement);
    const wb: xlsx.WorkBook = xlsx.utils.book_new();
    xlsx.utils.book_append_sheet(wb, ws, 'Sheet1');
    xlsx.writeFile(wb, this.cookieService.get('_name')+'的课表.xlsx');
   }
}
