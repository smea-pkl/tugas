import { Button, Card } from "@chakra-ui/react";
import { Input, Center } from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { useState } from "react";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";
import { toaster } from "@/components/ui/toaster";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const userServiceClient = createClient(UserService, transport);

function Register() {
  const [username, setUsername] = useState();
  const [email, setEmail] = useState();
  const [password, setPassword] = useState();

  const register = () => {
    const payload = {
      username: username,
      email: email,
      password: password,
    };

    userServiceClient
      .register(payload)
      .then(() => {
        toaster.create({
          title: "Register",
          description: "Register Berhasil",
          type: "success",
        });
      })
      .catch(() => {
        toaster.create({
          title: "Register",
          description: "Register Gagal",
          type: "error",
        });
      });
  };

  return (
    <Center mt={10}>
      <Card.Root width="500px">
        <Card.Body gap="2">
          <Card.Title mt="2">Register</Card.Title>
          <Card.Description>Username</Card.Description>
          <Input
            placeholder="Enter your username"
            value={username}
            onChange={(e) => setUsername(e.target.value as any)}
          />
          <Card.Description>Email</Card.Description>
          <Input
            placeholder="Enter your email"
            value={email}
            onChange={(e) => setEmail(e.target.value as any)}
          />
          <Card.Description>Password</Card.Description>
          <PasswordInput
            placeholder="Enter your password"
            value={password}
            onChange={(e) => setPassword(e.target.value as any)}
          />
        </Card.Body>

        <Card.Footer justifyContent="flex-end">
          <Button onClick={register}>Register</Button>
        </Card.Footer>
      </Card.Root>
    </Center>
  );
}

export default Register;
