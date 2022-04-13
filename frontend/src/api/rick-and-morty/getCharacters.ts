import type { UseQueryOptions, QueryClient } from 'react-query';
import { useQuery } from 'react-query';

import { apiClientRickAndMorty } from '../apiClient';

import type { Character } from './character';
import type { RickAndMortyResponse } from './response';

type IGetCharactersProps = UseQueryOptions<RickAndMortyResponse<Character[]>>;

interface IClientProps {
  client: QueryClient;
}

const fetchCharacters = async () => {
  const response = await apiClientRickAndMorty.get('character', {});
  return (await response.json()) as RickAndMortyResponse<Character[]>;
};

export const useGetCharactersQuery = ({ ...options }: IGetCharactersProps) => {
  return useQuery<RickAndMortyResponse<Character[]>>(
    'characters',
    async () => await fetchCharacters(),
    {
      ...options,
    }
  );
};

export const getCharactersQueryWithClient = async ({
  client,
  ...options
}: IGetCharactersProps & IClientProps) => {
  return await client.prefetchQuery<RickAndMortyResponse<Character[]>>(
    'characters',
    async () => await fetchCharacters(),
    {
      ...options,
    }
  );
};
