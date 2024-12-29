import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DashboardModule } from '@modules/dashboard/dashboard.module';
import { LayoutModule } from '@modules/layout/layout.module';
import { LoginModule } from '@modules/login/login.module';
import { ProductModule } from '@modules/product/product.module';
import { JWTInterceptor } from '@shared/commons/jwt.interceptor';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    LayoutModule,
    DashboardModule,
    ProductModule,
    LoginModule,
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS, useClass: JWTInterceptor, multi: true,
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
