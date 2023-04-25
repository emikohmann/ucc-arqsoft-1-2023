import React from "react";

import {Routes, Route} from 'react-router-dom';

import Sidebar from '../components/sidebar';

import Home from './home';
import Item from './item';
import Categories from './categories';

const Portal = () => {
  return (
    <>
      <div className="row">
        <div className="col s12 m3 l2 xl2" >
            <Sidebar />
        </div>
        <div className="col s12 m7 l10 xl10">
            <Routes>
                <Route path='/' element={<Home/>} />
                <Route path='/items/:id' element={<Item/>} />
                <Route path='/categories' element={<Categories/>} />
            </Routes>
        </div>
      </div>
    </>
  );
};
  
export default Portal;