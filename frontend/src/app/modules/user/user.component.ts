import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute } from '@angular/router';
import { User } from '@shared/models/user.model';
import { UserService } from '@shared/services/user.service';
import { UserEditComponent } from './user-edit/user-edit.component';

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
    private _matDialog: MatDialog,
    private _snackBar: MatSnackBar,
  ) { }

  ngOnInit(): void {
    this.getById();
  }

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action, {
      horizontalPosition: 'end',
      verticalPosition: 'top',
      duration: 5000,
      panelClass: ['custom-snackbar'],
    });
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

  edit(item: User): void {
    const dialogRef = this._matDialog.open(UserEditComponent, {
      width: '600px',
      height: '500px',
      data: item,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.openSnackBar('Produto editado com sucesso!', 'x');
        this.getById();
      }
    })
  }
}
