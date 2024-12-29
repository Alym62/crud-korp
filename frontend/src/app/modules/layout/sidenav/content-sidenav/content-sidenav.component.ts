import { Component, computed, Input, signal, Signal } from '@angular/core';
import { UtilHelper } from '@shared/helpers/util.helper';

@Component({
  selector: 'content-sidenav-component',
  templateUrl: './content-sidenav.component.html',
  styleUrls: ['./content-sidenav.component.scss']
})
export class ContentSidenavComponent {
  currentUser = UtilHelper.getCurrentUser();

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
