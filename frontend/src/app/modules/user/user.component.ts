import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from '@shared/models/user.model';
import { UserService } from '@shared/services/user.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.scss']
})
export class UserComponent implements OnInit {
  user!: User;

  constructor(
    private service: UserService,
    private route: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    this.getById();
  }

  getById(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.service.getById(Number(id)).subscribe({
      next: (value) => {
        if (value.success) {
          this.user = value.data;
        }
      },
      error: (err) => console.error(err),
    });
  }
}
