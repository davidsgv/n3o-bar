import React from "react";
import { TableRow } from "./TableRow";

export function Table({products}){
    if(products.length > 0){
        return (
            <table className="mdl-data-table mdl-js-data-table mdl-shadow--2dp full-width table-responsive">
                <thead>
                    <tr>
                        <th className="mdl-data-table__cell--non-numeric">Name</th>
                        <th>Stock</th>
                        <th>Category</th>
                        <th>Purchase price</th>
                        <th>Sale price</th>
                        <th>options</th>
                    </tr>
                </thead>
                <tbody>
                    {products.map((product,index) =>(<TableRow key={index} product={product}/>))}
                </tbody>
            </table>
        );
    }
    else{
        return (
            <table className="mdl-data-table mdl-js-data-table mdl-shadow--2dp full-width table-responsive">
                <thead>
                    <tr>
                        <th className="mdl-data-table__cell--non-numeric">Name</th>
                        <th>Stock</th>
                        <th>Category</th>
                        <th>Purchase price</th>
                        <th>Sale price</th>
                        <th>options</th>
                    </tr>
                </thead>
            </table>
        )
    }
}