import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from '@modules/dashboard/dashboard.component';
import { BaseComponent } from '@modules/layout/base/base.component';
import { ProductComponent } from '@modules/product/product.component';

const routes: Routes = [
  {
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
