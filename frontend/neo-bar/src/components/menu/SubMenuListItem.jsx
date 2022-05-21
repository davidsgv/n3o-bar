import React from "react";

export function SubMenuListItem({menuSubItem}){
    const {text, icon, link} = menuSubItem
    
    return (
            <li className="full-width">
                <a href={link} className="full-width">
                    <div className="navLateral-body-cl">
                        <i className={"zmdi " +  icon}></i>
                    </div>
                    <div className="navLateral-body-cr">
                        {text}
                    </div>
                </a>
            </li>
    )
}