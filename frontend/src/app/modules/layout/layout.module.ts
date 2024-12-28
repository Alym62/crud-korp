import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { SharedModule } from '@shared/shared.module';
import { BaseComponent } from './base/base.component';
import { HeaderComponent } from './header/header.component';
import { ContentSidenavComponent } from './sidenav/content-sidenav/content-sidenav.component';
import { SidenavComponent } from './sidenav/sidenav.component';

@NgModule({
  declarations: [
    BaseComponent,
    HeaderComponent,
    SidenavComponent,
    ContentSidenavComponent
  ],
  imports: [
    CommonModule,
    SharedModule
  ],
  exports: [
    BaseComponent
  ]
})
export class LayoutModule { }
