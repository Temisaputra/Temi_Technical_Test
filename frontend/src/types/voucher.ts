export interface Voucher {
  no: number;
  voucher_code: string;
  discount: number;
  expired_date: string;
}

export interface Meta {
  total_data: number;
  total_page: number;
  page: number;
  page_size: number;
}

export interface Props {
  vouchers: Voucher[];
  onSort: (field: string) => void;
  orderBy: string;
  orderType: "asc" | "desc";
  onDelete: (id: number) => void; // callback dari parent
}