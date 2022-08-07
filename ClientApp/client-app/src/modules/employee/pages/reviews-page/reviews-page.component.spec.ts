import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReviewsPageComponent } from './reviews-page.component';

describe('ReviewsPageComponent', () => {
  let component: ReviewsPageComponent;
  let fixture: ComponentFixture<ReviewsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReviewsPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReviewsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
