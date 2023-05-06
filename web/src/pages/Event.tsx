import { Box, Heading, Text, VStack } from "@chakra-ui/react";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { client } from "../api/client";
import { Event as EventData, FightResult } from "../api/psc_pb";
import { FightResultRow } from "../components/FightResultCard";
import { PageLoading } from "../components/PageLoading";

export const Event = () => {
  const { id } = useParams();

  const eventId = parseInt(id!);

  const [loading, setLoading] = useState(true);
  const [event, setEvent] = useState<EventData>(new EventData());
  const [fightResults, setFightResults] = useState<FightResult[]>();

  useEffect(() => {
    (async () => {
      const resp = await client.listResultsForEvent({ eventId: eventId });
      setEvent(resp.event!);
      setFightResults(resp.fightResults);
      setLoading(false);
    })();
  }, [eventId]);

  if (loading) {
    return <PageLoading />;
  }

  return (
    <Box>
      <Box pb="5">
        <Heading p={"3"}>{event.name}</Heading>
        <Text>{event.date!.toDate().toLocaleDateString()}</Text>
      </Box>

      <VStack>
        {fightResults?.map((fr) => (
          <Box py="3">
            <FightResultRow fightResult={fr} />
          </Box>
        ))}
      </VStack>
    </Box>
  );
};
