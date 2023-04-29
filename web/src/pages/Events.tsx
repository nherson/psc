import { Text } from "@chakra-ui/react";
import { PageLoading } from "../components/PageLoading";
import { client } from "../api/client";
import { useEffect, useState } from "react";
import { Event } from "../api/psc_pb";

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

  if (loading) {
    return <PageLoading />;
  }

  return <Text>Done</Text>;
};
