import { useNavigate } from "react-router";
import {
  Box,
  Flex,
  Text,
  Button,
  HStack,
  Avatar,
  Menu,
  Portal,
  Spinner,
} from "@chakra-ui/react";
import { useUser } from "@/features/authentication";

export function Navbar() {
  const { state, changeState } = useUser();
  const { userData } = state;
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("token");
    changeState.changeUserData({ type: "idle" });
    navigate("/login");
  };

  return (
    <Box bg="#3d2c2e" px={5} py={3}>
      <Flex align="center" justify="space-between">
        <Text color="white" fontSize="xl" fontWeight="bold">
          PKL
        </Text>

        <HStack>
          <Button onClick={() => navigate("/")} bg="#886f68" color="#f5edf0">
            Home
          </Button>
          {userData.type === "success" && (
            <Button
              onClick={() => navigate("/profile")}
              bg="#886f68"
              color="#f5edf0"
            >
              Profile
            </Button>
          )}
        </HStack>

        <Box>
          {userData.type === "pending" && <Spinner size="sm" color="#f5edf0" />}

          {(userData.type === "idle" || userData.type === "error") && (
            <HStack gap={2}>
              <Button
                size="sm"
                variant="outline"
                color="#f5edf0"
                borderColor="#f5edf0"
                _hover={{ bg: "#886f68" }}
                onClick={() => navigate("/login")}
              >
                Login
              </Button>
              <Button
                size="sm"
                variant="outline"
                color="#f5edf0"
                borderColor="#f5edf0"
                _hover={{ bg: "#886f68" }}
                onClick={() => navigate("/register")}
              >
                Register
              </Button>
            </HStack>
          )}

          {userData.type === "success" && (
            <Menu.Root>
              <Menu.Trigger>
                <Box>
                  <HStack>
                    <Avatar.Root size="sm">
                      <Avatar.Fallback name={userData.data.username} />
                    </Avatar.Root>
                  </HStack>
                </Box>
              </Menu.Trigger>

              <Portal>
                <Menu.Positioner>
                  <Menu.Content>
                    <Menu.Item
                      value="user"
                      onClick={() => navigate("/profile")}
                    >
                      <Avatar.Root>
                        <Avatar.Fallback />
                      </Avatar.Root>
                      <Text>{userData.data.username}</Text>
                    </Menu.Item>
                    <Menu.Item
                      value="logout"
                      color="red.500"
                      onClick={handleLogout}
                    >
                      Logout
                    </Menu.Item>
                  </Menu.Content>
                </Menu.Positioner>
              </Portal>
            </Menu.Root>
          )}
        </Box>
      </Flex>
    </Box>
  );
}
