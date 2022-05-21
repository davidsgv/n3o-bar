import React from "react";
import { getUserName } from "../../functions/user";

export function TopBar({toggleMenu}){

    const handleIconClik = ()=>{
        toggleMenu();
    }

    return (
        <div className="full-width navBar">
            <div className="full-width navBar-options">
                <i onClick={handleIconClik} className="zmdi zmdi-swap btn-menu" id="btn-menu"></i>	
                <div className="mdl-tooltip" htmlFor="btn-menu">Hide / Show MENU</div>
                <nav className="navBar-options-list">
                    <ul className="list-unstyle">
                        <li className="btn-Notification" id="notifications">
                            <i className="zmdi zmdi-notifications"></i>
                            <div className="mdl-tooltip" htmlFor="notifications">Notifications</div>
                        </li>
                        <li className="btn-exit" id="btn-exit">
                            <i className="zmdi zmdi-power"></i>
                            <div className="mdl-tooltip" htmlFor="btn-exit">LogOut</div>
                        </li>
                        <li className="text-condensedLight noLink" ><small>{getUserName()}</small></li>
                        <li className="noLink">
                            <figure>
                                <img src="/assets/img/avatar-male.png" alt="Avatar" className="img-responsive"/>
                            </figure>
                        </li>
                    </ul>
                </nav>
            </div>
        </div>
    )
}