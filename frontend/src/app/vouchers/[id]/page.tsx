"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import api from "@/lib/api";
import { Voucher } from "@/types/voucher";
import VoucherForm from "@/components/VoucherForm";

export default function EditVoucherPage() {
  const router = useRouter();
  const params = useParams();
  const id = params.id;

  const [voucher, setVoucher] = useState<Partial<Voucher> | null>(null);

  useEffect(() => {
    const fetchVoucher = async () => {
      try {
        const res = await api.get(`/voucher/${id}`);
        setVoucher(res.data.data); // ⬅️ pastikan response "data" sesuai
      } catch (err) {
        console.error("Error fetching voucher:", err);
      }
    };
    fetchVoucher();
  }, [id]);

  const handleUpdate = async (data: Voucher) => {
    try {
      await api.put(`/voucher-update/${id}`, data);
      router.push("/vouchers");
    } catch (err) {
      console.error("Error updating voucher:", err);
    }
  };

  if (!voucher) {
    return <p className="p-6">Loading...</p>;
  }

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Edit Voucher</h1>
      <VoucherForm defaultValues={voucher} onSubmit={handleUpdate} mode="edit" />
    </div>
  );
}
