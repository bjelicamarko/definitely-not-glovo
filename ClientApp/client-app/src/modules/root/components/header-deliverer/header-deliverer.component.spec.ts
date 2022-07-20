import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HeaderDelivererComponent } from './header-deliverer.component';

describe('HeaderDelivererComponent', () => {
  let component: HeaderDelivererComponent;
  let fixture: ComponentFixture<HeaderDelivererComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HeaderDelivererComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HeaderDelivererComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
