import React, { useCallback } from 'react';
import { FormProvider, useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { Button, Center } from '@chakra-ui/react';

import { FormInput } from '~/components/FormInput';
import { NavbarLayout } from '~/components/NavbarLayout';

const validationSchema = z.object({
  firstName: z.string().nonempty(),
  lastName: z.string().nonempty(),
});
type FormValues = z.infer<typeof validationSchema>;

export function FormPage() {
  const formMethods = useForm<FormValues>({
    resolver: zodResolver(validationSchema),
  });
  const { handleSubmit } = formMethods;

  const onSubmit = useCallback(
    handleSubmit((data) => console.log(data)),
    [handleSubmit]
  );

  return (
    <NavbarLayout>
      <Center mt={8}>
        <FormProvider {...formMethods}>
          <form onSubmit={onSubmit}>
            <FormInput
              name="firstName"
              label="First name"
              inputProps={{ placeholder: 'First name' }}
              mb={2}
            />
            <FormInput
              name="lastName"
              label="Last name"
              inputProps={{ placeholder: 'Last name' }}
              mb={2}
            />
            <Button w="100%" type="submit" colorScheme="teal">
              Submit
            </Button>
          </form>
        </FormProvider>
      </Center>
    </NavbarLayout>
  );
}
