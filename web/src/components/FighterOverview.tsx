import {
  Box,
  Card,
  CardBody,
  CardHeader,
  Center,
  Divider,
  Flex,
  Heading,
  Text,
  VStack,
  useColorModeValue,
} from "@chakra-ui/react";
import { FightResult, Fighter, FighterResult } from "../api/psc_pb";

interface AverageFightResults {
  significantStrikes: number;
  takedowns: number;
  knockdowns: number;
  controlTimeSeconds: number;
  score: number;
}

export const FighterOverview = ({
  fighterResults,
}: {
  fighterResults: FighterResult[];
}) => {
  let avgFightResults = fighterResults.reduce<AverageFightResults>(
    (avg, fr) => ({
      significantStrikes: avg.significantStrikes + fr.significantStrikes,
      takedowns: avg.takedowns + fr.takedowns,
      knockdowns: avg.knockdowns + fr.knockdowns,
      score: avg.score + fr.score,
      controlTimeSeconds: avg.controlTimeSeconds + fr.controlTimeSeconds,
    }),
    {
      significantStrikes: 0,
      takedowns: 0,
      knockdowns: 0,
      score: 0,
      controlTimeSeconds: 0,
    }
  );
  avgFightResults = {
    significantStrikes:
      avgFightResults.significantStrikes / fighterResults.length,
    takedowns: avgFightResults.takedowns / fighterResults.length,
    knockdowns: avgFightResults.knockdowns / fighterResults.length,
    score: avgFightResults.score / fighterResults.length,
    controlTimeSeconds:
      avgFightResults.controlTimeSeconds / fighterResults.length,
  };

  return (
    <Card bgColor={useColorModeValue("gray.100", "gray.700")}>
      <CardHeader>
        <Box>
          <Heading size="md">Overview</Heading>
        </Box>
      </CardHeader>
      <Divider />
      <CardBody>
        <Center w="100%">
          <VStack>
            <Flex w="100%" px="3">
              <Box>
                <Text fontWeight={"bold"} fontSize="md">
                  Significant Strikes
                </Text>
              </Box>
              <Box flexGrow={1} w="50px"></Box>
              <Box>
                <Text fontSize="md">
                  {avgFightResults.significantStrikes.toFixed(2)}
                </Text>
              </Box>
            </Flex>
            <Flex w="100%" px="3">
              <Box>
                <Text fontWeight={"bold"} fontSize="md">
                  Takedowns
                </Text>
              </Box>
              <Box flexGrow={1} w="50px"></Box>
              <Box>
                <Text fontSize="md">
                  {avgFightResults.takedowns.toFixed(2)}
                </Text>
              </Box>
            </Flex>
            <Flex w="100%" px="3">
              <Box>
                <Text fontWeight={"bold"} fontSize="md">
                  Knockdowns
                </Text>
              </Box>
              <Box flexGrow={1} w="50px"></Box>
              <Box>
                <Text fontSize="md">
                  {avgFightResults.knockdowns.toFixed(2)}
                </Text>
              </Box>
            </Flex>
            <Flex w="100%" px="3">
              <Box>
                <Text fontWeight={"bold"} fontSize="md">
                  Control Time
                </Text>
              </Box>
              <Box flexGrow={1} w="50px"></Box>
              <Box>
                <Text fontSize="md">
                  {formatSeconds(avgFightResults.controlTimeSeconds)}
                </Text>
              </Box>
            </Flex>
          </VStack>
        </Center>
      </CardBody>
      <Divider />
      <CardBody>
        <VStack>
          <Flex w="100%" px="3">
            <Box>
              <Text fontWeight={"bold"} fontSize="md">
                Average Score
              </Text>
            </Box>
            <Box flexGrow={1} w="50px"></Box>
            <Box>
              <Text fontSize="md">{avgFightResults.score.toFixed(2)}</Text>
            </Box>
          </Flex>
          <Flex w="100%" px="3">
            <Box>
              <Text fontWeight={"bold"} fontSize="md">
                Fight Count
              </Text>
            </Box>
            <Box flexGrow={1} w="50px"></Box>
            <Box>
              <Text fontSize="md">{fighterResults.length}</Text>
            </Box>
          </Flex>
        </VStack>
      </CardBody>
    </Card>
  );
};

const formatSeconds = (n: number): string => {
  const m = Math.floor(n / 60);
  const s = n % 60;
  const ss = s < 10 ? `0${s}` : `${s}`; // pad the 0 if needed
  return `${m}:${ss}`;
};
