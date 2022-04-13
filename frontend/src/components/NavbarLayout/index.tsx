import type { ReactNode } from 'react';
import type { LinkProps } from '@chakra-ui/react';
import {
  Box,
  Flex,
  Link,
  Button,
  useColorModeValue,
  Stack,
  useColorMode,
  useDisclosure,
  IconButton,
  HStack,
} from '@chakra-ui/react';
import {
  AddIcon,
  CloseIcon,
  HamburgerIcon,
  MoonIcon,
  SunIcon,
} from '@chakra-ui/icons';

import { ROUTES } from '~/constants/routes';
import Logo from '../Logo';

type NavLinkProps = LinkProps;

const NavLink = (props: NavLinkProps) => (
  <Link
    px={2}
    py={1}
    rounded="md"
    _hover={{
      textDecoration: 'none',
      bg: useColorModeValue('gray.200', 'gray.700'),
    }}
    {...props}
  />
);

export function NavbarLayout({ children }: { children: ReactNode }) {
  const { colorMode, toggleColorMode } = useColorMode();
  const { isOpen, onOpen, onClose } = useDisclosure();
  return (
    <>
      <Box bg={useColorModeValue('gray.100', 'gray.900')} px={4}>
        <Flex h={16} alignItems="center" justifyContent="space-between">
          <IconButton
            size="sm"
            icon={isOpen ? <CloseIcon /> : <HamburgerIcon />}
            aria-label="Open Menu"
            display={{ md: 'none' }}
            onClick={isOpen ? onClose : onOpen}
          />
          <HStack maxH={'100%'} spacing={8} alignItems="center">
            <Logo style={{ maxHeight: '100%', width: '100px' }} />
            <HStack as="nav" spacing={4} display={{ base: 'none', md: 'flex' }}>
              {ROUTES.map((route) => (
                <NavLink key={route.path} href={route.path}>
                  {route.label}
                </NavLink>
              ))}
            </HStack>
          </HStack>
          <Flex alignItems="center">
            {/* <Button
              variant="solid"
              colorScheme="teal"
              size="sm"
              mr={4}
              leftIcon={<AddIcon />}
            >
              Action
            </Button> */}
            <Button size="sm" onClick={toggleColorMode}>
              {colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
            </Button>
          </Flex>
        </Flex>

        {isOpen ? (
          <Box pb={4} display={{ md: 'none' }}>
            <Stack as="nav" spacing={4}>
              {ROUTES.map((route) => (
                <NavLink key={route.path} href={route.path}>
                  {route.label}
                </NavLink>
              ))}
            </Stack>
          </Box>
        ) : null}
      </Box>
      <main>{children}</main>
    </>
  );
}
