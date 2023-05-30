import Item from './pages/item'

import './App.css';

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';

function App() {
    return (
        <div>
            <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css" />
            
            <Router>
                <Routes>
                    <Route path='/items/:id' element={<Item />}></Route>
                </Routes>
            </Router>

        </div>
    )
}

export default App;