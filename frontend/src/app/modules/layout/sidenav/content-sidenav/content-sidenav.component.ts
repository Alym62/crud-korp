import { Component, computed, Input, signal, Signal } from '@angular/core';

@Component({
  selector: 'content-sidenav-component',
  templateUrl: './content-sidenav.component.html',
  styleUrls: ['./content-sidenav.component.scss']
})
export class ContentSidenavComponent {
  sidenavCollapsed = signal(false);
  @Input('collapsed') set collapsed(val: boolean) {
    this.sidenavCollapsed.set(val);
  }

  menuItem: Signal<Array<{ icon: string, label: string, router?: string }>> = signal(new Array(
    {
      icon: "dashboard",
      label: "Dashboard",
      router: "/dashboard",
    },
    {
      icon: "assessment",
      label: "Produtos",
      router: "/product",
    },
  ));

  profilePicSize = computed(() => this.sidenavCollapsed() ? '32' : '100');
}
