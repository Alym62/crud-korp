import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Product } from '@shared/models/product.model';
import { ProductService } from '@shared/services/product.service';

@Component({
  selector: 'app-product-edit',
  templateUrl: './product-edit.component.html',
  styleUrls: ['./product-edit.component.scss']
})
export class ProductEditComponent implements OnInit {
  isNew: boolean = false;

  formBuilder: FormBuilder = new FormBuilder();
  formGroup!: FormGroup;

  product!: Product;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: Product,
    private service: ProductService,
    private _dialogRef: MatDialogRef<ProductEditComponent>,
    private _snackBar: MatSnackBar,
  ) {
    this.createForm();
  }

  ngOnInit(): void {
    this.isNew = !!this.data ? false : true;
    if (!this.isNew) {
      this.product = this.data;
      this.formGroup.patchValue(this.data);
    }
  }

  private createForm(): void {
    this.formGroup = this.formBuilder.group({
      name: [null, Validators.required],
      description: [null, Validators.required],
      price: [0, Validators.required],
    });
    this.formGroup.valueChanges.subscribe(value => {
      Object.assign(this.product, value);
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

  saveOrUpdate(): void {
    if (!this.formGroup.valid) {
      this.formGroup.markAllAsTouched();
      return;
    }

    if (this.isNew) {
      this.service.create(this.formGroup.value).subscribe({
        next: (value) => {
          this._dialogRef.close(value.success);
        },
        error: (err) => this.openSnackBar(err.error.error, 'x'),
      });
    } else {
      this.service.update(this.product.id, this.formGroup.value).subscribe({
        next: (value) => {
          this._dialogRef.close(value.success);
        },
        error: (err) => this.openSnackBar(err.error.error, 'x'),
      });
    }
  }
}
