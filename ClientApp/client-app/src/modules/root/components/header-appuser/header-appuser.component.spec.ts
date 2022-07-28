import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HeaderAppuserComponent } from './header-appuser.component';

describe('HeaderAppuserComponent', () => {
  let component: HeaderAppuserComponent;
  let fixture: ComponentFixture<HeaderAppuserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HeaderAppuserComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HeaderAppuserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
