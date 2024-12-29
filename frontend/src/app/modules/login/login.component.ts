import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
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
}
