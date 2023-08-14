import { FC } from "react";
import { Flex, Text } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import { getAthlete, getRoute } from "../../api/rest";
import { useQuery } from "@tanstack/react-query";
import { NotFound } from "../404/404";
import { Loading } from "../../components/Loading/Loading";
import { AthleteOtherPage } from "./AthleteOtherPage";
import { useAuthenticated } from "../../contexts/Authenticated";
import { AthleteMePage } from "./AthleteMePage";

export const AthletePage: FC<{}> = ({}) => {
  const { athlete_id } = useParams();
  const { authenticatedUser, isLoading } = useAuthenticated();

  if (isLoading) {
    return <Loading />;
  }

  if (
    authenticatedUser?.athlete_id === athlete_id ||
    // Or Steven
    authenticatedUser?.athlete_id === "2661162"
  ) {
    return <AthleteMePage />;
  }
  return <AthleteOtherPage />;
};
