import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { PageResponse } from '@shared/commons/page.response';
import { Product } from '@shared/models/product.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  baseUrl: string = 'http://localhost:8080/api/v1/product';

  constructor(
    private http: HttpClient,
  ) { }

  pager(page: number, limit: number): Observable<{ data: PageResponse<Product>, success: boolean }> {
    const params: HttpParams = new HttpParams()
      .set("page", page.toString())
      .set("limit", limit.toString());

    return this.http.get<{ data: PageResponse<Product>, success: boolean }>(`${this.baseUrl}/pager`, { params });
  }

  getById(id: number): Observable<{ data: Product, success: boolean }> {
    return this.http.get<{ data: Product, success: boolean }>(`${this.baseUrl}/${id}`);
  }

  create(obj: Product): Observable<{ data: Product, success: boolean }> {
    return this.http.post<{ data: Product, success: boolean }>(`${this.baseUrl}/create`, obj);
  }

  update(id: number, obj: Product): Observable<{ data: Product, success: boolean }> {
    return this.http.put<{ data: Product, success: boolean }>(`${this.baseUrl}/update/${id}`, obj);
  }

  delete(id: number): Observable<{ data: Product, success: boolean }> {
    return this.http.delete<{ data: Product, success: boolean }>(`${this.baseUrl}/delete/${id}`);
  }
}
