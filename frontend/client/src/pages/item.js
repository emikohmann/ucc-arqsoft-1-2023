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

      // los corchetes del final son una lista vacia que hace que se ejecute solamente el primer ... ?
    }, []);

    return (
      

      <div className="Item">
      
      {typeof itemData === 'undefined' ? (
        <div><h3>Loading...</h3></div>
      ) : (
        <div>
          <h2>{itemData["name"]}</h2>
          <div class="row">
          <div class="col s7 push-s5">
              <span class="flow-text"> Juegazo del año. 
              Elden Ring (エルデンリング Eruden Ringu?) es un videojuego de rol de acción desarrollado por FromSoftware y publicado por Bandai Namco Entertainment. El videojuego surge de una colaboración entre el director y diseñador Hidetaka Miyazaki y el novelista de fantasía George R. R. Martin. Fue lanzado a nivel mundial el 25 de febrero de 2022 para las plataformas Xbox One, Xbox Series X/S, Microsoft Windows, PlayStation 4, PlayStation 5.1​2​ </span> 
              <h2  id="precio"> ${itemData["price"]}</h2>

              <a class="waves-effect waves-light ltm-large btn-large"> Comprar </a>
          </div>

          <div class="col s5 pull-s7"><span class="flow-text"><img class="responsive-img" src="https://gorilagames.com/img/Public/1019-producto-elden-ring-ps4-1647.jpg"></img></span></div>
      

          </div>
          <hr />
          <ul class="pagination">
                <li class="disabled"><a href="#!"><i class="material-icons">Previo</i></a></li>
                <li class="active"><a href="#!">1</a></li>
                <li class="waves-effect"><a href="#!">2</a></li>
                <li class="waves-effect"><a href="#!">3</a></li>
                <li class="waves-effect"><a href="#!">4</a></li>
                <li class="waves-effect"><a href="#!">5</a></li>
                <li class="waves-effect"><a href="#!"><i class="material-icons">-</i></a></li>
          </ul>
          <hr />
          <div class="col s12 m2">
          <p >Planas S.A ©</p>
          </div>
        </div>
      )}
        
      </div>
    );
    
    
  }
  
  export default Item;
