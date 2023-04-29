import { Center, Spinner } from "@chakra-ui/react";

export const PageLoading = () => (
  <Center h="50vh">
    <Spinner speed="1s" size="xl"></Spinner>
  </Center>
);
