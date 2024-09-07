'use client'

import { Space_Mono } from "next/font/google";
import "./globals.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const space_mono = Space_Mono({ weight: ["400", "700"], subsets: ["latin"] });

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 60,
      refetchOnWindowFocus: false,
    },
  }
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  return (
    <html lang="en">
      <QueryClientProvider client={queryClient}>
        <body className={space_mono.className}>
          {children}
        </body>
      </QueryClientProvider>
    </html>
  );
}
