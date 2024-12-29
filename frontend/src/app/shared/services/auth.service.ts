import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { CurrentUserDTO } from '@shared/models/dto/current-user.dto';
import { LoginDTO } from '@shared/models/dto/login.dto';
import { Observable, tap } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  baseUrl: string = 'http://localhost:8080/auth';

  constructor(
    private http: HttpClient,
    private router: Router,
  ) { }

  login(obj: LoginDTO): Observable<{ data: { currentUser: CurrentUserDTO, token: string }, success: boolean }> {
    return this.http.post<{ data: { currentUser: CurrentUserDTO, token: string }, success: boolean }>(`${this.baseUrl}/login`, obj)
      .pipe(
        tap((res) => this.saveTokenInLocalStorage(res.data.token, res.data.currentUser)),
      );
  }

  getToken(): string | null {
    return localStorage.getItem('token');
  }

  logout(): void {
    localStorage.removeItem('token');
    localStorage.removeItem('currentUser');

    this.router.navigate(['/login']);
  }

  private saveTokenInLocalStorage(token: string, currentUser: object): void {
    localStorage.setItem('token', token);
    localStorage.setItem('currentUser', JSON.stringify(currentUser));
  }
}
