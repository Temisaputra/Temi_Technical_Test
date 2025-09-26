"use client";

import { useRouter } from "next/navigation";
import {  Props } from "@/types/voucher";

export default function VoucherTable({ vouchers, onSort, orderBy, orderType, onDelete }: Props) {
  const router = useRouter();

  const renderSortIcon = (field: string) => {
    if (orderBy !== field) return "↕";
    return orderType === "asc" ? "↑" : "↓";
  };

  return (
    <table className="w-full border">
      <thead className="bg-gray-100">
        <tr>
          <th className="p-2 cursor-pointer text-black" onClick={() => onSort("no")}>
            No {renderSortIcon("no")}
          </th>
          <th className="p-2 cursor-pointer text-black" onClick={() => onSort("voucher_code")}>
            Voucher Code {renderSortIcon("voucher_code")}
          </th>
          <th className="p-2 cursor-pointer text-black" onClick={() => onSort("discount")}>
            Discount (%) {renderSortIcon("discount")}
          </th>
          <th className="p-2 cursor-pointer text-black" onClick={() => onSort("expired_date")}>
            Expiry Date {renderSortIcon("expired_date")}
          </th>
          <th className="p-2 text-black">Actions</th>
        </tr>
      </thead>
      <tbody>
        {vouchers.map((v) => (
          <tr key={v.no} className="border-t">
            <td className="p-2 text-center">{v.no}</td>
            <td className="p-2">{v.voucher_code}</td>
            <td className="p-2 text-center">{v.discount}%</td>
            <td className="p-2">{v.expired_date || "-"}</td>
            <td className="p-2 text-center">
              <button
                onClick={() => router.push(`/vouchers/${v.no}`)}
                className="px-2 py-1 bg-yellow-500 text-white rounded hover:bg-yellow-600 mr-2"
              >
                Edit
              </button>
              <button
                onClick={() => onDelete(v.no)}
                className="px-2 py-1 bg-red-500 text-white rounded hover:bg-red-600"
              >
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}