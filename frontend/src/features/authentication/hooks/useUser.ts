import { useContext } from "react";
import { UserContext } from "../context/user-context";

export function useUser() {
  const ctx = useContext(UserContext);

  if (!ctx) {
    throw Error("useProfileContext must be called within UserProvider");
  }

  return ctx;
}
