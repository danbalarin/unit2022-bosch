/* eslint-disable @typescript-eslint/no-explicit-any */
import React from 'react';
import type { AppProps } from 'next/app';
import { ChakraProvider } from '@chakra-ui/react';
import { Hydrate, QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';

import '@fontsource/inter/variable.css';
import { theme } from '../styles/theme';
import { AuthContextProvider } from '../contexts/user';

const App: React.FC<AppProps> = ({ Component, pageProps }) => {
  const [queryClient] = React.useState(() => new QueryClient());

  const QueryClientProviderWithoutType = QueryClientProvider as any;
  const HydrateWithoutType = Hydrate as any;

  return (
    <QueryClientProviderWithoutType client={queryClient}>
      <HydrateWithoutType state={pageProps?.dehydratedState}>
        <AuthContextProvider>
          <ChakraProvider theme={theme}>
            <Component {...pageProps} />
            <ReactQueryDevtools initialIsOpen={false} />
          </ChakraProvider>
        </AuthContextProvider>
      </HydrateWithoutType>
    </QueryClientProviderWithoutType>
  );
};

export default App;
