import ky from 'ky';

export const apiClient = ky.extend({
  prefixUrl: `https://rickandmortyapi.com/api/`,
  retry: 0,
});
