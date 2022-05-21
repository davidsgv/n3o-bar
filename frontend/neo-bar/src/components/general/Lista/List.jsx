import React from "react";
import { ListItem } from "./ListItem";

export function List({list}){
    if(list.length > 0){
        return (
            <div className="mdl-list" id="categorias-container">
                {list.map((item,index) =>(<ListItem key={item.id} category={item}/>))}
            </div>
        )
    }
}