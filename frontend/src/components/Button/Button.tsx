'use client'

import { ReactNode } from "react";

type ButtonColor = 'primary' | 'secondary';

type ButtonProps = {
    onClick?: () => void
    fullWidth?: boolean
    color?: ButtonColor
    children?: ReactNode
    disabled?: boolean
    type?: 'submit' | 'button'
}

const computeButtonColor = (color: ButtonColor) => {
    switch (color) {
        case 'primary':
            return '#D9D9D9';
        case 'secondary':
            return '#393939';
        default:
            return '#D9D9D9';
    }
}

// EXAMPLE COMPONENT:
export default function Button({ onClick, disabled, fullWidth, color = 'primary', children, type = 'button' }: ButtonProps) {
    return (
        <button
            className={`${fullWidth ? 'w-full' : ''} text-white rounded-md pt-3 pb-3 pr-5 pl-5 hover:bg-opacity-65 transition-all ease-in-out`}
            onClick={onClick}
            disabled={disabled}
            style={{
                backgroundColor: computeButtonColor(color),
            }}
            type={type}
        >
            {children}
        </button>
    );
}
