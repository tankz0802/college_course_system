import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-assign-course',
  templateUrl: './assign-course.component.html',
  styleUrls: ['./assign-course.component.less']
})
export class AssignCourseComponent implements OnInit {
  editLine:number;
  courseList:any[];
  showModal:boolean;
  timeRule = [{rule:'上午', select: false}, {rule:'下午', select:false}, {rule:'晚上', select:false}];
  dayRule = [{rule: '星期一', select:false},{rule: '星期二', select:false},{rule:'星期三', select:false}, {rule:'星期四', select: false},{rule:'星期五', select: false},{rule:'星期六', select:false},{rule:'星期日', select:false}];
  compulsoryFlag:boolean;
  class:any[];
  showDayRule:boolean;
  showTimeRule:boolean;
  showClass:boolean;
  selectCourseIndex:number;
  constructor(private http:HttpClient, private cookieService:CookieService) { }

  ngOnInit(): void {
    this.editLine = -1;
    this.showModal = false;
    this.courseList = [];
    this.compulsoryFlag = false;
    this.showDayRule = false;
    this.showTimeRule = false;
    this.class = [];
    this.http.get('api/teach_course_list/'+this.cookieService.get('_id')).toPromise().then((data)=>{
      console.log(data["data"]);
      this.courseList = data["data"];
    }).catch((err)=>{
      alert(err.error.msg);
    })

    this.http.get('api/class_list').toPromise().then((data)=>{
      for(let i=0;i<data["data"].length;i++) {
        this.class.push({rule:data["data"][i].name, select:false, id: data["data"][i].id});
      }
      console.log(this.courseList);
    }).catch((err)=>{
      alert(err.error)
    })
  }

  edit(i:number) {
    this.editLine = i;
  }

  showAssignCourseModal(i:number) {
    this.showModal = true;
    this.selectCourseIndex = i;
    if(String(this.courseList[i].category).includes("必修")) {
      this.compulsoryFlag = true;
    }else{
      this.compulsoryFlag = false;
    }
  }

  confirmEdit() {
    this.editLine = -1;
  }

  selectTime(i:number) {
    this.timeRule[i].select = !this.timeRule[i].select;
  }

  selectDay(i:number) {
    this.dayRule[i].select = !this.dayRule[i].select;
  }

  selectClass(i:number) {
    this.class[i].select = !this.class[i].select;
  }

  confirmAssign() {
    let assignCourseRequest = {
      time: [],
      weekday:[],
      class: [],
      cid: this.courseList[this.selectCourseIndex].id,
      tid: Number(this.cookieService.get('_id')),
    }
    this.timeRule.forEach((val, index)=>{
      if(val.select) {
        assignCourseRequest.time.push(index+1);
      }
    })
    this.dayRule.forEach((val, index)=>{
      if(val.select) {
        assignCourseRequest.weekday.push(index+1)
      }
    })
    
    this.class.forEach((val)=>{
      if(val.select) {
        assignCourseRequest.class.push(val.id);
      }
    })
    if(this.compulsoryFlag && assignCourseRequest.class.length === 0) {
      alert('必修课必须指定班级');
      return;
    }
    console.log(assignCourseRequest);
    this.http.post('api/assign_course', assignCourseRequest).toPromise().then((data)=>{

    }).catch((err)=>{
      alert(JSON.stringify(err.error));
    })
  }

}
