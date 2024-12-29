import { inject } from "@angular/core";
import { ActivatedRouteSnapshot, CanActivateFn, Router, RouterStateSnapshot } from "@angular/router";
import { AuthService } from "@shared/services/auth.service";

export const AuthGuard: CanActivateFn = (route: ActivatedRouteSnapshot, state: RouterStateSnapshot) => {
  const router: Router = inject(Router);
  const authService: AuthService = inject(AuthService);

  const token: string | null = authService.getToken();

  if (token != null) return true;

  return router.navigate(['/login']);
}
