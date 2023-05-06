import { ChakraProvider, Box, Grid, theme } from "@chakra-ui/react";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import { Home } from "./pages/Home";
import NavBar from "./components/NavBar";
import { Fighters } from "./pages/Fighters";
import { Events } from "./pages/Events";
import { Event } from "./pages/Event";
import { Fighter } from "./pages/Fighter";

export const App = () => (
  <ChakraProvider theme={theme}>
    <BrowserRouter>
      <Box textAlign="center" fontSize="xl">
        <NavBar />
        <Grid minH="100vh" p={5} px={20}>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/fighters" element={<Fighters />} />
            <Route path="/events" element={<Events />} />
            <Route path="/events/:id" element={<Event />} />
            <Route path="/fighters/:id" element={<Fighter />} />
          </Routes>
        </Grid>
      </Box>
    </BrowserRouter>
  </ChakraProvider>
);
