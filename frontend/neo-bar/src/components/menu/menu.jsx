import React, {Fragment, useRef, useEffect} from "react";
import { NavBar } from "./NavBar";
import { TopBar } from "./TopBar";

const KEY = "login.token"

export function Menu({children}){

    const token = localStorage.getItem(KEY);

    useEffect(()=>{
        if(!token && window.location.href != "http://localhost:3000/"){
            window.location.href = "/"
        }
    },[]);

    const content = useRef();
    const nav = useRef();

    const toggleMenu = () =>{
        content.current.classList.toggle("pageContent-change")
        nav.current.classList.toggle("navLateral-change");
    }

    if(!token){
        return (
            <Fragment>
                {children}
            </Fragment> 
        )
    }

    return (
        <Fragment>
            <NavBar ref={nav} toggleMenu={toggleMenu}/>
            <section ref={content} className="full-width pageContent">
                <TopBar toggleMenu={toggleMenu}/>
                {children}
            </section>
        </Fragment>
    )
} 