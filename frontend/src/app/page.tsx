// src/app/page.tsx
"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  useEffect(() => {
    // redirect ke halaman login
    router.replace("/auth/login");
  }, [router]);

  return null; // kosongkan karena langsung redirect
}
