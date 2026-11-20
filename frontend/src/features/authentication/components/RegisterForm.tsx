import { useState, type FormEvent } from "react";
import { useNavigate } from "react-router";
import { Button, Card, Input } from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { toaster } from "@/components/ui/toaster";
import { registerUser } from "../services/auth-service";

export function RegisterForm() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const navigate = useNavigate();

  const handleRegister = async (e: FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      await registerUser({ username, email, password });
      toaster.create({
        title: "Register",
        description: "Register Berhasil",
        type: "success",
      });
      navigate("/login");
    } catch (err) {
      toaster.create({
        title: "Register",
        description: "Register Gagal",
        type: "error",
      });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Card.Root as="form" onSubmit={handleRegister} width="500px">
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
        <Button type="submit" loading={isLoading}>
          Register
        </Button>
      </Card.Footer>
    </Card.Root>
  );
}
