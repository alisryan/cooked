'use client'

import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import { Dispatch, SetStateAction, useState } from 'react';

type DropdownMenuProps = {
    items: string[];
    selectedItem: any;
    onSelect: Dispatch<SetStateAction<any>>
}

export default function DropdownMenu(props: DropdownMenuProps) {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <div className='relative'>
            <button
                onClick={() => setIsOpen(!isOpen)}
            >
                {props.selectedItem}
                <ArrowDropDownIcon />
            </button>   
            <div className={`absolute bottom-100 bg-primary rounded-md overflow-hidden ${isOpen ? '' : 'hidden'}`}>
                {props.items.map((item, index) => {
                    const isSelected = item === props.selectedItem

                    return (
                        <div
                            key={index} 
                            className={`cursor-pointer pr-4 pl-4 pt-2 pb-2 ${isSelected ? 'hover:!bg-secondaryHover' : 'hover:!bg-primaryHover'}`}
                            style={{
                                backgroundColor: isSelected ? '#393939' : '#d9d9d9',
                                color: isSelected ? 'white' : 'black'
                            }}
                            onClick={() => {
                                props.onSelect(item)
                                setIsOpen(false)
                            }}
                        >
                            {item}
                        </div>
                )})}
            </div>
        </div>
    )
}
