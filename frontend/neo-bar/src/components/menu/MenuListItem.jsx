import React, {useRef} from "react";
import {SubMenuList} from "./SubMenuList";

export function MenuListItem({menuItem}){
    const {text, icon, link, separator, submenus} = menuItem

    const subMenuLista = useRef();

    // var subMenuItem = <SubMenuList ref={subMenuLista} subMenu={submenus} onCLick={handleClick}/>

    if(separator){
        return (<li className="full-width divider-menu-h"></li>)
    }

    if(submenus){
        var icono = <span className="zmdi zmdi-chevron-left"></span>
    }else{
        var icono = ""
    }

    return (
        <li className="full-width">
            <a href={link} className="full-width btn-subMenu" onClick={() => subMenuLista.current.toogleMenu()}>
                <div className="navLateral-body-cl">
                    <i className={"zmdi " + icon}></i>
                </div>
                <div className="navLateral-body-cr">
                    {text}
                </div>
                {icono}
            </a>
            <SubMenuList ref={subMenuLista} subMenu={submenus}/>
            {/* {submenus ? subMenuItem: ""} */}
        </li>
    )
}