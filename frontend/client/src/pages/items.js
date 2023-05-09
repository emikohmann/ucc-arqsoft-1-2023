import React, {useEffect, useState} from "recat";

import {useParams} from "react-router-dom";

function Item() {
  const [itemData, setItemData] = useState([{}]);
  let (id) = useParams();

  useEffect(() => {
      setItemData(undefined);
      fetch(`http://localhost:5001/api/items/$(id)`).then(
        respo
        .then(
          data =>
        )
      )
    } 

  )


  return (
    <div className="Item">
        <h2 className="card-panel light-blue lighten-2">Titulos del item</h2>
        
        <h2 className="card-panel light-blue lighten-3">$ 100.00</h2>
        
        <p>Descripcion del item</p>
    </div>
  );
}

export default Item;
