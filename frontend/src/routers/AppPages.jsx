import { useRoutes } from 'react-router-dom';
import routes from './Routes';

function AppPages() {
    return useRoutes(Object.values(routes));
};

export default AppPages;