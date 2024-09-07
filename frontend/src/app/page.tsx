
'use client'

import { fetchAllRecipes } from "@/api/recipe";
import Navbar from "@/components/Navbar/Navbar";
import CreateRecipe from "@/forms/CreateRecipe";
import { Modal } from "@mui/material";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";

export default function Home() {  
  const [isModalOpen, setIsModalOpen] = useState(false);

  const { data, isLoading, error } = useQuery({
    queryKey: ["recipes"],
    queryFn: fetchAllRecipes,
  })

  return (
    <>
      <Modal
        className='flex flex-col items-center justify-center'
        open={isModalOpen} 
        disableAutoFocus
        onClose={() => setIsModalOpen(false)}
      >
        <div className='bg-primary flex flex-col items-center p-16 w-[550px] h-[600px] overflow-scroll rounded-lg'>
          <CreateRecipe onSuccess={() => setIsModalOpen(false)} />
        </div>
      </Modal>
      <Navbar createRecipe={() => setIsModalOpen(true)} />
      <div className='flex flex-col p-16 items-start gap-6'>
        { /** TODO: THIS IS WHERE THE MAIN APP LIVES */}
      </div>
    </>
  );
}
