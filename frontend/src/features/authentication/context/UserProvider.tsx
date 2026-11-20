import { useEffect, useState, type ReactNode } from "react";
import { create } from "mutative";
import { withRequest } from "@/lib/with-request";
import { getUser } from "../services/get-user";
import {
  UserContext,
  type AppState,
  type Result,
  type UserData,
} from "./user-context";

export function UserProvider({ children }: { children: ReactNode }) {
  const [appState, setAppState] = useState<AppState>({
    userData: { type: "idle" },
  });

  function changeUserData(userData: Result<UserData>) {
    setAppState((state) => {
      return create(state, (draft) => {
        draft.userData = userData;
      });
    });
  }

  useEffect(() => {
    withRequest(getUser, changeUserData);
  }, []);

  return (
    <UserContext.Provider
      value={{ state: appState, changeState: { changeUserData } }}
    >
      {children}
    </UserContext.Provider>
  );
}
