import "./globals.css";
import Navbar from "@/components/Navbar";
import { AuthProvider } from "@/context/AuthContext";

export const metadata = {
  title: "Voucher Management",
  description: "Manage discount vouchers easily",
};

import { ReactNode } from "react";

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-gray-50 text-gray-900">
        <div className="min-h-screen flex flex-col">
          {/* Wrapper untuk seluruh halaman */}
           <AuthProvider>

          <Navbar />
            <main className="flex-1 container mx-auto px-4 py-6">{children}</main>
          </AuthProvider>
        </div>
      </body>
    </html>
  );
}
