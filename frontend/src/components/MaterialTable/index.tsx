import {
  Table,
  TableContainer,
  Tag,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import React from 'react';
import { compareAsc } from 'date-fns';

import { useGetWarehousesQuery } from '../../api/items/getWarehouse';
import type { IItemRequest, IJourney } from '../../api/types/journey';

interface Props {
  selectedWarehouse?: number;
  data: IJourney[];
}

export function MaterialTable({ data, selectedWarehouse }: Props) {
  const sortedData = data.sort((a, b) =>
    compareAsc(new Date(a.DepartureTime), new Date(b.DepartureTime))
  );
  const { data: warehouses } = useGetWarehousesQuery({});

  const finalData = sortedData.reduce<
    Record<number, Array<IItemRequest & { Departed: boolean }>>
  >((acc, val) => {
    val.ItemRequests.forEach(({ WarehouseID, ...rest }) => {
      if (!acc[WarehouseID]) {
        acc[WarehouseID] = [];
      }
      if (selectedWarehouse && selectedWarehouse === WarehouseID) {
        acc[WarehouseID].push({ ...rest, WarehouseID, Departed: val.Departed });
      } else if (!selectedWarehouse) {
        acc[WarehouseID].push({ ...rest, WarehouseID, Departed: val.Departed });
      }
    });
    return acc;
  }, {});

  return (
    <TableContainer>
      <Table size="sm">
        <Thead>
          <Tr>
            {!selectedWarehouse && <Th>Mezisklad</Th>}
            <Th>Jmeno</Th>
            <Th isNumeric>Pocet</Th>
            <Th />
          </Tr>
        </Thead>
        <Tbody>
          {Object.entries(finalData).map(([warehouse, requests]) =>
            requests.map(
              ({ Item: { name: itemName }, counts, Departed }, index) => (
                <Tr key={`${warehouse}-${index}`}>
                  {!selectedWarehouse && index === 0 && (
                    <Td isNumeric>
                      {warehouses?.warehouses.find((w) => w.ID === +warehouse)
                        ?.Name ?? warehouse}
                    </Td>
                  )}
                  {!selectedWarehouse && index !== 0 && <Td />}
                  <Td>{itemName}</Td>
                  <Td isNumeric>{counts}</Td>
                  <Td isNumeric>
                    {Departed ? (
                      <Tag variant="subtle" colorScheme="green">
                        Na ceste
                      </Tag>
                    ) : (
                      ''
                    )}
                  </Td>
                </Tr>
              )
            )
          )}
        </Tbody>
      </Table>
    </TableContainer>
  );
}
