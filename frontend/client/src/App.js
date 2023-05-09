
import './App.css';
import Item from'./pages/items';

import {BrowserRouter as Router, Routers, Route} from 'react-router-dom';

function App() {
  return (
    <div className="App">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css"></link>
        <h1 className="card-panel light-blue lighten-1">Shop</h1>
        <Item />

        <body>

        <form id="formulario_cliente"> 

          <h5>Formulario Registro</h5>
         <input type="text" name="Nombre" placeholder="Nombre"></input>
         <input type="text" lastName="Apellido" placeholder="Apellido"></input>
         <input type="email" email="Email" placeholder="dominio@email.com"></input>
         <input type="text" name="usuario" placeholder="usuario_"></input>
         <input type="password" name="pass" placeholder="Password"></input>
         <p>Estoy de acuerdo con <a href="#">Terminos y Condiciones</a></p>
         <button type="button"> Registrar </button>
         <p><a href="#">¿Ya tengo Cuenta?</a></p>
  
      </form>
    </body>

    <Router>
      <Routers>
        <Route path = '/items/:id' element={<Item />}></Route>
      </Routers>
    </Router>

    </div>


    


  );
}

export default App;
