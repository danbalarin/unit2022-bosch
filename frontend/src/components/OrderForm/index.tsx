import React, { useCallback } from 'react';
import { FormProvider, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { Button } from '@chakra-ui/react';
import { AddIcon } from '@chakra-ui/icons';

import { useGetItemsQuery } from '../../api/items/getItems';
import { FormSelect } from '../../components/FormSelect';
import { FormInput } from '../../components/FormInput';

const validationSchema = z.object({
  warehouse: z.string().nonempty(),
  quantity: z
    .number()
    .or(z.string().regex(/\d+/u).transform(Number))
    .refine((n) => n >= 0),
  item: z.string().nonempty(),
});
type FormValues = z.infer<typeof validationSchema>;

const WAREHOUSES = ['1', '2', '3', '4', '5'];

interface IOrderFormProps {
  enableWarehouseSelection?: boolean;
}

export function OrderForm({ enableWarehouseSelection }: IOrderFormProps) {
  const formMethods = useForm<FormValues>({
    resolver: zodResolver(validationSchema),
    defaultValues: {
      warehouse: '1',
      item: '1',
      quantity: 1,
    },
  });
  const { handleSubmit } = formMethods;
  const { data } = useGetItemsQuery({});
  const onSubmit = useCallback(
    handleSubmit(async (formData) => {
      try {
        console.log(formData);
        // await handleUserSignIn(formData);
        // setError(false);
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
          id="warehouse"
          name="warehouse"
          label="Mezisklad"
          options={WAREHOUSES.map((val) => ({ value: val, label: val }))}
          mb={4}
          disabled={!enableWarehouseSelection}
        />
        <FormSelect
          id="item"
          name="item"
          label="Material"
          options={
            data?.items.map((val) => ({
              label: val.name,
              value: val.ID,
            })) ?? []
          }
          mb={4}
        />
        <FormInput
          id="quantity"
          name="quantity"
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
