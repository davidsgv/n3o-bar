import React from "react";
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



export function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Login/>}></Route>
                <Route path="/home" element={<Home/>}></Route>
                <Route path="/admin" element={<Admin/>}></Route>
                <Route path="/categories" element={<Categories/>}></Route>
                <Route path="/inventory" element={<Inventory/>}></Route>
            </Routes>
        </BrowserRouter>
    );
}