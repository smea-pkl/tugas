import { Button, Card } from "@chakra-ui/react";
import { Input, Center } from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { useState } from "react";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";
import { toaster } from "@/components/ui/toaster";
import { useNavigate } from "react-router";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  interceptors: [
    function (next) {
      return function (req) {
        const token = localStorage.getItem("token");

        if (!token) {
          if (req.header.has("Authorization")) {
            req.header.delete("Authorization");
          }

          return next(req);
        }

        req.header.set("Authorization", `Bearer ${token}`);

        return next(req);
      };
    },
  ],
});

export const userServiceClient = createClient(UserService, transport);

function Login() {
  const [username, setUsername] = useState();
  const [password, setPassword] = useState();
  const navigate = useNavigate();

  const login = async () => {
    const payload = {
      username: username,
      password: password,
    };

    try {
      const data = await userServiceClient.login(payload);
      localStorage.setItem("token", data.token);
      toaster.create({
        title: "Login",
        description: "Login Berhasil",
        type: "success",
      });
      navigate("/home");
    } catch (error) {
      toaster.create({
        title: "Login",
        description: "Login Gagal",
        type: "error",
      });
    }
  };

  return (
    <Center mt={10}>
      <Card.Root width="500px">
        <Card.Body gap="2">
          <Card.Title mt="2">Login</Card.Title>
          <Card.Description>Username</Card.Description>
          <Input
            placeholder="Enter your username"
            value={username}
            onChange={(e) => setUsername(e.target.value as any)}
          />
          <Card.Description>Password</Card.Description>
          <PasswordInput
            placeholder="Enter your password"
            value={password}
            onChange={(e) => setPassword(e.target.value as any)}
          />
        </Card.Body>

        <Card.Footer justifyContent="flex-end">
          <Button onClick={login}>Login</Button>
        </Card.Footer>
      </Card.Root>
    </Center>
  );
}

export default Login;
