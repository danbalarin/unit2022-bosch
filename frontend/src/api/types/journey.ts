import type { IItem } from './item';

interface ItemCount {
  count: number;
  itemId: IItem['ID'];
}

export enum JourneyStatus {
  LOADING = 'LOADING',
  ON_THE_WAY = 'ON_THE_WAY',
}

export interface IJourney {
  id: string;
  warehouseId: string;
  expeditionTime: string;
  items: ItemCount[];
}
