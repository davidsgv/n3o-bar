import React, {useEffect} from "react";
import {NavBar} from "../components/menu/NavBar"
import { TopBar } from "../components/TopBar";

const KEY = "login.token"

export function Home(){

    useEffect(()=>{
        const token = localStorage.getItem(KEY);
        if(!token){
            window.location.href = "/"
        }
    },[]);

    return (
        <div>
            <NavBar/>
                <section className="full-width pageContent">
                    <TopBar/>
                    <section className="full-width text-center" style={{padding: "40px 0"}}>
                        <h3 className="text-center tittles">RESPONSIVE TILES</h3>
                        <article className="full-width tile">
                            <div className="tile-text">
                                <span className="text-condensedLight">
                                    2<br/>
                                    <small>Administrators</small>
                                </span>
                            </div>
                            <i className="zmdi zmdi-account tile-icon"></i>
                        </article>
                        <article className="full-width tile">
                            <div className="tile-text">
                                <span className="text-condensedLight">
                                    71<br/>
                                    <small>Clients</small>
                                </span>
                            </div>
                            <i className="zmdi zmdi-accounts tile-icon"></i>
                        </article>
                        <article className="full-width tile">
                            <div className="tile-text">
                                <span className="text-condensedLight">
                                    7<br/>
                                    <small>Providers</small>
                                </span>
                            </div>
                            <i className="zmdi zmdi-truck tile-icon"></i>
                        </article>
                        <article className="full-width tile">
                            <div className="tile-text">
                                <span className="text-condensedLight">
                                    9<br/>
                                    <small>Categories</small>
                                </span>
                            </div>
                            <i className="zmdi zmdi-label tile-icon"></i>
                        </article>
                        <article className="full-width tile">
                            <div className="tile-text">
                                <span className="text-condensedLight">
                                    121<br/>
                                    <small>Products</small>
                                </span>
                            </div>
                            <i className="zmdi zmdi-washing-machine tile-icon"></i>
                        </article>
                        <article className="full-width tile">
                            <div className="tile-text">
                                <span className="text-condensedLight">
                                    47<br/>
                                    <small>Sales</small>
                                </span>
                            </div>
                            <i className="zmdi zmdi-shopping-cart tile-icon"></i>
                        </article>
                    </section>
                </section>
        </div>
    )
}