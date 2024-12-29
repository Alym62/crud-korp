import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { LoginDTO } from '@shared/models/dto/login.dto';
import { Observable, tap } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  baseUrl: string = 'http://localhost:8080/auth';

  constructor(
    private http: HttpClient,
  ) { }

  login(obj: LoginDTO): Observable<{ data: string, success: boolean }> {
    return this.http.post<{ data: string, success: boolean }>(`${this.baseUrl}/login`, obj)
      .pipe(
        tap((res) => this.saveTokenInLocalStorage(res.data)),
      );
  }

  getToken(): string | null {
    return localStorage.getItem('token');
  }

  logout(): void {
    localStorage.removeItem('token');
  }

  private saveTokenInLocalStorage(token: string): void {
    localStorage.setItem('token', token);
  }
}
