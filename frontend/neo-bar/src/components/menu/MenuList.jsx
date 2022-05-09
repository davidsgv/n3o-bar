import React from "react";
import { MenuListItem } from "./MenuListItem";

export function MenuList({menu}){
    if(menu.length > 0){
        return (
            <nav className="full-width">
                <ul className="full-width list-unstyle menu-principal">
                    {menu.map((menuItem,index) =>(<MenuListItem key={index} menuItem={menuItem}/>))}
                </ul>
            </nav>
        )
    }
}