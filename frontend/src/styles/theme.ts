import { theme as chakraTheme } from '@chakra-ui/pro-theme';
import type { ChakraTheme } from '@chakra-ui/react';
import { extendTheme } from '@chakra-ui/react';

const customTheme: Partial<ChakraTheme> = {
  config: {
    initialColorMode: 'dark',
  },
  colors: { ...chakraTheme.colors, brand: chakraTheme.colors.purple },
  fonts: {
    ...chakraTheme.fonts,
    heading: "'Bosch Sans','Helvetica Neue',Helvetica,Arial,sans-serif",
    body: "'Bosch Sans','Helvetica Neue',Helvetica,Arial,sans-serif",
  },
};

export const theme = extendTheme(customTheme, chakraTheme);
