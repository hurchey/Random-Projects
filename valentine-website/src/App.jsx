import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from './Home';
import Suggestions from './Suggestions';
import HeartsBackground from './Heartsbackground'; // if applicable
import './App.css';

function App() {
  return (
    <>
      <HeartsBackground />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/suggestions" element={<Suggestions />} />
      </Routes>
    </>
  );
}

export default App;