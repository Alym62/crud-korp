import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { UserEditComponent } from '@modules/user/user-edit/user-edit.component';
import { AuthService } from '@shared/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  formBuilder: FormBuilder = new FormBuilder();
  formGroup!: FormGroup;

  constructor(
    private authService: AuthService,
    private router: Router,
    private _snackBar: MatSnackBar,
    private _matDialog: MatDialog,
  ) {
    this.createForm();
  }

  private createForm(): void {
    this.formGroup = this.formBuilder.group({
      email: [null, Validators.required],
      password: [null, Validators.required],
    });
  }

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action, {
      horizontalPosition: 'end',
      verticalPosition: 'top',
      duration: 5000,
      panelClass: ['custom-snackbar'],
    });
  }

  login(): void {
    if (!this.formGroup.valid) {
      return;
    }

    this.authService.login(this.formGroup.value).subscribe({
      next: (value) => {
        if (value.success) {
          this.router.navigate(['/dashboard']);
          this.openSnackBar('Bem-vindo ao nosso sistema!', 'x');
        }
      },
      error: (err) => {
        this.openSnackBar(err.error.error, 'x');
      }
    });
  }

  register(): void {
    const dialogRef = this._matDialog.open(UserEditComponent, {
      width: '600px',
      height: '600px'
    })

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.openSnackBar('Parabéns! Você acaba de ser registrado no nosso sistema.', 'x');
      }
    })
  }
}
