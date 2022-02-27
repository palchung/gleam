import { getNavItems } from '../helper/getNavItems'

import PatternRoundedIcon from '@mui/icons-material/PatternRounded';
import LockOpenRoundedIcon from '@mui/icons-material/LockOpenRounded';
import HomeMaxRoundedIcon from '@mui/icons-material/HomeMaxRounded';

const routes = {
    Home: { name: 'Home', path: '/home', icon: <HomeMaxRoundedIcon /> },
    Signup: { name: 'Sign Up', path: '/home/signup', icon: <LockOpenRoundedIcon /> },
    Login: { name: 'Login', path: '/home/login', icon: <PatternRoundedIcon /> },
};

const appBarItems = getNavItems(routes, ['Home', 'Signup', 'Login']);

export { routes, appBarItems }
