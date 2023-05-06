import {
  Box,
  Card,
  CardBody,
  Heading,
  Stack,
  Text,
  VStack,
  useColorModeValue,
} from "@chakra-ui/react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { client } from "../api/client";
import { PageLoading } from "../components/PageLoading";
import { Fighter } from "../api/psc_pb";

export const Fighters = () => {
  const [loading, setLoading] = useState(true);
  const [fighters, setEvents] = useState(new Array<Fighter>());

  useEffect(() => {
    (async () => {
      const resp = await client.listFighters({});
      setEvents(resp.fighters);
      setLoading(false);
    })();
  }, []);

  const navigate = useNavigate();

  const cardHoverColors = useColorModeValue("blue.200", "blue.700");

  if (loading) {
    return <PageLoading />;
  }

  const sortedFighters = fighters.sort((f1, f2) => {
    if (f1.lastName.toLowerCase() < f2.lastName.toLowerCase()) {
      return -1;
    } else if (f2.lastName.toLowerCase() < f1.lastName.toLowerCase()) {
      return 1;
    }
    return 0;
  });

  return (
    <Box>
      <VStack>
        {sortedFighters.map((f) => (
          <Card
            w="lg"
            _hover={{
              cursor: "pointer",
              textDecoration: "none",
              bg: cardHoverColors,
            }}
            onClick={() => navigate(`/fighters/${f.id}`)}
          >
            <CardBody>
              <Stack mt="3" spacing="1">
                <Heading size="md">
                  {f.firstName} {f.lastName}
                </Heading>
                {f.nickName !== "" && <Text fontSize="lg">{f.nickName}</Text>}
              </Stack>
            </CardBody>
          </Card>
        ))}
      </VStack>
    </Box>
  );
};
