import React from 'react';
import { Link } from 'react-router-dom';
import routes from '../routers/Routes';

const menu = ['Home', 'Signup'];

function getNavItems() {
  const n = Object.keys(routes)
  .filter(key => menu.includes(key))
  .reduce((obj, key) => {
    obj[key] = routes[key];
    return obj;
  }, {});
  return Object.values(n);
};

const navItems = getNavItems();

function Header() {
  return <ul className='App-nav-list'>
    {
      navItems.map((navItem, i) => (
        <li key={i} className="App-nav-item">
            <Link to={navItem.path}>{navItem.name}</Link>
        </li>
      ))
    }
  </ul>
}

export default Header;