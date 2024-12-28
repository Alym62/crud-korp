import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ProductEditComponent } from './product-edit/product-edit.component';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.scss']
})
export class ProductComponent {
  constructor(
    private _matDialog: MatDialog,
  ) { }

  add(): void {
    this._matDialog.open(ProductEditComponent)
  }
}
