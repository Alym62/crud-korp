import { Component, EventEmitter, Output, signal } from '@angular/core';
import { AuthService } from '@shared/services/auth.service';

@Component({
  selector: 'header-component',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {
  @Output() collapsedChange = new EventEmitter<boolean>();
  collapsed = signal(false);

  constructor(
    private authService: AuthService,
  ) { }

  toggleCollapsed() {
    this.collapsed.set(!this.collapsed());
    this.collapsedChange.emit(this.collapsed());
  }

  logout(): void {
    this.authService.logout();
  }
}
