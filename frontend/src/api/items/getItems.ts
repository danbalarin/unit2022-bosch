import { useQuery } from 'react-query';
import type { UseQueryOptions } from 'react-query';

import { apiClient } from '../apiClient';
import type { IItem } from '../types/item';
import type { IRequestException } from '../types/requestException';

interface IGetItemsResponse {
  items: IItem[];
}

export const getItems = async () => {
  const response = await apiClient.get('items', {});
  return (await response.json()) as IGetItemsResponse;
};

export const useGetItemsQuery = (
  options: UseQueryOptions<IGetItemsResponse, IRequestException>
) => {
  return useQuery<IGetItemsResponse, IRequestException>(
    'items',
    async () => await getItems(),
    options
  );
};
