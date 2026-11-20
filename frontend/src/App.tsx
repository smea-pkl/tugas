import { StrictMode } from "react";
import { BrowserRouter, Routes, Route } from "react-router";
import { Provider } from "@/components/ui/provider";
import { Toaster } from "@/components/ui/toaster";
import { UserProvider } from "@/features/authentication";
import Register from "./pages/Register.tsx";
import Login from "./pages/Login.tsx";
import Home from "./pages/Home.tsx";
import Profile from "./pages/Profile.tsx";

export default function App() {
  return (
    <StrictMode>
      <UserProvider>
        <Provider defaultTheme="light">
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Home />}></Route>
              <Route path="/login" element={<Login />}></Route>
              <Route path="/register" element={<Register />}></Route>
              <Route path="/profile" element={<Profile />}></Route>
            </Routes>
          </BrowserRouter>
          <Toaster />
        </Provider>
      </UserProvider>
    </StrictMode>
  );
}
