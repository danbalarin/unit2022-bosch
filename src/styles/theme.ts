import { theme as chakraTheme } from '@chakra-ui/pro-theme';
import type { ChakraTheme } from '@chakra-ui/react';
import { extendTheme } from '@chakra-ui/react';

const colors: Partial<ChakraTheme> = {
  config: {
    initialColorMode: 'dark',
  },
  colors: { ...chakraTheme.colors, brand: chakraTheme.colors.purple },
};

export const theme = extendTheme({ colors }, chakraTheme);
