import React, { forwardRef, useRef } from "react";
import { getUserName, getUserRol } from "../../functions/user";
import { MenuList } from "./MenuList";

export const NavBar = forwardRef((props, ref)=>{
    const menu = [
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
            icon: "zmdi-case",
            submenus: [
                {
                    text: "USUARIOS",
                    link: "/admin",
                    icon: "zmdi-face"
                },
                // {
                //     text: "COMPANY",
                //     link: "#!",
                //     icon: "zmdi-balance",
                // },
                {
                    text: "PROVEEDORES",
                    link: "/notDo",
                    icon: "zmdi-truck"
                },
                {
                    text: "FORMAS DE PAGO",
                    link: "/notDo",
                    icon: "zmdi-card"
                },
                {
                    text: "CATEGORÍAS",
                    link: "/categories",
                    icon: "zmdi-label"
                },
                {
                    text: "PRODUCTOS",
                    link: "/notDo",
                    icon: "zmdi-washing-machine"
                },
                {
                    text: "INVENTORY",
                    link: "/inventory",
                    icon: "zmdi-store"
                },
                {
                    text: "MESAS",
                    link: "/notDo",
                    icon: "zmdi-drink"
                },
                {
                    text: "CAJAS",
                    link: "/notDo",
                    icon: "zmdi-money"
                }
            ]
        },
        {
            separator: true,
        },
        {
            text: "SALES (informe)",
            link: "/notDo",
            icon: "zmdi-shopping-cart"
        },
        {
            separator: true,
        },
        {
            text: "PEDIDOS (mesero)",
            link: "/notDo",
            icon: "zmdi-coffee"
        },
        {
            separator: true,
        },
        {
            text: "CAJA (cajero)",
            link: "/notDo",
            icon: "zmdi-money-box"
        }
    ];

    const handleCloseClick = ()=>{
        props.toggleMenu();
    }

    return (
        <section ref={ref} className="full-width navLateral" style={{position: "fixed", left:0, top:0}}>
            <div className="full-width navLateral-bg btn-menu"></div>
            <div className="full-width navLateral-body">
                <div className="full-width navLateral-body-logo text-center tittles">
                    <i onClick={handleCloseClick} className="zmdi zmdi-close btn-menu"></i> Inventory 
                </div>
                <figure className="full-width navLateral-body-tittle-menu">
                    <div>
                        <img src="/assets/img/avatar-male.png" alt="Avatar" className="img-responsive"/>
                    </div>
                    <figcaption>
                        <span id="nombreusuario" >
                            {getUserName()}<br/>
                            <small id="rolusuario">{getUserRol()}</small>
                        </span>
                    </figcaption>
                </figure>

                <MenuList menu={menu} />
            </div>
        </section>
    )
});