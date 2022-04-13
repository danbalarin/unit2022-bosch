import React from 'react';
import type { FC } from 'react';
import {
  Center,
  Table,
  TableContainer,
  Tbody,
  Thead,
  Tr,
  Th,
  VStack,
  Checkbox,
  Td,
} from '@chakra-ui/react';

import { Card } from '../../components/Card';
import { SEO } from '../../components/Seo';
import { NavbarLayout } from '../../components/NavbarLayout';

export const WarehousePage: FC = () => {
  const MOCK_DATA = [
    { name: 'Sroubek', count: 1 },
    { name: 'Lopata', count: 2 },
    { name: 'Matka', count: 5 },
    { name: 'Lozisko', count: 15 },
  ];

  return (
    <NavbarLayout>
      <SEO title="Warehouse" description="" />
      <Center pt={16}>
        <VStack spacing={8}>
          <Card>
            <TableContainer>
              <Table size="sm">
                <Thead>
                  <Tr>
                    <Th>Jmeno</Th>
                    <Th isNumeric>Pocet</Th>
                    <Th />
                  </Tr>
                </Thead>
                <Tbody>
                  {MOCK_DATA.map(({ name, count }, index) => (
                    <Tr key={index}>
                      <Td>{name}</Td>
                      <Td isNumeric>{count}</Td>
                      <Td>
                        <Checkbox />
                      </Td>
                    </Tr>
                  ))}
                </Tbody>
              </Table>
            </TableContainer>
          </Card>
        </VStack>
      </Center>
    </NavbarLayout>
  );
};
