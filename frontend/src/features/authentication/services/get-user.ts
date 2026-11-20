import type { UserData } from "../context/user-context";
import { userServiceClient } from "../services/auth-service";

export async function getUser(): Promise<UserData> {
  return await userServiceClient.getUser({});
}
