import React, {useState, useEffect} from "react";
import { NavBar } from "../components/menu/NavBar";
import { TopBar } from "../components/TopBar";
import { Table } from "../components/inventory/Table";
import { getServerUrl } from "../functions/configuration";

const URL = getServerUrl() + "/api/producto";
const KEY = "login.token"

export function Inventory() {
    const [products, setProducts] = useState([]);

    //obtiene las categorias haciendo un fetch
    const getProducts = ()=>{
        fetch(URL, {
            method: 'GET',
            headers: {
                'Authorization': 'BEARER ' + localStorage.getItem(KEY),
            },
        }).
        then(response => response.json()).
        then(content =>{
            if(content.status.message != null && content.status.message != ""){
                alert(content.status.message);
            }
        
            if(content.data != null){
                var index = 1;
                var products = content.data.map(prod => ({id: index++, nombre: prod.Nombre,
                    stock: prod.Cantidad, category: prod.IdCategoria, purchasePrice: prod.PrecioCompra, 
                    salePrice: prod.PrecioVenta}))
                setProducts(products);
                return
            }
            
            alert("Error");
        })
    }

    useEffect(()=>{
        getProducts();
    },[]);


    return (
        <div>
            <NavBar/>
            <section className="full-width pageContent">
                <TopBar/>
                <section className="full-width header-well">
                    <div className="full-width header-well-icon">
                        <i className="zmdi zmdi-store"></i>
                    </div>
                    <div className="full-width header-well-text">
                        <p className="text-condensedLight">
                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde aut nulla accusantium minus corporis accusamus fuga harum natus molestias necessitatibus.
                        </p>
                    </div>
                </section>
                <div className="full-width divider-menu-h"></div>
                <div className="mdl-grid">
                    <div className="mdl-cell mdl-cell--4-col-phone mdl-cell--8-col-tablet mdl-cell--12-col-desktop">
                        <div className="table-responsive">
                            <Table products={products}/>
                        </div>
                    </div>
                </div>
            </section>
        </div >
    )
}