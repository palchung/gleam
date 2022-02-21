import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import AppPages from './routers/AppPages';

import Header from './components/Header';
import './App.css';


function App() {
  return (
    <Router>
    <div className="App">
      
      <Header />
      
      <div className="container">
        <AppPages />
      </div>

    </div>
    </Router>
  );
}

export default App;
