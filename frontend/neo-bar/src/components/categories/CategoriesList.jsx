import React from "react";
import { CategoriesListItem } from "./CategoriesListItem";

export function CategoriesList({categories}){
    if(categories.length > 0){
        return (
            <div className="mdl-list" id="categorias-container">
                {categories.map((cat,index) =>(<CategoriesListItem key={cat.id} category={cat}/>))}
            </div>
        )
    }
}