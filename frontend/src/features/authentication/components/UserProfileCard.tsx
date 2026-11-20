import { useNavigate } from "react-router";
import { useUser } from "../hooks/useUser";
import { Card, VStack, Text, Avatar, Button, Spinner } from "@chakra-ui/react";

export function UserProfileCard() {
  const { state } = useUser();
  const { userData } = state;
  const navigate = useNavigate();

  if (userData.type === "idle" || userData.type === "pending") {
    return <Spinner size="lg" color="purple.400" />;
  }

  if (userData.type === "error") {
    return (
      <VStack gap={3}>
        <Text color="red.400">Gagal memuat profil atau belum login.</Text>
        <Button onClick={() => navigate("/login")}>Ke Halaman Login</Button>
      </VStack>
    );
  }

  return (
    <Card.Root w="full" maxW="md" shadow="md" rounded="md" bg="#424c55">
      <Card.Body>
        <VStack gap={4} textAlign="center">
          <Avatar.Root size="2xl" colorPalette="purple">
            <Avatar.Fallback />
          </Avatar.Root>
          <VStack gap={1}>
            <Text fontSize="xl" fontWeight="bold" color="#f5edf0">
              {userData.data.username}
            </Text>
            <Text fontSize="md" color="#d1ccdc">
              {userData.data.email}
            </Text>
          </VStack>
        </VStack>
      </Card.Body>
      <Card.Footer justifyContent="center">
        <Button
          onClick={() => navigate("/")}
          w="full"
          variant="outline"
          borderColor="#d1ccdc"
          color="#f5edf0"
          _hover={{ bg: "#d1ccdc", color: "#424c55" }}
        >
          Back
        </Button>
      </Card.Footer>
    </Card.Root>
  );
}
