"use client";

import VoucherForm from "@/components/VoucherForm";
import api from "@/lib/api";
import { useRouter } from "next/navigation";

export default function CreateVoucherPage() {
  const router = useRouter();

  const handleCreate = async (data: any) => {
    await api.post("/voucher-create", data);
    router.push("/vouchers");
  };

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Create Voucher</h1>
      <VoucherForm defaultValues={{}} onSubmit={handleCreate} mode="create" />
    </div>
  );
}
