import { useQuery } from 'react-query';
import type { UseQueryOptions } from 'react-query';

import { apiClient } from '../apiClient';
import type { IItem } from '../types/item';
import type { IRequestException } from '../types/requestException';

export const getItems = async () => {
  const response = await apiClient.get('items', {});
  return (await response.json()) as IItem;
};

export const useGetItemsQuery = (
  options: UseQueryOptions<IItem, IRequestException>
) => {
  return useQuery<IItem, IRequestException>(
    'items',
    async () => await getItems(),
    options
  );
};
