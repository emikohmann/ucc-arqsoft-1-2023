import logo from './logo.svg';
import './App.css';
import Item from './pages/item'

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css" />
      <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" />
      <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
      
      <body>
        <div class="wrapper text-lighten-2 text-aling-center">
          <h3>Shop</h3>
        </div>

        <Router>
          <Routes>
              <Route path='/items/:id' element={<Item />} />
          </Routes>
        </Router>
      </body>
    </div>
  );
}

export default App;
