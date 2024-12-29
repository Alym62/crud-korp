export interface User {
  id: number;
  email: string;
  password: string;
  position: string;
  role: string;
  createdAt: Date;
  updatedAt: Date;
  removed: boolean;
}
