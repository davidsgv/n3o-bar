import React, {useEffect,useRef, useState} from "react";
import { getServerUrl } from "../functions/configuration";

const KEY = "login.token"
const URL = getServerUrl() + "/api/login/";


export function Login(){
    const [token, setToken] = useState([""]); 

    const EmailInputRef =  useRef();
    const PasswordRef =  useRef();
    const LoginButtonRef = useRef();

    useEffect(()=>{
        const token = localStorage.getItem(KEY);
        if(token){
            window.location.href = "/home"
        }
    },[]);

    const handleLogin = (e) =>{
        e.preventDefault()

        var username= EmailInputRef.current.value
        var password= PasswordRef.current.value

        if (username==null || username==""){
            alert("Ingrese un nombre de usuario") 
            return; 
        }

        if (password==null || password==""){
            alert("Ingrese una contraseña de usuario") 
            return; 
        }

        var info = {
            "username": username,
            "password": password
        }
        
        fetch(URL, {
            method: 'POST',
            headers: {
                'Accept': 'application/json;charset=utf-8',
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(info)
        })
        .then(response => response.json())
        .then(content => {
            if(content.data != null){
                localStorage.setItem(KEY, content.data);
                window.location.href = "/home"
                return;
            }
            alert("Credenciales invalidas")
        });
    } 

    return (
        <div className="container" id="container">
            <link rel="stylesheet" href="./css/style.css"></link>
            <div className="form-container sing-in-container">
                {/* <form action="./Vista/Dashboard/home.html"> */}
                <form>
                    <h1>Iniciar sesión</h1>
                    <div className="social-container">
                        <a href="" className="social"><i className="fab fa-facebook-f"></i></a>
                        <a href="" className="social"><i className="fab fa-google-plus-g"></i></a>
                        <a href="" className="social"><i className="fab fa-linkedin-in"></i></a>
                    </div>
                    <span>Continuar con otra cuenta</span>
                    <input ref={EmailInputRef} type="email" id="username" placeholder="Correo"/>
                    <input ref={PasswordRef} type="password" id="password" placeholder="Contraseña"/>
                    <button ref={LoginButtonRef} onClick={handleLogin} id="Iniciar" >Iniciar sesión</button>
                </form>
            </div>
        </div>
    )
}