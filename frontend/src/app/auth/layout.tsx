export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="bg-gray-50 text-gray-900 min-h-screen flex items-center justify-center">
      {children}
    </div>
  );
}
