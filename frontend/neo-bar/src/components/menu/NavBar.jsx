import React, {useState} from "react";
import { getUserName, getUserRol } from "../../functions/user";
import { MenuList } from "./MenuList";

export function NavBar(){
    const [menu, setMenu] = useState([
        {
            text: "HOME",
            link: "/home",
            icon: "zmdi-view-dashboard"
        },
        {
            separator: true,
        },
        {
            text: "ADMINISTRACION",
            link: "#!",
            icon: "zmdi-case"
        },
        {
            separator: true,
        },
        {
            text: "USUARIOS",
            link: "/admin",
            icon: "zmdi-face"
        },
        {
            separator: true,
        },
        {
            text: "PRODUCTS",
            link: "#!",
            icon: "zmdi-washing-machine"
        },
        {
            separator: true,
        },
        {
            text: "SALES",
            link: "#!",
            icon: "zmdi-shopping-cart"
        },
        {
            separator: true,
        },
        {
            text: "INVENTORY",
            link: "/inventory",
            icon: "zmdi-store"
        },
        {
            separator: true,
        },
        {
            text: "CATEGORIES",
            link: "/categories",
            icon: "zmdi-label"
        }
    ]);

    return (
        <section className="full-width navLateral" style={{position: "fixed", left:0, top:0}}>
            <div className="full-width navLateral-bg btn-menu"></div>
            <div className="full-width navLateral-body">
                <div className="full-width navLateral-body-logo text-center tittles">
                    <i className="zmdi zmdi-close btn-menu"></i> Inventory 
                </div>
                <figure className="full-width navLateral-body-tittle-menu">
                    <div>
                        <img src="assets/img/avatar-male.png" alt="Avatar" className="img-responsive"/>
                    </div>
                    <figcaption>
                        <span id="nombreusuario" >
                            {getUserName()}<br/>
                            <small id="rolusuario">{getUserRol()}</small>
                        </span>
                    </figcaption>
                </figure>

                <MenuList menu={menu} />
                        {/* <li className="full-width">
                            <a href="#!" className="full-width btn-subMenu">
                                <div className="navLateral-body-cl">
                                    <i className="zmdi zmdi-case"></i>
                                </div>
                                <div className="navLateral-body-cr">
                                    ADMINISTRACION
                                </div>
                                <span className="zmdi zmdi-chevron-left"></span>
                            </a>
                            <ul className="full-width menu-principal sub-menu-options">
                                <li className="full-width">
                                    <a href="company.html" className="full-width">
                                        <div className="navLateral-body-cl">
                                            <i className="zmdi zmdi-balance"></i>
                                        </div>
                                        <div className="navLateral-body-cr">
                                            COMPANY
                                        </div>
                                    </a>
                                </li>
                                <li className="full-width">
                                    <a href="providers.html" className="full-width">
                                        <div className="navLateral-body-cl">
                                            <i className="zmdi zmdi-truck"></i>
                                        </div>
                                        <div className="navLateral-body-cr">
                                            PROVIDERS
                                        </div>
                                    </a>
                                </li>
                                <li className="full-width">
                                    <a href="payments.html" className="full-width">
                                        <div className="navLateral-body-cl">
                                            <i className="zmdi zmdi-card"></i>
                                        </div>
                                        <div className="navLateral-body-cr">
                                            PAYMENTS
                                        </div>
                                    </a>
                                </li>
                                <li className="full-width">
                                    <a href="/categories" className="full-width">
                                        <div className="navLateral-body-cl">
                                            <i className="zmdi zmdi-label"></i>
                                        </div>
                                        <div className="navLateral-body-cr">
                                            CATEGORIAS
                                        </div>
                                    </a>
                                </li>
                            </ul>
                        </li> */}
                        {/* <li className="full-width">
                            <a href="#!" className="full-width btn-subMenu">
                                <div className="navLateral-body-cl">
                                    <i className="zmdi zmdi-face"></i>
                                </div>
                                <div className="navLateral-body-cr">
                                    USUARIOS
                                </div>
                                <span className="zmdi zmdi-chevron-left"></span>
                            </a>
                            <ul className="full-width menu-principal sub-menu-options">
                                <li className="full-width">
                                    <a href="admin.html" className="full-width">
                                        <div className="navLateral-body-cl">
                                            <i className="zmdi zmdi-account"></i>
                                        </div>
                                        <div className="navLateral-body-cr">
                                            ADMINISTRADORES
                                        </div>
                                    </a>
                                </li>
                            </ul>
                        </li> */}
                        {/* <li className="full-width">
                            <a href="#!" className="full-width btn-subMenu">
                                <div className="navLateral-body-cl">
                                    <i className="zmdi zmdi-wrench"></i>
                                </div>
                                <div className="navLateral-body-cr">
                                    SETTINGS
                                </div>
                                <span className="zmdi zmdi-chevron-left"></span>
                            </a>
                            <ul className="full-width menu-principal sub-menu-options">
                                <li className="full-width">
                                    <a href="#!" className="full-width">
                                        <div className="navLateral-body-cl">
                                            <i className="zmdi zmdi-widgets"></i>
                                        </div>
                                        <div className="navLateral-body-cr">
                                            OPTION
                                        </div>
                                    </a>
                                </li>
                                <li className="full-width">
                                    <a href="#!" className="full-width">
                                        <div className="navLateral-body-cl">
                                            <i className="zmdi zmdi-widgets"></i>
                                        </div>
                                        <div className="navLateral-body-cr">
                                            OPTION
                                        </div>
                                    </a>
                                </li>
                            </ul>
                        </li> */}
            </div>
        </section>
    )
}