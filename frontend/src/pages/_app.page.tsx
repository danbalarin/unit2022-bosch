/* eslint-disable @typescript-eslint/no-explicit-any */
import React from 'react';
import type { AppProps } from 'next/app';
import { ChakraProvider } from '@chakra-ui/react';
import { Hydrate, QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';
import '@fontsource/inter/variable.css';

const App: React.FC<AppProps> = ({ Component, pageProps }) => {
  const [queryClient] = React.useState(() => new QueryClient());

  const QueryClientProviderWithoutType = QueryClientProvider as any;
  const HydrateWithoutType = Hydrate as any;

  return (
    <QueryClientProviderWithoutType client={queryClient}>
      <HydrateWithoutType state={pageProps?.dehydratedState}>
        <ChakraProvider>
          <Component {...pageProps} />
          <ReactQueryDevtools initialIsOpen={false} />
        </ChakraProvider>
      </HydrateWithoutType>
    </QueryClientProviderWithoutType>
  );
};

export default App;
