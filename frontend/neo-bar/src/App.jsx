import React, { Fragment } from "react";
import {
    BrowserRouter,
    Routes,
    Route,
    Link
} from "react-router-dom";
import {Home} from './pages/home'
import { Login } from "./pages/login";
import { Admin } from "./pages/admin";
import { Categories } from "./pages/categories";
import { Inventory } from "./pages/inventory";
import { InventoryDetail } from "./pages/inventoryDetail";
import { NotImplemented } from "./pages/NotImplemented";
import { Menu } from "./components/menu/menu";


export function App() {
    return (
        <Menu>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Login/>}></Route>
                    <Route path="/home" element={<Home/>}></Route>
                    <Route path="/admin" element={<Admin/>}></Route>
                    <Route path="/categories" element={<Categories/>}></Route>
                    <Route path="/product" element={<Inventory/>}>
                    </Route>
                    <Route path="/product/:productId" element={<InventoryDetail/>}/>
                    <Route path="*" element={<NotImplemented/>}></Route>
                </Routes>
            </BrowserRouter>
        </Menu>
    );
}