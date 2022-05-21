import React, { Fragment } from "react";
import { useParams } from "react-router-dom";

export function InventoryDetail(){
    // let { id } = useParams();
    // console.log(id)

    var nombre, categoria;

    return (
        <Fragment>
            <section className="full-width header-well">
                <div className="full-width header-well-icon">
                    <i className="zmdi zmdi-account"></i>
                </div>
                <div className="full-width header-well-text">
                    <p className="text-condensedLight">
                        Lorem ipsum dolor sit amet, consectetur adipisicing elit. Unde aut nulla accusantium minus corporis accusamus fuga harum natus molestias necessitatibus.
                    </p>
                </div>
            </section>
            <div className="mdl-grid">
                <div className="mdl-cell mdl-cell--12-col">
                    <div className="full-width panel mdl-shadow--2dp">
                        <div className="full-width panel-tittle bg-primary text-center tittles">
                            Gestionar Productos
                        </div>
                        <div className="full-width panel-content">
                            <form>
                                <div className="mdl-grid">
                                    {/* <div className="mdl-cell mdl-cell--12-col">
                                        <legend className="text-condensedLight"><i className="zmdi zmdi-border-color"></i> &nbsp; DATA ADMINISTRATOR</legend><br/>
                                    </div> */}
                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                        <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <input className="mdl-textfield__input" type="text" pattern="[A-Z]*"/>
                                            <label className="mdl-textfield__label" htmlFor="DNIAdmin">Nombre</label>
                                            <span className="mdl-textfield__error">Nombre invalido</span>
                                        </div>
                                    </div>
                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                    <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <input className="mdl-textfield__input" type="number" pattern="[1-9]*"/>
                                            <label className="mdl-textfield__label" htmlFor="DNIAdmin">Cantidad</label>
                                            <span className="mdl-textfield__error">Cantidad invalida</span>
                                        </div>
                                    </div>

                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                        <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <input className="mdl-textfield__input" type="number" pattern="-?[A-Za-záéíóúÁÉÍÓÚ ]*(\.[0-9]+)?" id="NameAdmin"/>
                                            <label className="mdl-textfield__label" htmlFor="NameAdmin">Precio de compra</label>
                                            <span className="mdl-textfield__error">Precio de compra invalido</span>
                                        </div>
                                    </div>
                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                        <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <input className="mdl-textfield__input" type="text" pattern="-?[A-Za-záéíóúÁÉÍÓÚ ]*(\.[0-9]+)?" id="LastNameAdmin"/>
                                            <label className="mdl-textfield__label" htmlFor="LastNameAdmin">Precio de venta</label>
                                            <span className="mdl-textfield__error">Precio de venta</span>
                                        </div>
                                    </div>
                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                        <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <input className="mdl-textfield__input" type="date" id="phoneAdmin" disabled/>
                                            <label className="mdl-textfield__label" htmlFor="phoneAdmin">Fecha ultimo cambio</label>
                                        </div>
                                    </div>
                                    
                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                        <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <input className="mdl-textfield__input" type="date" id="addressAdmin" disabled/>
                                            <label className="mdl-textfield__label" htmlFor="addressAdmin">Fecha creacion</label>
                                        </div>
                                    </div>
                                    <div className="mdl-cell mdl-cell--12-col">
                                        <legend className="text-condensedLight"><i className="zmdi zmdi-border-color"></i> &nbsp; ACCOUNT DETAILS</legend><br/>
                                    </div>
                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                        <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <input className="mdl-textfield__input" type="checkbox" id="emailAdmin"/>
                                            <label className="mdl-textfield__label" htmlFor="emailAdmin">Activo</label>
                                        </div>
                                    </div>
                                    <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                        <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                            <select className="mdl-textfield__input"  name="" id="">
                                                <option value="1">Categoria de prueba</option>
                                            </select>
                                            <label className="mdl-textfield__label" htmlFor="passwordAdmin">Categoría</label>
                                        </div>
                                    </div>
                                </div>
                            </form>
                            <p className="text-center">
                                <button className="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect mdl-button--colored bg-primary" id="btn-addAdmin">
                                    <i className="zmdi zmdi-plus"></i>
                                </button>
                            </p>
                            <div className="mdl-tooltip" htmlFor="btn-addAdmin">Add Administrator</div>
                        </div>
                    </div>
                </div>
            </div>
        </Fragment>
    )
}