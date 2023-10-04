import React, { useState } from 'react';
import './App.css';
import {BrowserRouter as Router, Routes, Route, Link} from 'react-router-dom';
import Home from './components/Home';
import Login from './components/Login';
import Register from './components/Register';

function Navbar(){
  return(
    <>
      <button><Link to='/'> Home </Link></button>
      <button><Link to='/Login'> Login </Link></button>
      <button><Link to='/Register'> Register </Link></button>
    </>
  )
}

function App() {
  return (
    <Router>
      <Navbar/>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/Login" element={<Login/>}/>
        <Route path="/Register" element={<Register/>}></Route>
      </Routes>
    </Router>
    
  );
}

export default App;