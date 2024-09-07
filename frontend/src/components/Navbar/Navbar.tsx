'use client'

import Button from "../Button/Button";

type NavbarProps = {
    createRecipe: () => void
}

export default function Navbar({ createRecipe } : NavbarProps) {
    return (
        <div className='w-screen bg-primary items-center flex flex-row pt-[35px] pb-[35px] pr-[43px] pl-[43px] justify-between'>
            <h1 className='text-[35px] font-medium'>
                GENERATE COOK BOOK
            </h1>
            <Button
                onClick={createRecipe}
                color='secondary'
            >
                Create New
            </Button>
        </div>
    )
}
