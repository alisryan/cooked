import { ReactNode } from "react";

type TagProps = {
    children: ReactNode;
}

export default function Tag({ children }: TagProps) {
    return (
        <div className='bg-secondary text-white font-medium text-[12px] rounded-full pl-8 pr-8 pt-3 pb-3'>
            { children }
        </div>
    )
}
