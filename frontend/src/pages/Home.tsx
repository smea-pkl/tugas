import { useEffect, useState } from "react";
import { userServiceClient } from "./Login";
import { Box, Flex, Text, Link, Menu, Button, Portal } from "@chakra-ui/react";

function Home() {
  const [username, setUsername] = useState("");

  const getProfile = async () => {
    const profile = await userServiceClient.getProfile({});
    setUsername(profile.username);
  };

  useEffect(() => {
    getProfile();
  }, []);

  return (
    <Box backgroundColor="gray">
      <Flex justifyContent="space-between" mx={10}>
        <Text textStyle="4xl" fontWeight="bold">
          PKL
        </Text>

        <Link href="/home">Home</Link>

        <Menu.Root>
          <Menu.Trigger>
            <Button variant="outline" size="sm">
              {username}
            </Button>
          </Menu.Trigger>

          <Portal>
            <Menu.Positioner>
              <Menu.Content>
                <a href="">Profile</a>
              </Menu.Content>
              <Menu.Content>Logout</Menu.Content>
            </Menu.Positioner>
          </Portal>
        </Menu.Root>
      </Flex>
    </Box>
  );
}

export default Home;
