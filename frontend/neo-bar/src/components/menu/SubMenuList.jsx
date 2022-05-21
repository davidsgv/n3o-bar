import React, {useRef,forwardRef,useImperativeHandle } from "react";
import { SubMenuListItem } from "./SubMenuListItem";


export const SubMenuList = forwardRef((props, ref) => {
    const menu = useRef()

    useImperativeHandle(ref, () => ({
        toogleMenu(){
            menu.current.classList.toggle("sub-menu-options-show");
        },
    }))

    if(props.subMenu){
        return (
            <ul ref={menu} className="full-width menu-principal sub-menu-options">
                {props.subMenu.map((subMenuItem,index) =>(<SubMenuListItem key={index} menuSubItem={subMenuItem}/>))}
            </ul>
        )
    }
});