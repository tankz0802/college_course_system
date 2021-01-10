import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-grade',
  templateUrl: './grade.component.html',
  styleUrls: ['./grade.component.less']
})
export class GradeComponent implements OnInit {

  gradeList:any[];
  constructor(private http:HttpClient, private cookieService:CookieService) { }

  ngOnInit(): void {
    this.gradeList = [];
    this.http.get('api/grade_list/'+this.cookieService.get('_id')).toPromise().then((data)=>{
      this.gradeList = data["data"];
    }).catch((err)=>{
      alert(err.error)
    })
    
  }

}
