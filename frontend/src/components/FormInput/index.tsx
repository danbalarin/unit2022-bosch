import type { FormControlProps } from '@chakra-ui/react';
import {
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from '@chakra-ui/react';
import React from 'react';
import { useController } from 'react-hook-form';

interface FormInputProps extends FormControlProps {
  name: string;
  label: string;
  inputProps?: React.HTMLAttributes<HTMLInputElement>;
}

export function FormInput({
  name,
  label,
  inputProps,
  ...props
}: FormInputProps) {
  const {
    field,
    fieldState: { error },
  } = useController({ name });
  return (
    <FormControl {...props} isInvalid={!!error?.message}>
      <FormLabel htmlFor={name}>{label}</FormLabel>
      <Input id={name} {...inputProps} {...field} />
      <FormErrorMessage>
        &nbsp;{error?.message && error.message}
      </FormErrorMessage>
    </FormControl>
  );
}

FormInput.defaultProps = {
  inputProps: {},
};
