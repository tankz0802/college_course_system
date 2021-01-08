import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CollegeInfoComponent } from './college-info.component';

describe('CollegeInfoComponent', () => {
  let component: CollegeInfoComponent;
  let fixture: ComponentFixture<CollegeInfoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CollegeInfoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CollegeInfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
