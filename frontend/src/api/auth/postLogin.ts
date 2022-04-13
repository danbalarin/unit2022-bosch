import type { UseMutationOptions } from 'react-query';
import { useMutation } from 'react-query';

import { apiClient } from '../apiClient';
import type { IRequestException } from '../types/requestException';
import type { IUser, IUserLoginData } from '../types/user';

export const loginUser = async (data: IUserLoginData) => {
  const response = await apiClient.post('login', {
    json: { ...data },
  });
  return (await response.json()) as IUser;
};

export const useLoginMutation = (
  options: UseMutationOptions<IUser, IRequestException, IUserLoginData>
) => {
  return useMutation<IUser, IRequestException, IUserLoginData>(
    async (data) => await loginUser(data),
    options
  );
};
