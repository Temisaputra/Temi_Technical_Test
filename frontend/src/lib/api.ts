import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8085/api", // ganti sesuai backend kamu
});

// Tambahkan interceptor untuk setiap request
api.interceptors.request.use((config) => {
  // Ambil token dari localStorage (atau cookies)
  if (typeof window !== "undefined") {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  }
  return config;
});

export default api;
