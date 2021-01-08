import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { NgModule } from '@angular/core'; 
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { HomeComponent } from './components/home/home.component';
import { CookieService } from "ngx-cookie-service";
import { HttpClientModule } from "@angular/common/http";
import { SelectCourseComponent } from './components/home/select-course/select-course.component';
import { WeekdayPipe } from './weekday.pipe';
import { CourseTableComponent } from './components/home/course-table/course-table.component';
import { IndexComponent } from './components/home/index/index.component';
import { GradeComponent } from './components/home/grade/grade.component';
import { TeachCourseComponent } from './components/home/teach-course/teach-course.component';
import { CollegeInfoComponent } from './components/home/college-info/college-info.component';
import { AssignCourseComponent } from './components/home/assign-course/assign-course.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent,
    SelectCourseComponent,
    WeekdayPipe,
    CourseTableComponent,
    IndexComponent,
    GradeComponent,
    TeachCourseComponent,
    CollegeInfoComponent,
    AssignCourseComponent,
    PageNotFoundComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [CookieService],
  bootstrap: [AppComponent]
})
export class AppModule { }
