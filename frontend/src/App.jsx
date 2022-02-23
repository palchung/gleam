import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';

import Layout from './components/Layout';
import AppPages from './routers/AppPages';


function App() {
  return (
    <Router>
    <div className="App">
            
      <Layout>
        <AppPages />
      </Layout>

    </div>
    </Router>
  );
}

export default App;
