import { createContext } from "react";

export type Result<T> =
  | { type: "idle" }
  | { type: "pending" }
  | { type: "success"; data: T }
  | { type: "error"; msg: string };

export type UserData = {
  username: string;
  email: string;
};

export type AppState = {
  userData: Result<UserData>;
};

export type ChangeUserData = (userData: Result<UserData>) => void;

export type IUserContext = {
  state: AppState;
  changeState: {
    changeUserData: ChangeUserData;
  };
};

export const UserContext = createContext<IUserContext | null>(null);
