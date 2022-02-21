import React from 'react';

import Home from '../pages/Home';
import Signup from '../pages/Signup';

const routes = {
    Home:   { name: 'Home', path: '/', element: <Home /> },
    Signup: { name: 'Sign Up', path: '/signup', element: <Signup /> },
};


export default routes ;