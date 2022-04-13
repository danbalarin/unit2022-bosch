export interface RickAndMortyResponseInfo {
  count: number;
  pages: number;
  next: string;
  prev: string;
}

export interface RickAndMortyResponse<TData> {
  info: RickAndMortyResponseInfo;
  results: TData;
}
