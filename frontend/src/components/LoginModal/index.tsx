import type { ModalProps } from '@chakra-ui/react';
import {
  Alert,
  AlertIcon,
  AlertTitle,
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  useDisclosure,
} from '@chakra-ui/react';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import React, { useCallback, useEffect, useState } from 'react';
import { FormProvider, useForm } from 'react-hook-form';

import { FormInput } from '../FormInput';
import { useAuthContext } from '../../contexts/user';

type Props = Partial<ModalProps> & {};

const validationSchema = z.object({
  email: z.string().email().nonempty(),
  password: z.string().nonempty(),
});
type FormValues = z.infer<typeof validationSchema>;

export function LoginModal(props: Props) {
  const formMethods = useForm<FormValues>({
    resolver: zodResolver(validationSchema),
  });
  const { handleSubmit } = formMethods;
  const [isError, setError] = useState(false);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { user, handleUserSignIn } = useAuthContext();

  const onSubmit = useCallback(
    handleSubmit(async (formData) => {
      try {
        await handleUserSignIn(formData);
        setError(false);
      } catch (err) {
        setError(true);
      }
    }),
    [handleSubmit]
  );

  useEffect(() => {
    user?.token && onClose();
  }, [onClose, user]);

  const errorMessage = isError && 'Invalid credentials';

  return (
    <>
      <Button variant="solid" colorScheme="primary" size="sm" onClick={onOpen}>
        Login
      </Button>
      <Modal
        motionPreset="slideInBottom"
        isOpen={isOpen}
        onClose={onClose}
        {...props}
      >
        <ModalOverlay />
        <ModalContent pb={5}>
          <ModalHeader>Login</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <>
              {errorMessage && (
                <Alert status="error" mb={2}>
                  <AlertIcon />
                  <AlertTitle mr={2}>{errorMessage}</AlertTitle>
                </Alert>
              )}
              <FormProvider {...formMethods}>
                <form onSubmit={onSubmit}>
                  <FormInput
                    name="email"
                    label="Email"
                    inputProps={{ placeholder: 'Email', type: 'email' }}
                    mb={2}
                  />
                  <FormInput
                    name="password"
                    label="Password"
                    inputProps={{ placeholder: 'Password', type: 'password' }}
                    mb={2}
                  />
                  <Button mt={2} w="100%" type="submit" colorScheme="primary">
                    Submit
                  </Button>
                </form>
              </FormProvider>
            </>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}
