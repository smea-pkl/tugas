import { useState, type FormEvent } from "react";
import { useNavigate } from "react-router";
import { Button, Card, Input } from "@chakra-ui/react";
import { toaster } from "@/components/ui/toaster";
import { PasswordInput } from "@/components/ui/password-input";
import { useUser } from "../hooks/useUser";
import { loginUser } from "../services/auth-service";
import { withRequest } from "@/lib/with-request";
import { getUser } from "../services/get-user";

export function LoginForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const navigate = useNavigate();
  const { changeState } = useUser();

  const handleLogin = async (e: FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      const data = await loginUser({ username, password });
      localStorage.setItem("token", data.token);

      await withRequest(getUser, changeState.changeUserData);
      toaster.create({
        title: "Login",
        description: "Login Berhasil",
        type: "success",
      });
      navigate("/");
    } catch (err) {
      toaster.create({
        title: "Login",
        description: "Login Gagal",
        type: "error",
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Card.Root as="form" onSubmit={handleLogin} width="500px">
      <Card.Body gap="2">
        <Card.Title mt="2">Login</Card.Title>
        <Card.Description>Username</Card.Description>
        <Input
          placeholder="Enter your username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <Card.Description>Password</Card.Description>
        <PasswordInput
          placeholder="Enter your password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </Card.Body>

      <Card.Footer justifyContent="flex-end">
        <Button type="submit" loading={isLoading}>
          Login
        </Button>
      </Card.Footer>
    </Card.Root>
  );
}
