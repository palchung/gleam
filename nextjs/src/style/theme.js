import { createTheme } from '@mui/material/styles';
import { amber, red, pink } from '@mui/material/colors';

// Create a theme instance.
const theme = createTheme({
  palette: {
    primary: {
      main: amber[400],
    },
    secondary: {
      main: pink[400],
    },
    error: {
      main: red[400],
    },
  },
});

export default theme;
