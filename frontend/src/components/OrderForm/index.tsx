import React, { useCallback } from 'react';
import { FormProvider, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { Button, useToast } from '@chakra-ui/react';
import { AddIcon } from '@chakra-ui/icons';
import { useGetWarehousesQuery } from 'src/api/items/getWarehouse';

import { useGetJourneysQuery } from '../../api/items/getJourney';
import { useOrderMutation } from '../../api/items/postOrder';
import { useGetItemsQuery } from '../../api/items/getItems';
import { FormSelect } from '../../components/FormSelect';
import { FormInput } from '../../components/FormInput';

const validationSchema = z.object({
  warehouseId: z
    .number()
    .or(z.string().regex(/\d+/u).transform(Number))
    .refine((n) => n >= 0),
  count: z
    .number()
    .or(z.string().regex(/\d+/u).transform(Number))
    .refine((n) => n >= 0),
  itemId: z
    .number()
    .or(z.string().regex(/\d+/u).transform(Number))
    .refine((n) => n >= 0),
});
type FormValues = z.infer<typeof validationSchema>;

interface IOrderFormProps {
  enableWarehouseSelection?: boolean;
}

export function OrderForm({ enableWarehouseSelection }: IOrderFormProps) {
  const formMethods = useForm<FormValues>({
    resolver: zodResolver(validationSchema),
    defaultValues: {
      warehouseId: 1,
      itemId: 1,
      count: 1,
    },
  });
  const { mutateAsync: order } = useOrderMutation({});
  const { handleSubmit, reset } = formMethods;
  const { data: items } = useGetItemsQuery({});
  const { data: warehouses } = useGetWarehousesQuery({});
  const toast = useToast();
  const onSubmit = useCallback(
    handleSubmit(async (formData) => {
      try {
        await order(formData);
        toast({
          title: 'Pridano do objednavky',
          status: 'success',
          duration: 3000,
          isClosable: true,
        });
        reset();
      } catch (err) {
        // setError(true);
      }
    }),
    [handleSubmit]
  );

  return (
    <FormProvider {...formMethods}>
      <form onSubmit={onSubmit}>
        <FormSelect
          id="warehouseId"
          name="warehouseId"
          label="Mezisklad"
          options={
            warehouses?.warehouses.map((val) => ({
              value: val.ID,
              label: val.Name,
            })) ?? []
          }
          mb={4}
          disabled={!enableWarehouseSelection}
        />
        <FormSelect
          id="itemId"
          name="itemId"
          label="Material"
          options={
            items?.items.map((val) => ({
              label: val.name,
              value: val.ID,
            })) ?? []
          }
          mb={4}
        />
        <FormInput
          id="count"
          name="count"
          label="Pocet ks"
          inputProps={{ type: 'number' }}
          mb={4}
        />
        <Button
          size="lg"
          type="submit"
          colorScheme="primary"
          w="100%"
          leftIcon={<AddIcon />}
        >
          Objednat
        </Button>
      </form>
    </FormProvider>
  );
}
