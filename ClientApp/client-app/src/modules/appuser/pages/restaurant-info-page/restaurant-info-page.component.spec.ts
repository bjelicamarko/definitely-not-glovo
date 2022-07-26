import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RestaurantInfoPageComponent } from './restaurant-info-page.component';

describe('RestaurantInfoPageComponent', () => {
  let component: RestaurantInfoPageComponent;
  let fixture: ComponentFixture<RestaurantInfoPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RestaurantInfoPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RestaurantInfoPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
