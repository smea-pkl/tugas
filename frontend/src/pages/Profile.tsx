import { Card, Center } from "@chakra-ui/react";

function Profile() {
  return (
    <Center>
      <Card.Root>
        <Card.Body>
          <Card.Title>Profile</Card.Title>
          <Card.Description>Username : </Card.Description>
          <Card.Description>Email : </Card.Description>
        </Card.Body>
        <Card.Footer>
          <a href="/home">Back</a>
        </Card.Footer>
      </Card.Root>
    </Center>
  );
}

export default Profile;
