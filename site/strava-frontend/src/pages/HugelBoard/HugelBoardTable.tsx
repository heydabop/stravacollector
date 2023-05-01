import { FC, useEffect } from "react"
import {
  Grid,
  Table,
  Thead,
  Tbody,
  Tfoot,
  Tr,
  Th,
  Td,
  TableCaption,
  TableContainer,
  Box,
  Spinner,
  Text,
  Link,
  LinkBox,
  LinkOverlay,
  Skeleton,
  SkeletonCircle,
  SkeletonText,
  useToast,
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
  Tabs, TabList, TabPanels, Tab, TabPanel,
} from '@chakra-ui/react'
import { getErrorDetail, getErrorMessage, getHugelLeaderBoard } from "../../api/rest"
import { useMutation, useQuery } from "@tanstack/react-query";
import { AthleteAvatar } from "../../components/AthleteAvatar/AthleteAvatar";
import * as moment from 'moment'
import { HugelBoardProps } from "./HugelBoard";

export const HugelBoardTable: FC<HugelBoardProps> = ({
  data, error, isLoading, isFetched
}) => {

  return <>
    {
      error && <Alert status='error'>
        <AlertIcon />
        <AlertTitle>Failed to load leaderboard</AlertTitle>
        <AlertDescription>{getErrorMessage(error, "Leaderboard did not load.")}</AlertDescription>
      </Alert>
    }
    <TableContainer>
      <Table size="sm" variant='striped' colorScheme='gray'>
        <TableCaption>Das Hugel Leaderboad</TableCaption>
        <Thead>
          <Tr>
            <Th>Rank</Th>
            <Th>Athlete</Th>
            <Th>Elapsed Time</Th>
            {/* Add this later <Th>Difference</Th> */}
          </Tr>
        </Thead>
        {/* {tbody} */}
      </Table>
    </TableContainer>
  </>
  // let tbody = <TableLoading />
  // if (hugelLeaderboard) {
  //   tbody = <Tbody>{hugelLeaderboard.activities.map((activity, index) => {
  //     const elapsedDuration = moment.duration(activity.elapsed * 1000)

  //     // TODO: Idk why this hover does not override the striped columns
  //     return <LinkBox key={activity.athlete_id + index} as="tr" _hover={{ background: "yellow !important" }} >
  //       <Td>
  //         <LinkOverlay href={"https://www.strava.com/activities/" + activity.activity_id} target="_blank">
  //           <Box>
  //             <Text as="span">{activity.rank}</Text>
  //             <Text as="span">{activity.activity_name}</Text>
  //           </Box>
  //         </LinkOverlay>
  //       </Td>

  //       <Td>
  //         <Box display="flex" alignItems="center">
  //           <AthleteAvatar
  //             firstName={activity.athlete.firstname}
  //             lastName={activity.athlete.lastname}
  //             athleteID={activity.athlete_id}
  //             username={activity.athlete.username}
  //             profilePicLink={activity.athlete.profile_pic_link}
  //             size="sm"
  //           />

  //           <Text as="span" paddingLeft={"5px"}>
  //             {activity.athlete.firstname} {activity.athlete.lastname}
  //           </Text>
  //         </Box>

  //       </Td>

  //       <Td>
  //         {`${elapsedDuration.hours()}:${elapsedDuration.minutes()}:${elapsedDuration.seconds()}`}
  //         {/* {activity.elapsed} */}
  //       </Td>
  //     </LinkBox>
  //   })}</Tbody>
  // }

  // return <>
  //   {
  //     hugelLeaderboardError && <Alert status='error'>
  //       <AlertIcon />
  //       <AlertTitle>Failed to load leaderboard</AlertTitle>
  //       <AlertDescription>getErrorMessage(hugelLeaderboardError)</AlertDescription>
  //     </Alert>
  //   }
  //   <Box>
  //     <TableContainer>
  //       <Table size="sm" variant='striped' colorScheme='gray'>
  //         <TableCaption>Das Hugel Leaderboad</TableCaption>
  //         <Thead>
  //           <Tr>
  //             <Th>Rank</Th>
  //             <Th>Athlete</Th>
  //             <Th>Elapsed Time</Th>
  //             {/* Add this later <Th>Difference</Th> */}
  //           </Tr>
  //         </Thead>
  //         {tbody}
  //       </Table>
  //     </TableContainer>
  //   </Box >
  // </>
}

// const TableLoading: FC = () => {
//   const amount = [1, 1, 1]
//   const skeletonRows = amount.map((_, index) => {
//     return <Tr key={index}>
//       <Td> <SkeletonText skeletonHeight='4' noOfLines={1} width='20px' /></Td>
//       <Td>
//         <Box display="flex" alignItems="center">
//           <SkeletonCircle size='10' /><Skeleton width='20px' />
//           <SkeletonText skeletonHeight='4' noOfLines={1} width='100px' />
//         </Box>
//       </Td>
//       <Td> <SkeletonText skeletonHeight='4' noOfLines={1} width='80px' /></Td>
//     </Tr>
//   })
//   return <Tbody>{skeletonRows}</Tbody>
// }