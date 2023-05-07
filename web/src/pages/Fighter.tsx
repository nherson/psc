import { Box, Heading, Text, VStack } from "@chakra-ui/react";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { client } from "../api/client";
import {
  Event as EventData,
  FightResult,
  Fighter,
  FighterResult,
} from "../api/psc_pb";
import { FightResultRow } from "../components/FightResultCard";
import { PageLoading } from "../components/PageLoading";
import { Link as RouteLink } from "react-router-dom";
import { FighterOverview } from "../components/FighterOverview";

const FighterPage = () => {
  const { id } = useParams();

  const fighterId = parseInt(id!);

  const [loading, setLoading] = useState(true);
  const [fighter, setFighter] = useState<Fighter>(new Fighter());
  const [fightResults, setFightResults] = useState<FightResult[]>();

  useEffect(() => {
    (async () => {
      setLoading(true);
      const resp = await client.listResultsForFighter({ fighterId: fighterId });
      setFighter(resp.fighter!);
      setFightResults(resp.fightResults);
      setLoading(false);
    })();
  }, [fighterId]);

  if (loading) {
    return <PageLoading />;
  }

  // Find this fighter's fight results to create an overview
  const fighterResults =
    fightResults?.map((f) => {
      // should never happen
      if (f.fighterResults.length != 2) {
        throw new Error("unexpected fighterResults for fight");
      }

      if (f.fighterResults[0].fighter?.id === id) {
        return f.fighterResults[0];
      } else if (f.fighterResults[1].fighter?.id === id) {
        return f.fighterResults[1];
      } else {
        // should never happen
        throw new Error("fighter not found in fight results");
      }
    }) || [];

  return (
    <Box>
      <Box pb="5">
        <Heading p={"3"}>
          {fighter.firstName} {fighter.lastName}
        </Heading>
        {fighter.nickName && <Text>{fighter.nickName}</Text>}
      </Box>

      <VStack>
        <Box pb={5}>
          <FighterOverview fighterResults={fighterResults} />
        </Box>
        {fightResults?.map((fr) => (
          <Box>
            <Box py="2">
              <RouteLink to={`/events/${fr.event!.id}`}>
                <Text>{fr.event?.name}</Text>
              </RouteLink>

              <Text fontSize="md">
                {fr.event!.date!.toDate().toLocaleDateString()}
              </Text>
            </Box>
            <Box py="2">
              <FightResultRow fightResult={fr} />
            </Box>
          </Box>
        ))}
      </VStack>
    </Box>
  );
};

export { FighterPage as Fighter };
