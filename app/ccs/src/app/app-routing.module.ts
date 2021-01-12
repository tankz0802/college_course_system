import { StudentListComponent } from './components/home/student-list/student-list.component';
import { RoleGuard } from './role.guard';
import { AssignCourseComponent } from './components/home/assign-course/assign-course.component';
import { CollegeInfoComponent } from './components/home/college-info/college-info.component';
import { TeachCourseComponent } from './components/home/teach-course/teach-course.component';
import { GradeComponent } from './components/home/grade/grade.component';
import { IndexComponent } from './components/home/index/index.component';
import { CourseTableComponent } from './components/home/course-table/course-table.component';
import { SelectCourseComponent } from './components/home/select-course/select-course.component';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/login/login.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule, CanActivate } from '@angular/router';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';


const routes: Routes = [{
  path: 'login',
  component: LoginComponent
},
{
  path: 'home',
  component: HomeComponent,
  canActivate: [RoleGuard],
  children: [
    {
      path: 'index',
      component: IndexComponent
    },
    {
      path: 'select_course',
      component: SelectCourseComponent
    },
    {
      path: 'assign_course',
      component: AssignCourseComponent
    },
    {
      path: 'course_table',
      component: CourseTableComponent
    },
    {
      path: 'grade',
      component: GradeComponent
    },
    {
      path: 'teach_course',
      component: TeachCourseComponent
    },
    {
      path: 'college_info',
      component: CollegeInfoComponent
    },
    {
      path: 'student_list',
      component: StudentListComponent
    }
  ]
},
{path: '', redirectTo: 'login', pathMatch: 'full'},
{ path: '**', component: LoginComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
