import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { SharedModule } from '@shared/shared.module';
import { UserEditComponent } from './user-edit/user-edit.component';
import { UserComponent } from './user.component';

@NgModule({
  declarations: [
    UserComponent,
    UserEditComponent
  ],
  imports: [
    CommonModule,
    SharedModule
  ]
})
export class UserModule { }
