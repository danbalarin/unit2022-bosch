import ky from 'ky';

export const apiClientRickAndMorty = ky.extend({
  prefixUrl: `https://rickandmortyapi.com/api/`,
  retry: 0,
});

export const apiClient = ky.extend({
  prefixUrl: `https://unit2022.herokuapp.com/api/`,
  retry: 0,
  headers: { 'Content-Type': 'application/json' },
});
