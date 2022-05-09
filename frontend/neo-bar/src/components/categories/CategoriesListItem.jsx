import React from "react";

export function CategoriesListItem({category}){
    const {id, nombre, visible} = category

    if(visible){
        return (
            <div className="mdl-list__item mdl-list__item--two-line">
                <span className="mdl-list__item-primary-content">{id}. {nombre}
                    <i className="zmdi zmdi-label mdl-list__item-avatar"></i>
                </span>
            </div>
        )
    }
}