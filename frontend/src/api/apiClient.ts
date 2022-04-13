import ky from 'ky';

export const apiClientRickAndMorty = ky.extend({
  prefixUrl: `https://rickandmortyapi.com/api/`,
  retry: 0,
});

function authHeader() {
  /* return authorization header with jwt token */
  let token = localStorage.getItem('token');

  if (token) {
    return { Authorization: token };
  } else {
    return {};
  }
}

export const apiClient = ky.extend({
  prefixUrl: `https://unit2022.herokuapp.com/api/`,
  retry: 0,
  headers: authHeader(),
});
