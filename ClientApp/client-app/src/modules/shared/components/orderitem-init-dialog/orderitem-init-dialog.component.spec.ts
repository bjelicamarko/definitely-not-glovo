import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrderitemInitDialogComponent } from './orderitem-init-dialog.component';

describe('OrderitemInitDialogComponent', () => {
  let component: OrderitemInitDialogComponent;
  let fixture: ComponentFixture<OrderitemInitDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OrderitemInitDialogComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OrderitemInitDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
