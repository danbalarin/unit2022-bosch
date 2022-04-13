import {
  Center,
  Heading,
  Stat,
  StatHelpText,
  StatLabel,
  StatNumber,
} from '@chakra-ui/react';
import React, { useEffect } from 'react';
import { dehydrate, QueryClient } from 'react-query';

import {
  getCharactersQueryWithClient,
  useGetCharactersQuery,
} from '~/api/rick-and-morty/getCharacters';
import { NavbarLayout } from '~/components/NavbarLayout';

export async function getStaticProps() {
  const queryClient = new QueryClient();

  await getCharactersQueryWithClient({ client: queryClient });

  return {
    props: {
      dehydratedState: dehydrate(queryClient),
    },
  };
}

export function RickAndMortyPage() {
  const { data } = useGetCharactersQuery({});
  useEffect(() => {
    if (data?.results) {
      console.table(data.results);
    }
  }, [data]);
  return (
    <NavbarLayout>
      <Center
        display="flex"
        flexDirection="column"
        alignItems="center"
        justifyContent="center"
        mt={8}
      >
        <Heading>Rick and Morty API</Heading>
        <Stat mt={4}>
          <StatLabel>Characters</StatLabel>
          <StatNumber>{data?.info.count}</StatNumber>
          <StatHelpText>Check console for full list</StatHelpText>
        </Stat>
      </Center>
    </NavbarLayout>
  );
}
