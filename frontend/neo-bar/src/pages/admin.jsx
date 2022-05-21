import React, { Fragment, useState } from "react";
import { ListPanel } from "../components/general/Lista/ListPanel";

export function Admin(){

    const [usuarios, setUsuarios] = useState([
        {
            name: "admin@admin.com",
            tag: "Administrador",
            visible: true,
            icon: "zmdi-account",
            id: 1
        }
    ]);
    
    //Datos necesarios para el ListPanel
    const listPanelData = {
        title: "Lista de usuarios",
        searchActive: false,
        //searchFunction: handleSearchKeyUp,
    }

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
            <div className="mdl-tabs mdl-js-tabs mdl-js-ripple-effect">
                <div className="mdl-tabs__tab-bar">
                    <a href="#tabNewAdmin" className="mdl-tabs__tab is-active">NEW</a>
                    <a href="#tabListAdmin" className="mdl-tabs__tab">LIST</a>
                </div>
                <div className="mdl-tabs__panel is-active" id="tabNewAdmin">
                    <div className="mdl-grid">
                        <div className="mdl-cell mdl-cell--12-col">
                            <div className="full-width panel mdl-shadow--2dp">
                                <div className="full-width panel-tittle bg-primary text-center tittles">
                                    New Administrator
                                </div>
                                <div className="full-width panel-content">
                                    <form>
                                        <div className="mdl-grid">
                                            <div className="mdl-cell mdl-cell--12-col">
                                                <legend className="text-condensedLight"><i className="zmdi zmdi-border-color"></i> &nbsp; DATA ADMINISTRATOR</legend><br/>
                                            </div>
                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <input className="mdl-textfield__input" type="number" pattern="-?[0-9]*(\.[0-9]+)?" id="DNIAdmin"/>
                                                    <label className="mdl-textfield__label" htmlFor="DNIAdmin">identificacion</label>
                                                    <span className="mdl-textfield__error">Identificacion invalida</span>
                                                </div>
                                            </div>
                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <select className="mdl-textfield__input" id="tipo-identificacion">
                                                        <option value="0"></option>
                                                        <option value="1">Cedula de ciudadania</option>
                                                        <option value="2">2</option>
                                                    </select>
                                                    <label className="mdl-textfield__label" htmlFor="tipo-identificacion">Tipo Identificacion</label>
                                                    <span className="mdl-textfield__error">Invalid number</span>
                                                </div>
                                            </div>

                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <input className="mdl-textfield__input" type="text" pattern="-?[A-Za-záéíóúÁÉÍÓÚ ]*(\.[0-9]+)?" id="NameAdmin"/>
                                                    <label className="mdl-textfield__label" htmlFor="NameAdmin">Name</label>
                                                    <span className="mdl-textfield__error">Invalid name</span>
                                                </div>
                                            </div>
                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <input className="mdl-textfield__input" type="text" pattern="-?[A-Za-záéíóúÁÉÍÓÚ ]*(\.[0-9]+)?" id="LastNameAdmin"/>
                                                    <label className="mdl-textfield__label" htmlFor="LastNameAdmin">Last Name</label>
                                                    <span className="mdl-textfield__error">Invalid last name</span>
                                                </div>
                                            </div>
                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <input className="mdl-textfield__input" type="tel" pattern="-?[0-9+()- ]*(\.[0-9]+)?" id="phoneAdmin"/>
                                                    <label className="mdl-textfield__label" htmlFor="phoneAdmin">Phone</label>
                                                    <span className="mdl-textfield__error">Invalid phone number</span>
                                                </div>
                                            </div>
                                            
                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <input className="mdl-textfield__input" type="text" id="addressAdmin"/>
                                                    <label className="mdl-textfield__label" htmlFor="addressAdmin">Address</label>
                                                    <span className="mdl-textfield__error">Invalid address</span>
                                                </div>
                                            </div>
                                            <div className="mdl-cell mdl-cell--12-col">
                                                <legend className="text-condensedLight"><i className="zmdi zmdi-border-color"></i> &nbsp; ACCOUNT DETAILS</legend><br/>
                                            </div>
                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <input className="mdl-textfield__input" type="email" id="emailAdmin"/>
                                                    <label className="mdl-textfield__label" htmlFor="emailAdmin">E-mail</label>
                                                    <span className="mdl-textfield__error">Invalid E-mail</span>
                                                </div>
                                            </div>
                                            <div className="mdl-cell mdl-cell--6-col mdl-cell--8-col-tablet">
                                                <div className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                                                    <input className="mdl-textfield__input" type="password" id="passwordAdmin"/>
                                                    <label className="mdl-textfield__label" htmlFor="passwordAdmin">Password</label>
                                                    <span className="mdl-textfield__error">Invalid password</span>
                                                </div>
                                            </div>
                                            {/* <div className="mdl-cell mdl-cell--12-col">
                                                <legend className="text-condensedLight"><i className="zmdi zmdi-border-color"></i> &nbsp; CHOOSE AVATAR</legend><br/>
                                            </div>
                                            <div className="mdl-cell mdl-cell--12-col mdl-cell--8-col-tablet">
                                                <div className="mdl-grid">
                                                    <div className="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone">
                                                        <label className="mdl-radio mdl-js-radio mdl-js-ripple-effect" htmlFor="option-1">
                                                            <input type="radio" id="option-1" className="mdl-radio__button" name="options" value="avatar-male.png"/>
                                                            <img src="assets/img/avatar-male.png" alt="avatar" style={{height: "45px", width:"45px"}}/>
                                                            <span className="mdl-radio__label">Avatar 1</span>
                                                        </label>
                                                    </div>
                                                    <div className="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone">
                                                        <label className="mdl-radio mdl-js-radio mdl-js-ripple-effect" htmlFor="option-2">
                                                            <input type="radio" id="option-2" className="mdl-radio__button" name="options" value="avatar-female.png"/>
                                                            <img src="assets/img/avatar-female.png" alt="avatar" style={{height: "45px", width:"45px"}}/>
                                                            <span className="mdl-radio__label">Avatar 2</span>
                                                        </label>
                                                    </div>
                                                    <div className="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone">
                                                        <label className="mdl-radio mdl-js-radio mdl-js-ripple-effect" htmlFor="option-3">
                                                            <input type="radio" id="option-3" className="mdl-radio__button" name="options" value="avatar-male2.png"/>
                                                            <img src="assets/img/avatar-male2.png" alt="avatar" style={{height: "45px", width:"45px"}}/>
                                                            <span className="mdl-radio__label">Avatar 3</span>
                                                        </label>
                                                    </div>
                                                    <div className="mdl-cell mdl-cell--3-col mdl-cell--4-col-tablet mdl-cell--4-col-phone">
                                                        <label className="mdl-radio mdl-js-radio mdl-js-ripple-effect" htmlFor="option-4">
                                                            <input type="radio" id="option-4" className="mdl-radio__button" name="options" value="avatar-female2.png"/>
                                                            <img src="assets/img/avatar-female2.png" alt="avatar" style={{height: "45px", width:"45px"}}/>
                                                            <span className="mdl-radio__label">Avatar 4</span>
                                                        </label>
                                                    </div>
                                                </div>
                                            </div> */}
                                        </div>
                                        <p className="text-center">
                                            <button className="mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect mdl-button--colored bg-primary" id="btn-addAdmin">
                                                <i className="zmdi zmdi-plus"></i>
                                            </button>
                                        </p>
                                        <div className="mdl-tooltip" htmlFor="btn-addAdmin">Add Administrator</div>
                                    </form>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="mdl-tabs__panel" id="tabListAdmin">
                    <ListPanel panelData={listPanelData} listData={usuarios}/>
                </div>
		    </div>
        </Fragment>
    )
} 