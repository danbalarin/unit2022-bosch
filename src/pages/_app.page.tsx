import React from 'react';
import type { AppProps } from 'next/app';
import NextApp from 'next/app';
import { ChakraProvider } from '@chakra-ui/react';
import '@fontsource/inter/variable.css';

const App: React.FC<AppProps> = (props) => (
  <ChakraProvider>
    <NextApp {...props} />
  </ChakraProvider>
);

export default App;
