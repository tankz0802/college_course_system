import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TeachCourseComponent } from './teach-course.component';

describe('TeachCourseComponent', () => {
  let component: TeachCourseComponent;
  let fixture: ComponentFixture<TeachCourseComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TeachCourseComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TeachCourseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
