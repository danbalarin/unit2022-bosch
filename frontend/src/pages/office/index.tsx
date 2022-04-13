import { Center, VStack } from '@chakra-ui/react';
import React from 'react';
import { useGetJourneysQuery } from 'src/api/items/getJourney';

import { JOURNEY_MOCK } from '../../api/mocks/journey';
import { Card } from '../../components/Card';
import { MaterialTable } from '../../components/MaterialTable';
import { NavbarLayout } from '../../components/NavbarLayout';
import { OrderForm } from '../../components/OrderForm';
import { SEO } from '../../components/Seo';

export function OfficePage() {
  const { data: journeys } = useGetJourneysQuery({});
  return (
    <NavbarLayout>
      <SEO title="Order" description="creating new orders" />
      <Center pt={16}>
        <VStack spacing={8}>
          <Card>
            <OrderForm enableWarehouseSelection />
          </Card>
          <Card>
            <MaterialTable data={journeys?.journeys ?? []} />
          </Card>
        </VStack>
      </Center>
    </NavbarLayout>
  );
}
