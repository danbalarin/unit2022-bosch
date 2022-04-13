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
import { isBefore, compareAsc } from 'date-fns';

import { useGetItemsQuery } from '../../api/items/getItems';
import type { IItem } from '../../api/types/item';
import type { IItemCount, IJourney } from '../../api/types/journey';

interface Props {
  selectedWarehouse?: number;
  data: IJourney[];
}

export function MaterialTable({ data, selectedWarehouse }: Props) {
  const sortedData = data.sort((a, b) =>
    compareAsc(new Date(a.DepartureTime), new Date(b.DepartureTime))
  );
  const { data: itemData } = useGetItemsQuery({});
  const itemsMap = itemData?.items.reduce((acc, val) => {
    return { ...acc, [val.ID]: val };
  }, {} as Record<number, IItem>);

  const finalData = sortedData.reduce<
    Record<number, Array<Partial<IItemCount & { DepartureTime: string }>>>
  >((acc, val) => {
    val.ItemRequests.forEach(({ WarehouseID, ...rest }) => {
      if (!acc[WarehouseID]) {
        acc[WarehouseID] = [];
      }
      if (selectedWarehouse && selectedWarehouse === WarehouseID) {
        acc[WarehouseID].push({ ...rest, DepartureTime: val.DepartureTime });
      } else if (!selectedWarehouse) {
        acc[WarehouseID].push({ ...rest, DepartureTime: val.DepartureTime });
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
            requests.map(({ itemId, count, DepartureTime }, index) => (
              <Tr key={`${warehouse}-${index}`}>
                {!selectedWarehouse && index === 0 && (
                  <Td isNumeric>{warehouse}</Td>
                )}
                {!selectedWarehouse && index !== 0 && <Td />}
                <Td>{itemsMap?.[itemId ?? 1].name}</Td>
                <Td isNumeric>{count}</Td>
                <Td isNumeric>
                  {isBefore(new Date(DepartureTime ?? ''), new Date()) ? (
                    <Tag variant="subtle" colorScheme="green">
                      Na ceste
                    </Tag>
                  ) : (
                    ''
                  )}
                </Td>
              </Tr>
            ))
          )}
        </Tbody>
      </Table>
    </TableContainer>
  );
}
