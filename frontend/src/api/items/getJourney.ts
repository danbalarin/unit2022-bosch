import { useQuery } from 'react-query';
import type { UseQueryOptions } from 'react-query';

import { apiClient } from '../apiClient';
import type { IRequestException } from '../types/requestException';
import type { IJourney } from '../types/journey';

interface IGetJourneysResponse {
  journeys: IJourney[];
}

export const getJourneys = async () => {
  const response = await apiClient.get('journeys', {});
  return (await response.json()) as IGetJourneysResponse;
};

export const useGetJourneysQuery = (
  options: UseQueryOptions<IGetJourneysResponse, IRequestException>
) => {
  return useQuery<IGetJourneysResponse, IRequestException>(
    'journeys',
    async () => await getJourneys(),
    options
  );
};
