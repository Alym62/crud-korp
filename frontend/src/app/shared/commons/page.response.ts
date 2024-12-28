export interface PageResponse<T> {
  list: Array<T>;
  total: number;
  page: number;
  limit: number;
  totalPages: number;
}
