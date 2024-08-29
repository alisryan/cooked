'use client'

import { fetchAllRecipes } from "@/api/recipe";
import { useQuery } from "@tanstack/react-query";

export default function Home() {
  // EXAMPLE OF DATA FETCHING:
  // const { data, isLoading, error } = useQuery({
  //   queryKey: ["recipes"],
  //   queryFn: fetchAllRecipes,
  // })

  return (
    <main className="flex max-h-screen flex-col items-center justify-between">
      {/* TODO: BUILD THE MAIN PAGE OF THE APPLICATION HERE */}
    </main>
  );
}
