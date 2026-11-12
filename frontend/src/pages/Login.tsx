import { Button, Card } from "@chakra-ui/react";
import { Input, Center } from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";

function Login() {
  return (
    <Center mt={10}>
      <Card.Root width="500px">
        <Card.Body gap="2">
          <Card.Title mt="2">Login</Card.Title>
          <Card.Description>Username</Card.Description>
          <Input placeholder="Enter your username" />
          <Card.Description>Password</Card.Description>
          <PasswordInput placeholder="Enter your password" />
        </Card.Body>

        <Card.Footer justifyContent="flex-end">
          <Button>Login</Button>
        </Card.Footer>
      </Card.Root>
    </Center>
  );
}

export default Login;
