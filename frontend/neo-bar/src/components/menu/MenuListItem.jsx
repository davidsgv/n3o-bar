import React from "react";

export function MenuListItem({menuItem}){
    const {text, icon, link, separator} = menuItem

    if(separator){
        return (<li className="full-width divider-menu-h"></li>)
    }

    return (
        <li className="full-width">
            <a href={link} className="full-width">
                <div className="navLateral-body-cl">
                    <i className={"zmdi " + icon}></i>
                </div>
                <div className="navLateral-body-cr">
                    {text}
                </div>
            </a>
        </li>
    )
}