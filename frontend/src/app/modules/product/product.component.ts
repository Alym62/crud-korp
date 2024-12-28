import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { PageResponse } from '@shared/commons/page.response';
import { Product } from '@shared/models/product.model';
import { ProductService } from '@shared/services/product.service';
import { ProductEditComponent } from './product-edit/product-edit.component';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.scss']
})
export class ProductComponent implements OnInit {

  page: number = 1;
  limit: number = 5;
  totalPages: number = 0;

  pager: PageResponse<Product> = {
    list: new Array(),
    total: 0,
    page: this.page,
    limit: this.limit,
    totalPages: 0,
  };

  constructor(
    private _matDialog: MatDialog,
    private service: ProductService,
    private _snackBar: MatSnackBar,
  ) { }

  ngOnInit(): void {
    this.getAll();
  }

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action, {
      horizontalPosition: 'end',
      verticalPosition: 'top',
      duration: 5000,
      panelClass: ['custom-snackbar'],
    });
  }

  getAll(): void {
    this.service.pager(this.page, this.limit).subscribe({
      next: (pager: { data: PageResponse<Product>, success: boolean }) => {
        this.pager = pager.data;
        this.totalPages = pager.data.totalPages;
      },
      error: (err) => this.openSnackBar('Ops! Não foi possível buscar os registros em nossa base de dados.', 'x'),
    });
  }

  add(): void {
    const dialogRef = this._matDialog.open(ProductEditComponent, {
      width: '600px',
      height: '500px'
    })

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.openSnackBar('Produto adicionado com sucesso!', 'x');
        this.getAll();
      }
    })
  }

  edit(item: Product): void {
    const dialogRef = this._matDialog.open(ProductEditComponent, {
      width: '600px',
      height: '500px',
      data: item,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.openSnackBar('Produto editado com sucesso!', 'x');
        this.getAll();
      }
    })
  }

  delete(id: number, event: Event): void {
    event.stopPropagation();

    this.service.delete(id).subscribe({
      next: (value) => {
        if (value.success) {
          this.openSnackBar('Produto removido com sucesso!', 'x');
          this.getAll();
        }
      },
      error: (err) => {
        this.openSnackBar('Ops! Não foi possível remover o produto!', 'x');
      },
    });
  }

  loadNextPage(): void {
    if (this.page < this.totalPages) {
      this.page++;
      this.getAll();
    }
  }

  loadPreviousPage(): void {
    if (this.page > 1) {
      this.page--;
      this.getAll();
    }
  }

  goToPage(page: number): void {
    if (page > 0 && page <= this.totalPages) {
      this.page = page;
      this.getAll();
    }
  }
}
