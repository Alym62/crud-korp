import { Component, computed, signal } from '@angular/core';

@Component({
  selector: 'base-component',
  templateUrl: './base.component.html',
  styleUrls: ['./base.component.scss']
})
export class BaseComponent {
  collapsed = signal(false);
  sidenavWidth = computed(() => (this.collapsed() ? '65px' : '250px'));

  toggleCollapsed(collapsed: boolean) {
    this.collapsed.set(collapsed);
  }
}
