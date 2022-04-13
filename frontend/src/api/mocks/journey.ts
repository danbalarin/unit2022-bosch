import { subMinutes, addMinutes } from 'date-fns';

import type { IJourney } from '../types/journey';
import { JourneyStatus } from '../types/journey';

export const JOURNEY_MOCK: IJourney[] = [
  {
    id: '1',
    warehouseId: '1',
    expeditionTime: addMinutes(new Date(), 2).toISOString(),
    items: [
      {
        count: 5,
        itemId: 1,
      },
      { count: 10, itemId: 2 },
    ],
  },
  {
    id: '2',
    warehouseId: '1',
    expeditionTime: subMinutes(new Date(), 10).toISOString(),
    items: [
      {
        count: 3,
        itemId: 3,
      },
      { count: 9, itemId: 4 },
    ],
  },
];
