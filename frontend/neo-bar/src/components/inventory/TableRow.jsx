import React from "react";

export function TableRow({product}){
    const {name, stock, category, purchasePrice, salePrice} = product;
    return (
        <tr>
            <td className="mdl-data-table__cell--non-numeric">{name}</td>
            <td>{stock}</td>
            <td>{category}</td>
            <td>{purchasePrice}</td>
            <td>{salePrice}</td>
            <td><button className="mdl-button mdl-button--icon mdl-js-button mdl-js-ripple-effect"><i className="zmdi zmdi-more"></i></button></td>
        </tr>
    )
}