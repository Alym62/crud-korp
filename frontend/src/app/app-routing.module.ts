import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from '@modules/dashboard/dashboard.component';
import { BaseComponent } from '@modules/layout/base/base.component';
import { LoginComponent } from '@modules/login/login.component';
import { ProductComponent } from '@modules/product/product.component';
import { AuthGuard } from '@shared/commons/auth.guard';

const routes: Routes = [
  {
    path: 'login', component: LoginComponent,
  },
  {
    canActivate: [AuthGuard],
    path: '', component: BaseComponent,
    children: [
      {
        path: 'dashboard', component: DashboardComponent,
      },
      {
        path: 'product', component: ProductComponent,
      },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
