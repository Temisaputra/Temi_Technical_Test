"use client";

import { useEffect, useState, useRef } from "react";
import { useRouter } from "next/navigation";
import api from "@/lib/api";
import { Voucher, Meta } from "@/types/voucher";
import VoucherTable from "@/components/VoucherTable";
import Papa from "papaparse";

export default function VoucherListPage() {
  const router = useRouter();
  const [vouchers, setVouchers] = useState<Voucher[]>([]);
  const [meta, setMeta] = useState<Meta | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);


  // query state
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [keyword, setKeyword] = useState("");
  const [orderBy, setOrderBy] = useState("no");
  const [orderType, setOrderType] = useState<"asc" | "desc">("asc");

  const fetchVouchers = async () => {
    try {
      const res = await api.get("/vouchers", {
        params: {
          page,
          page_size: pageSize,
          keyword,
          order_by: orderBy,
          order_type: orderType,
        },
      });

      setVouchers(res.data.data || []);
      setMeta(res.data.meta || null);
    } catch (err) {
      console.error("Failed to fetch vouchers", err);
    }
    };
    
const handleDelete = async (id: number) => {
  if (!confirm("Are you sure you want to delete this voucher?")) return;
  try {
    await api.delete(`/voucher-delete/${id}`);
    fetchVouchers(); // refresh data setelah delete
  } catch (err: any) {
    console.error("Error deleting voucher:", err.response?.data || err.message);
    alert("Failed to delete voucher");
  }
  };
  
 // 📤 Export CSV (langsung download dari backend)
  const handleExport = async () => {
    try {
      const res = await api.get("/vouchers/export-csv", {
        responseType: "blob", // supaya dapat file binary
      });

      const url = window.URL.createObjectURL(new Blob([res.data]));
      const link = document.createElement("a");
      link.href = url;
      link.setAttribute("download", "vouchers.csv");
      document.body.appendChild(link);
      link.click();
      link.remove();
    } catch (err) {
      console.error("Export failed:", err);
      alert("Export gagal");
    }
  };

  // 📥 Import CSV (pakai multipart/form-data)
 const handleImport = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    const formData = new FormData();
    formData.append("file", file);

    try {
      const token = localStorage.getItem("token");
      await api.post("/vouchers/upload-csv", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
          Authorization: `Bearer ${token}`,
        },
      });

      alert("Import berhasil");
      await fetchVouchers(); // ⬅️ tunggu refresh data selesai
    } catch (err: any) {
      console.error("Import failed:", err.response?.data || err.message);
      alert("Import gagal: " + (err.response?.data?.message || err.message));
    } finally {
      if (fileInputRef.current) {
        fileInputRef.current.value = ""; // ⬅️ reset supaya bisa upload file sama lagi
      }
    }
  };
  useEffect(() => {
    fetchVouchers();
  }, [page, pageSize, keyword, orderBy, orderType]);

  return (
      <div className="space-y-4">
        <div className="flex justify-between items-end">
          <h1 className="text-2xl font-bold">Voucher List</h1>
         <div className="flex gap-2">
          <button
            onClick={() => router.push("/vouchers/create")}
            className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
          >
            Create Voucher
          </button>
          <button
            onClick={handleExport}
            className="px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700"
          >
            Export CSV
          </button>
          <label className="px-4 py-2 bg-yellow-600 text-white rounded hover:bg-yellow-700 cursor-pointer">
            Import CSV
            <input
              ref={fileInputRef}
              type="file"
              accept=".csv"
              onChange={handleImport}
              className="hidden"
            />
          </label>
        </div>
        </div>
      {/* Search box */}
      <input
        placeholder="Search by code..."
        value={keyword}
        onChange={(e) => setKeyword(e.target.value)}
        className="border p-2 rounded w-full"
      />

      {/* Table */}
      <VoucherTable
    vouchers={vouchers}
    onSort={(field) => {
    if (orderBy === field) {
      // toggle asc/desc
      setOrderType(orderType === "asc" ? "desc" : "asc");
    } else {
      setOrderBy(field);
      setOrderType("asc"); // default asc saat ganti kolom
    }
  }}
    orderBy={orderBy}
    orderType={orderType}
    onDelete={handleDelete}   // 👉 pasang callback di sini
  />

    {/* Pagination + Page Size */}
      {meta && (
        <div className="flex items-center justify-between mt-4">
          <div className="flex items-center gap-4">
            <span>
              Page {meta.page} of {meta.total_page} | Total {meta.total_data}
            </span>
            <select
              value={pageSize}
              onChange={(e) => {
                setPageSize(Number(e.target.value));
                setPage(1); // reset ke page 1 saat ganti page size
              }}
              className="border p-1 rounded text-white bg-black"
            >
              <option value={5}>5 / page</option>
              <option value={10}>10 / page</option>
              <option value={20}>20 / page</option>
              <option value={50}>50 / page</option>
            </select>
          </div>
          <div className="space-x-2">
            <button
              disabled={page <= 1}
              onClick={() => setPage(page - 1)}
              className="px-3 py-1 border rounded disabled:opacity-50"
            >
              Prev
            </button>
            <button
              disabled={page >= meta.total_page}
              onClick={() => setPage(page + 1)}
              className="px-3 py-1 border rounded disabled:opacity-50"
            >
              Next
            </button>
          </div>
        </div>
      )}
    </div>
  );
}