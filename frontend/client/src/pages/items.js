import '../App.css';

// todas las paginas que usan react tienen esto
import React, {useEffect, useState} from "react";

//nos permite adquirir todo lo que este asociado a las request
import {useParams} from "react-router-dom";

function Item() {
    // State es como un objeto que va guardando informacion mientras los componentes sigan vivos
    const [itemData, setItemData] = useState([{}]);
    let {id} = useParams();

    useEffect(() => {
        setItemData(undefined);
        fetch(`http://localhost:5001/api/items/${id}`).then(
            response => response.json()
        ).then(
            data => {
                setItemData(data)
            }
        )

        // los corchetes del final son una lista vacia que hace que se ejecute solamente el primer ... ?
    }, []);

  return (
    <div className="Item">
        <h2 className="card-panel light-blue lighten-2">Titulos del item</h2>
        
        <h2 className="card-panel light-blue lighten-3">$ 100.00</h2>
        
        <p>Descripcion del item</p>
    </div>
  );
}

export default Item;
