import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { SharedModule } from '@shared/shared.module';
import { ProductComponent } from './product.component';
import { ProductEditComponent } from './product-edit/product-edit.component';

@NgModule({
  declarations: [
    ProductComponent,
    ProductEditComponent
  ],
  imports: [
    CommonModule,
    SharedModule
  ],
})
export class ProductModule { }
