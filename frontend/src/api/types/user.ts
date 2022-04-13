export interface IUser {
  id: string;
  token: string;
  role: string;
}

export interface IUserLoginData {
  email: string;
  password: string;
}
