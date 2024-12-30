import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { User } from '@shared/models/user.model';
import { AuthService } from '@shared/services/auth.service';
import { UserService } from '@shared/services/user.service';

@Component({
  selector: 'app-user-edit',
  templateUrl: './user-edit.component.html',
  styleUrls: ['./user-edit.component.scss']
})
export class UserEditComponent implements OnInit {
  isNew: boolean = false;

  formBuilder: FormBuilder = new FormBuilder();
  formGroup!: FormGroup;

  isManagerSelected: boolean = false;
  isSellerSelected: boolean = false;

  user!: User;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: User,
    private service: UserService,
    private _dialogRef: MatDialogRef<UserEditComponent>,
    private router: Router,
    private authService: AuthService,
  ) {
    this.createForm();
  }

  ngOnInit(): void {
    this.isNew = !!this.data ? false : true;
    if (!this.isNew) {
      this.user = this.data;
      this.formGroup.patchValue({ ...this.data, role: this.data.role });
    }
  }

  private createForm(): void {
    this.formGroup = this.formBuilder.group({
      email: [null, Validators.required],
      password: [null, Validators.required],
      position: [null],
      roleManager: [null],
      roleSeller: [null],
    });
    this.formGroup.valueChanges.subscribe(value => {
      Object.assign(this.user, value);
    });
  }

  onRoleChange(selectedRole: string): void {
    if (selectedRole === 'manager') {
      this.isManagerSelected = true;
      this.isSellerSelected = false;
      this.formGroup.get('roleManager')?.setValue(selectedRole);
    } else if (selectedRole === 'seller') {
      this.isManagerSelected = false;
      this.isSellerSelected = true;
      this.formGroup.get('roleSeller')?.setValue(selectedRole);
    }
  }

  saveOrUpdate(): void {
    if (!this.formGroup.valid) {
      this.formGroup.markAllAsTouched();
      return;
    }

    let role = this.isManagerSelected ? this.formGroup.get('roleManager')?.value : this.formGroup.get('roleSeller')?.value;

    const userSaveOrUpdate = { ...this.formGroup.value, role: role }
    if (this.isNew) {
      this.service.create(userSaveOrUpdate).subscribe({
        next: (value) => {
          this._dialogRef.close(value.success);
        },
        error: (err) => console.error(err.error),
      });
    } else {
      this.service.update(this.user.id, userSaveOrUpdate).subscribe({
        next: (value) => {
          this._dialogRef.close(value.success);
          this.authService.logout();
          this.router.navigate(['/login']);
        },
        error: (err) => console.error(err.error),
      });
    }
  }
}
