import { theme as chakraTheme } from '@chakra-ui/pro-theme';
import type { ChakraTheme } from '@chakra-ui/react';
import { extendTheme } from '@chakra-ui/react';

const customTheme: Partial<ChakraTheme> = {
  config: {
    initialColorMode: 'light',
  },
  colors: {
    ...chakraTheme.colors,
    red: {
      '50': '#FFE5E6',
      '100': '#FFB8BA',
      '200': '#FF8A8D',
      '300': '#FF5C61',
      '400': '#FF2E34',
      '500': '#FF0008',
      '600': '#CC0006',
      '700': '#990005',
      '800': '#660003',
      '900': '#330002',
    },
    gray: {
      '50': '#F2F2F3',
      '100': '#D9DBDD',
      '200': '#C1C4C7',
      '300': '#A9ADB2',
      '400': '#91959C',
      '500': '#797E86',
      '600': '#61656B',
      '700': '#494C50',
      '800': '#303236',
      '900': '#18191B',
    },
    brand: chakraTheme.colors.gray,
    primary: chakraTheme.colors.red,
  },
  fonts: {
    ...chakraTheme.fonts,
    heading: "'Bosch Sans','Helvetica Neue',Helvetica,Arial,sans-serif",
    body: "'Bosch Sans','Helvetica Neue',Helvetica,Arial,sans-serif",
  },
};

export const theme = extendTheme(customTheme, chakraTheme);
