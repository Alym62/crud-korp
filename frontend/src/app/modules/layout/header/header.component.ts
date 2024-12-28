import { Component, EventEmitter, Output, signal } from '@angular/core';

@Component({
  selector: 'header-component',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {
  @Output() collapsedChange = new EventEmitter<boolean>();
  collapsed = signal(false);

  toggleCollapsed() {
    this.collapsed.set(!this.collapsed());
    this.collapsedChange.emit(this.collapsed());
  }
}
