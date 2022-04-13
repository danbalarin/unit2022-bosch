import type { IItem } from './item';

export interface IItemRequest {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
  Journey: null;
  JourneyID: number;
  WarehouseID: number;
  Item: IItem;
  ItemID: number;
  counts: number;
  RequestedBy: null;
  RequestedByID: number;
}

export enum JourneyStatus {
  LOADING = 'LOADING',
  ON_THE_WAY = 'ON_THE_WAY',
}

export interface IWarehouse {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
  Name: string;
}

export interface IWaypoint {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
  Route: null;
  RouteID: number;
  Warehouse: IWarehouse;
  WarehouseID: IWarehouse['ID'];
  Duration: number;
}

export interface IRoute {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
  Name: string;
  Waypoints: IWaypoint[];
  interval: number;
}

export interface IJourney {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
  Route: IRoute;
  RouteID: number;
  DepartureTime: string;
  Departed: boolean;
  ItemRequests: IItemRequest[];
  Place: number;
}
