import { Center } from "@chakra-ui/react";
import { UserProfileCard } from "@/features/authentication";

function Profile() {
  return (
    <Center minH="100vh">
      <UserProfileCard />
    </Center>
  );
}

export default Profile;
