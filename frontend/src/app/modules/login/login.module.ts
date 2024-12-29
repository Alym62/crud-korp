import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { AuthService } from '@shared/services/auth.service';
import { SharedModule } from '@shared/shared.module';
import { LoginComponent } from './login.component';

@NgModule({
  declarations: [
    LoginComponent
  ],
  imports: [
    CommonModule,
    SharedModule
  ],
  providers: [
    AuthService,
  ]
})
export class LoginModule { }
