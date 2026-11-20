import { createClient } from "@connectrpc/connect";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";
import { transport } from "@/lib/connect-transport";

export const userServiceClient = createClient(UserService, transport);

export type LoginPayload = {
  username: string;
  password: string;
};

export type RegisterPayload = {
  username: string;
  email: string;
  password: string;
};

export async function loginUser(payload: LoginPayload) {
  return userServiceClient.login(payload);
}

export async function registerUser(payload: RegisterPayload) {
  return userServiceClient.register(payload);
}
