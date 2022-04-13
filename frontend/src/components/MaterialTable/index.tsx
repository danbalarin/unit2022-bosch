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
import type { IJourney } from '../../api/types/journey';

interface Props {
  selectedWarehouse?: string;
  data: IJourney[];
}

export function MaterialTable({ data, selectedWarehouse }: Props) {
  const filteredData = data
    .filter((val) =>
      selectedWarehouse ? val.warehouseId === selectedWarehouse : true
    )
    .sort((a, b) =>
      compareAsc(new Date(a.expeditionTime), new Date(b.expeditionTime))
    );
  const { data: itemData } = useGetItemsQuery({});
  const itemsMap = itemData?.items.reduce((acc, val) => {
    return { ...acc, [val.ID]: val };
  }, {} as Record<string, IItem>);
  return (
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
          {filteredData.map((journey) =>
            journey.items.map(({ itemId, count }, index) => (
              <Tr key={index}>
                <Td>{itemsMap?.[itemId].name}</Td>
                <Td isNumeric>{count}</Td>
                <Td isNumeric>
                  {isBefore(new Date(journey.expeditionTime), new Date()) ? (
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
