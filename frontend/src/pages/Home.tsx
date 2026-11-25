import { Navbar } from "@/layouts/NavigationBar";
import { Box } from "@chakra-ui/react";
import { ChessBoard } from "@/features/catur";

function Home() {
  console.log("Home");

  return (
    <Box>
      <Navbar />
      <ChessBoard />
    </Box>
  );
}

export default Home;
