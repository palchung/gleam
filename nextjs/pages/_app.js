import * as React from 'react';
import PropTypes from 'prop-types';
import Head from 'next/head';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider } from '@emotion/react';
import theme from '../src/style/theme';
import createEmotionCache from '../src/helper/createEmotionCache';
import layouts from '../src/components/layout/appLayout';
import Copyright from '../src/components/modules/copyright';
import { AuthProvider } from '../src/api/auth/authContext';

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache();

export default function MyApp(props) {
  const { Component, emotionCache = clientSideEmotionCache, pageProps } = props;

  let Layout
  typeof Component.layout === 'undefined' ? Layout = layouts['Default'] : Layout = layouts[Component.layout]
  // const Layout = layouts[Component.layout] || ((children) => <>{children}</>);

  return (
    <CacheProvider value={emotionCache}>

      <Head>
        <meta name="viewport" content="initial-scale=1, width=device-width" />
      </Head>

      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <AuthProvider>
          <Layout>
            <Component {...pageProps} />
            <Copyright />
          </Layout>
        </AuthProvider>
      </ThemeProvider>
    </CacheProvider>
  );
}

MyApp.propTypes = {
  Component: PropTypes.elementType.isRequired,
  emotionCache: PropTypes.object,
  pageProps: PropTypes.object.isRequired,
};
