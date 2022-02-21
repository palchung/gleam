import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import appRoutes from './routers/routes';

import Header from './components/Header';
import './App.css';

const {
  HOME,
  SIGNUP
} = appRoutes

const navItems = [HOME, SIGNUP];

function App() {
  return (
    <Router>
    <div className="App">
      <Header navItems={navItems} />
      <div className="container">
        <Routes>
          <Route path={HOME.route} element={<Home />} />
          <Route path={SIGNUP.route} element={<Signup />} />
        </Routes>
      </div>
    </div>
    </Router>
  );
}

export default App;
