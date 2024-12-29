import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ProductService } from '@shared/services/product.service';
import { SharedModule } from '@shared/shared.module';
import { ProductEditComponent } from './product-edit/product-edit.component';
import { ProductComponent } from './product.component';

@NgModule({
  declarations: [
    ProductComponent,
    ProductEditComponent,
  ],
  imports: [
    CommonModule,
    SharedModule
  ],
  providers: [
    ProductService,
  ]
})
export class ProductModule { }
