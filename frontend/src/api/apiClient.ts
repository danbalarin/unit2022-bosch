import ky from 'ky';

import type { IUser } from './types/user';

export const apiClientRickAndMorty = ky.extend({
  prefixUrl: `https://rickandmortyapi.com/api/`,
  retry: 0,
});

export const apiClient = ky.extend({
  prefixUrl: `https://unit2022.herokuapp.com/api/`,
  retry: 0,
  hooks: {
    beforeRequest: [
      (request) => {
        const { token } = JSON.parse(
          localStorage.getItem('user') ?? '{}'
        ) as IUser;
        if (token) {
          request.headers.set('Authorization', `${token}`);
        }
      },
    ],
  },
});
