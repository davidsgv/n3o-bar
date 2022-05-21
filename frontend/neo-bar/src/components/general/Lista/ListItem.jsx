import React from "react";

export function ListItem({category}){
    const {id, name, visible, tag, icon, actions} = category

    if(tag){
        var tagContent = <span className="mdl-list__item-sub-title">{tag}</span>
    }

    if(actions){
        <a className="mdl-list__item-secondary-action" href="#!"><i className="zmdi zmdi-more"></i></a>
    }

    if(visible){
        return (
            <div className="mdl-list__item mdl-list__item--two-line">
                <span className="mdl-list__item-primary-content">
                    <i className={"zmdi mdl-list__item-avatar " + icon}></i>
                    <span>{name}</span>
                    {tagContent}
                </span>
            </div>
        )
    }
}