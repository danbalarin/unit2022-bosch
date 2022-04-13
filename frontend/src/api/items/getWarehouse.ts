import { useQuery } from 'react-query';
import type { UseQueryOptions } from 'react-query';

import { apiClient } from '../apiClient';
import type { IRequestException } from '../types/requestException';
import type { IWarehouse } from '../types/journey';

interface IGetWarehousesResponse {
  warehouses: IWarehouse[];
}

export const getWarehouses = async () => {
  const response = await apiClient.get('warehouses', {});
  return (await response.json()) as IGetWarehousesResponse;
};

export const useGetWarehousesQuery = (
  options: UseQueryOptions<IGetWarehousesResponse, IRequestException>
) => {
  return useQuery<IGetWarehousesResponse, IRequestException>(
    'warehouses',
    async () => await getWarehouses(),
    options
  );
};
