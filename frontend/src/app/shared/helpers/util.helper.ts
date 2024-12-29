import { CurrentUserDTO } from "@shared/models/dto/current-user.dto";

export class UtilHelper {
  private static instance: UtilHelper;

  public static get getInstance() {
    if (UtilHelper.instance) {
      return UtilHelper.instance;
    }
    return this.instance;
  };

  public static getCurrentUser(): CurrentUserDTO | null {
    const currentUser = localStorage.getItem('currentUser');
    if (!currentUser) {
      return null;
    }

    try {
      return JSON.parse(currentUser) as CurrentUserDTO;
    } catch (error) {
      return null;
    }
  }
}
