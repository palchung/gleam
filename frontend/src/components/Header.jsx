import React from 'react';
import { Link } from 'react-router-dom';

function Header({ navItems }) {
  return <ul className='App-nav-list'>
    {
      navItems.map(navItem =>
        <li className='App-nav-item'>
          <Link to={navItem.route}>{navItem.name}</Link>
        </li>  
      )
    }
  </ul>
}

export default Header;