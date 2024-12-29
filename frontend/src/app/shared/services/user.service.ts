import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from '@shared/models/user.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  baseUrl: string = 'http://localhost:8080/api/v1/user';

  constructor(
    private http: HttpClient,
  ) { }

  getById(id: number): Observable<{ data: User, success: boolean }> {
    return this.http.get<{ data: User, success: boolean }>(`${this.baseUrl}/${id}`);
  }

  create(obj: User): Observable<{ data: User, success: boolean }> {
    return this.http.post<{ data: User, success: boolean }>(`${this.baseUrl}/create`, obj);
  }

  update(id: number, obj: User): Observable<{ data: User, success: boolean }> {
    return this.http.put<{ data: User, success: boolean }>(`${this.baseUrl}/update/${id}`, obj);
  }

  delete(id: number): Observable<{ data: User, success: boolean }> {
    return this.http.delete<{ data: User, success: boolean }>(`${this.baseUrl}/delete/${id}`);
  }
}
