export interface IItem {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt?: string | null;
  name: string;
}

export interface IOrderData {
  itemId: number;
  warehouseId: number;
  count: number;
}
