'use client'

import { ReactNode } from "react";

type ButtonProps = {
    onClick: () => void;
    fullWidth?: boolean,
    color?: 'primary' | 'secondary'
    children?: ReactNode;
    disabled?: boolean
}

// EXAMPLE COMPONENT:
export default function Button({ onClick, disabled, fullWidth, color = 'primary', children }: ButtonProps) {
    return (
        <button
            className={`bg-${color} ${fullWidth ? 'w-full' : ''} text-white rounded-md pt-3 pb-3 pr-5 pl-5 hover:bg-opacity-65 transition-all ease-in-out`}
            onClick={onClick}
            disabled={disabled}
        >
            {children}
        </button>
    );
}