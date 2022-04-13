import type { BoxProps } from '@chakra-ui/react';
import { Box } from '@chakra-ui/react';
import React from 'react';

export function Card(props: BoxProps) {
  return (
    <Box
      p={8}
      w="full"
      bg="white"
      boxShadow="2xl"
      rounded="md"
      overflow="hidden"
      minW="320px"
      {...props}
    />
  );
}
