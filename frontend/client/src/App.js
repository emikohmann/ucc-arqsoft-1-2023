import './App.css';
import Item from './pages/item'
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';


function App() {
  return (
    <div className="App">
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css"></link>
      <Router>
        <Routes>
          <Route path = '/items/:id' element={<Item />}> </Route>
        </Routes>
      </Router>
    </div>
  );
}

export default App;
