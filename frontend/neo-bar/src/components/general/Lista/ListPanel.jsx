import React, { useRef } from "react";
import { List } from "./List";

export function ListPanel({panelData, listData}){
    const {title, searchFunction, searchActive} = panelData;
    const list = listData

    //filto de busqueda de categorias
    const searchInputRef =  useRef();

    const handleSearch = ()=>{
        searchFunction(searchInputRef);
    }

    
    if(searchActive){
        var searchContent = (
            <form action="#">
                <div className="mdl-textfield mdl-js-textfield mdl-textfield--expandable">
                    <label className="mdl-button mdl-js-button mdl-button--icon" htmlFor="searchAdmin">
                        <i className="zmdi zmdi-search"></i>
                    </label>
                    <div className="mdl-textfield__expandable-holder">
                        <input ref={searchInputRef} onChange={handleSearch} className="mdl-textfield__input" type="text" id="searchAdmin"/>
                        <label className="mdl-textfield__label"></label>
                    </div>
                </div>
            </form>
        )
    }

    return (
        <div className="mdl-grid">
            <div className="mdl-cell mdl-cell--4-col-phone mdl-cell--8-col-tablet mdl-cell--8-col-desktop mdl-cell--2-offset-desktop">
                <div className="full-width panel mdl-shadow--2dp">
                    <div className="full-width panel-tittle bg-success text-center tittles">
                        {title}
                    </div>
                    <div className="full-width panel-content">
                        {searchContent}
                        <List list={list}/>
                    </div>
                </div>
            </div>
        </div>
    )
}