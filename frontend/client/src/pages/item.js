import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";

function Item() {
  const [itemData, setItemData] = useState([{}]);
  let {id} = useParams();

  useEffect(() => {
        setItemData(undefined);
        fetch(`http://localhost:5001/api/items/${id}`).then(response => response.json()).then(data => {setItemData(data);})
  }, []);

  return (
    <div className="App">
      {typeof itemData == 'undefined' ? (
        <div><h3>Loading...</h3></div>):(
          <>
          <div><h1>{itemData["name"]}</h1></div>
          <div><h2>{itemData["price"]}</h2></div>
          </>
      )}
    </div>
  );
}

export default Item;