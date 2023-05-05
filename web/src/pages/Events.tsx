import { Link as RouteLink } from "react-router-dom";
import { PageLoading } from "../components/PageLoading";
import { client } from "../api/client";
import { useEffect, useState } from "react";
import { Event } from "../api/psc_pb";
import {
  Box,
  Center,
  Text,
  Stack,
  List,
  ListItem,
  ListIcon,
  Button,
  useColorModeValue,
  VStack,
  Card,
  ButtonGroup,
  CardBody,
  CardFooter,
  Divider,
  Heading,
  Flex,
  Spacer,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

export const Events = () => {
  const [loading, setLoading] = useState(true);
  const [events, setEvents] = useState(new Array<Event>());

  useEffect(() => {
    (async () => {
      const resp = await client.listEvents({});
      setEvents(resp.events);
      setLoading(false);
    })();
  }, []);

  const navigate = useNavigate();

  const cardHoverColors = useColorModeValue("blue.200", "blue.700");

  if (loading) {
    return <PageLoading />;
  }

  const sortedEvents = events.sort((e1, e2) => {
    if (!e1.date) {
      return 1;
    } else if (!e2.date) {
      return -1;
    } else {
      return Number(e2.date.seconds) - Number(e1.date.seconds);
    }
  });

  return (
    <Box>
      <VStack>
        {sortedEvents.map((e) => (
          <Card
            w="lg"
            _hover={{
              cursor: "pointer",
              textDecoration: "none",
              bg: cardHoverColors,
            }}
            onClick={() => navigate(`/events/${e.id}`)}
          >
            <CardBody>
              <Stack mt="3" spacing="1">
                <Heading size="md">{e.name}</Heading>
                <Text fontSize="lg">
                  {e.date!.toDate().toLocaleDateString()}
                </Text>
              </Stack>
            </CardBody>
          </Card>
        ))}
      </VStack>
    </Box>
  );
};
