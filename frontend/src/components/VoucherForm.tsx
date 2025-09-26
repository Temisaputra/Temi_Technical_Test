"use client";

import { useForm, Controller } from "react-hook-form";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { useEffect } from "react";

interface VoucherFormProps {
  defaultValues: any;
  onSubmit: (data: any) => void;
  mode?: "create" | "edit";
}

export default function VoucherForm({ defaultValues, onSubmit, mode = "create" }: VoucherFormProps) {
  const { register, handleSubmit, reset, control } = useForm({
    defaultValues,
  });

  useEffect(() => {
    reset(defaultValues);
  }, [defaultValues, reset]);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      {/* Voucher Code */}
      <input
        placeholder="Voucher Code (auto-generated on create)"
        className="w-full p-2 border rounded bg-gray-100 text-black"
        disabled={mode === "edit"}
        {...register("voucher_code")}
      />

      {/* Discount */}
      <input
        type="number"
        placeholder="Discount (%)"
        className="w-full p-2 border rounded"
        {...register("discount", { required: true, min: 1, max: 100 ,valueAsNumber: true })}
      />

       {/* Expiry Date - pakai react-datepicker */}
      <Controller
        name="expired_date"
        control={control}
        rules={{ required: true }}
        render={({ field }) => (
          <DatePicker
            className="w-full p-2 border rounded"
            selected={field.value ? new Date(field.value) : null}
            onChange={(date) => field.onChange(date?.toISOString().split("T")[0])}
            dateFormat="yyyy-MM-dd"
          />
        )}
      />

      <button
        type="submit"
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        {mode === "edit" ? "Update" : "Create"}
      </button>
    </form>
  );
}
