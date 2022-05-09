import React, {useRef, useState, useEffect} from "react";
import { NavBar } from "../components/menu/NavBar";
import { TopBar } from "../components/TopBar";
import { getServerUrl } from "../functions/configuration";
import { CategoriesList } from "../components/categories/CategoriesList";

const URL = getServerUrl() + "/api/categoria/";
const KEY = "login.token"

export function Categories(){

    //crear componente con las categorias
    const [categories, setCategories] = useState([]);

    //obtiene las categorias haciendo un fetch
    const getCategories = ()=>{
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
                var categories = content.data.map(cat => ({id: index++, nombre: cat.Nombre, visible: true}))
                setCategories(categories);
                return
            }
            
            alert("Error");
        })
    }

    //obtener las categorias al cargar la pagina
    useEffect(()=>{
        getCategories();
    },[]);

    //filto de busqueda de categorias
    const searchInputRef =  useRef();

    //evento del filtro de categorias
    const handleSearchKeyUp = ()=>{
        const text = searchInputRef.current.value;
        if(text.length == 0 ){
            var allCategories = [...categories];
            allCategories.map(cat =>{
                cat.visible = true;
            })
            setCategories(allCategories);
            return
        }

        var filter = categories.map(cat => {
            let shouldRender = cat.nombre.toLowerCase().indexOf(text.toLowerCase()) >= 0
            cat.visible = shouldRender
            return cat;
        });
        console.log(filter)
        setCategories(filter);
    }

    //agregar categoria
    const NombreInput =  useRef();

    const handleAddCategory = (e)=>{
        console.log(e);
        e.preventDefault();

        const nombre = NombreInput.current.value;

        if (nombre==null || nombre==""){
            alert("Ingrese una categoria") 
            return; 
        }

        var info = {
            "nombre": nombre,
        }

        fetch(URL, {
            method: 'POST',
            headers: {
                'Accept': 'application/json;charset=utf-8',
                'Content-Type': 'application/json;charset=utf-8',
                'Authorization': 'BEARER ' + localStorage.getItem(KEY),
            },
            body: JSON.stringify(info)
        }).
        then(response => response.json()).
        then(content =>{
            if(content.status.message != null && content.status.message != ""){
                alert(content.status.message)
                return;
            }
            alert("Categoria creada")
        });
    
        
    };

    return (
        <div>
            <NavBar/>
            <section className="full-width pageContent">
                <TopBar/>
                <section className="full-width header-well">
                    <div className="full-width header-well-icon">
                        <i className="zmdi zmdi-label"></i>
                    </div>
                    <div className="full-width header-well-text">
                        <p className="text-condensedLight">
                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde aut nulla accusantium minus corporis accusamus fuga harum natus molestias necessitatibus.
                        </p>
                    </div>
                </section>
                <div className="mdl-tabs mdl-js-tabs mdl-js-ripple-effect">
                    <div className="mdl-tabs__tab-bar">
                        <a href="#tabNewCategory" className="mdl-tabs__tab is-active">Nueva</a>
                        <a href="#tabListCategory" className="mdl-tabs__tab" id="lista-categorias">Lista</a>
                    </div>
                    <div className="mdl-tabs__panel is-active" id="tabNewCategory">
                        <div className="mdl-grid">
                            <div className="mdl-cell mdl-cell--12-col">
                                <div className="full-width panel mdl-shadow--2dp">
                                    <div className="full-width panel-tittle bg-primary text-center tittles">
                                        Agregar categoria
                                    </div>
                                    <div className="full-width panel-content">
                                        <form>
                                            <div className="mdl-grid">
                                                <div className="mdl-cell mdl-cell--12-col">
                                                    <legend className="text-condensedLight"><i className="zmdi zmdi-border-color"></i>Categorias</legend><br/>
                                                </div>
                                                <div className="mdl-cell mdl-cell--12-col mdl-cell--12-col-tablet">
                                                    <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                        <input ref={NombreInput} className="mdl-textfield__input" type="text" pattern="-?[A-Za-z0-9áéíóúÁÉÍÓÚ ]*(\.[0-9]+)?" id="NameCategory"/>
                                                        <label className="mdl-textfield__label" htmlFor="NameCategory">Nombre</label>
                                                        <span className="mdl-textfield__error">Invalid name</span>
                                                    </div>
                                                </div>
                                            </div>
                                            <p className="text-center">
                                                <button onClick={handleAddCategory} className="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect mdl-button--colored bg-primary" id="btn-addCategory">
                                                    <i className="zmdi zmdi-plus"></i>
                                                </button>
                                            </p>
                                            <div className="mdl-tooltip" htmlFor="btn-addCategory">AgregarCategoria</div>
                                        </form>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="mdl-tabs__panel" id="tabListCategory">
                        <div className="mdl-grid">
                            <div className="mdl-cell mdl-cell--4-col-phone mdl-cell--8-col-tablet mdl-cell--8-col-desktop mdl-cell--2-offset-desktop">
                                <div className="full-width panel mdl-shadow--2dp">
                                    <div className="full-width panel-tittle bg-success text-center tittles">
                                        Lista de categorias
                                    </div>
                                    <div className="full-width panel-content">
                                        <form action="#">
                                            <div className="mdl-textfield mdl-js-textfield mdl-textfield--expandable">
                                                <label className="mdl-button mdl-js-button mdl-button--icon" htmlFor="searchCategory">
                                                    <i className="zmdi zmdi-search"></i>
                                                </label>
                                                <div className="mdl-textfield__expandable-holder">
                                                    <input ref={searchInputRef} onChange={handleSearchKeyUp} className="mdl-textfield__input" type="text" id="searchCategory"/>
                                                    <label className="mdl-textfield__label"></label>
                                                </div>
                                            </div>
                                        </form>
                                        <CategoriesList categories={categories}/>

                                        {/* <div className="mdl-list" id="categorias-container">
                                        </div> */}
                                    </div>
                                </div>
                                
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    )
}