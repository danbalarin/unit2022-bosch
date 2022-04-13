export interface IRequestException {
  type: string;
  message: string;
  response: {
    status: number;
  };
}
