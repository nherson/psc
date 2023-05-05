import {
  Box,
  Card,
  CardBody,
  CardHeader,
  Center,
  Divider,
  Flex,
  HStack,
  Heading,
  Text,
  VStack,
  useColorModeValue,
} from "@chakra-ui/react";
import { FightResult, FighterResult } from "../api/psc_pb";

export const FightResultRow = ({
  fightResult,
}: {
  fightResult: FightResult;
}) => {
  return (
    <HStack spacing="1" align="stretch">
      <Flex grow={1}>
        <FighterResultCard
          fighterResult={fightResult.fighterResults[0]}
          fightResult={fightResult}
        />
      </Flex>
      <Flex p="4">
        <Center>vs</Center>
      </Flex>
      <Flex grow={1}>
        <FighterResultCard
          fighterResult={fightResult.fighterResults[1]}
          fightResult={fightResult}
        />
      </Flex>
    </HStack>
  );
};

const FighterResultCard = ({
  fighterResult,
  fightResult,
}: {
  fighterResult: FighterResult;
  fightResult: FightResult;
}) => (
  <Card w="100%" bgColor={useColorModeValue("gray.100", "gray.700")}>
    <Flex
      h="5px"
      w="100%"
      bgColor={fighterResult.corner === "red" ? "red.700" : "blue.700"}
    />
    <CardHeader>
      <Box>
        <Heading size="md">
          {fighterResult.fighter?.firstName} {fighterResult.fighter?.lastName}
        </Heading>
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
              <Text fontSize="md">{fighterResult.significantStrikes}</Text>
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
              <Text fontSize="md">{fighterResult.takedowns}</Text>
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
              <Text fontSize="md">{fighterResult.knockdowns}</Text>
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
                {formatSeconds(fighterResult.controlTimeSeconds)}
              </Text>
            </Box>
          </Flex>
        </VStack>
      </Center>
    </CardBody>
    <Divider />
    <CardBody>
      <Flex w="100%" px="3">
        <Box>
          <Text fontWeight={"bold"} fontSize="md">
            Score
          </Text>
        </Box>
        <Box flexGrow={1} w="50px"></Box>
        <Box>
          <Text fontSize="md">{fighterResult.score.toFixed(2)}</Text>
        </Box>
      </Flex>
      {fighterResult.win && (
        <Flex w="100%" pt="3">
          <Box flexGrow={1}>
            <Center>
              <VStack>
                <Text fontSize="md">Winner</Text>
                <Text fontSize="md">{fighterResult.winMethod}</Text>
                {fighterResult.winByStoppage && (
                  <Text fontSize="sm">
                    R{fightResult.resultEndingRound}{" "}
                    {formatSeconds(fightResult.resultEndingTimeSeconds)}
                  </Text>
                )}
              </VStack>
            </Center>
          </Box>
        </Flex>
      )}
    </CardBody>
  </Card>
);

const formatSeconds = (n: number): string => {
  const m = Math.floor(n / 60);
  const s = n % 60;
  const ss = s < 10 ? `0${s}` : `${s}`; // pad the 0 if needed
  return `${m}:${ss}`;
};
