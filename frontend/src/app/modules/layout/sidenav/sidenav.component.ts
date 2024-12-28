import { Component, Input } from '@angular/core';

@Component({
  selector: 'sidenav-component',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss']
})
export class SidenavComponent {
  @Input('sidenavWidth') sidenavWidth!: string;

  get isCollapsed(): boolean {
    return this.sidenavWidth === '65px';
  }
}
