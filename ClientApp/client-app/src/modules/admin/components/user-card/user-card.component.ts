import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { UserDTO } from 'src/modules/shared/models/UserDTO';
import { UsersService } from '../../services/users.service';

@Component({
  selector: 'app-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.scss']
})
export class UserCardComponent implements OnInit {

  @Input() user: UserDTO = {
    Id: 0,
    Email: '',
    FirstName: '',
    LastName: '',
    Contact: '',
    Role: '',
    Banned: false,
    Image: ''
  }

  @Output() renderList: EventEmitter<any> = new EventEmitter();
  
  constructor(private usersService: UsersService) { }

  ngOnInit(): void {
  }

}
