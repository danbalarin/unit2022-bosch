import type { FormControlProps, SelectProps } from '@chakra-ui/react';
import {
  Select,
  FormControl,
  FormErrorMessage,
  FormLabel,
} from '@chakra-ui/react';
import React from 'react';
import { useController } from 'react-hook-form';

interface IOption {
  value: string | number;
  label: string;
}

interface FormSelectProps extends FormControlProps {
  name: string;
  label: string;
  options: IOption[];
  disabled?: boolean;
  selectProps?: SelectProps;
}

export function FormSelect({
  name,
  label,
  disabled,
  options,
  selectProps,
  ...props
}: FormSelectProps) {
  const {
    field,
    fieldState: { error },
  } = useController({ name });
  return (
    <FormControl {...props} isInvalid={!!error?.message}>
      <FormLabel htmlFor={name}>{label}</FormLabel>
      <Select disabled={disabled} id={name} {...selectProps} {...field}>
        {options.map(({ value, label: optLabel }) => (
          <option key={value} value={value}>
            {optLabel}
          </option>
        ))}
      </Select>
      <FormErrorMessage>
        &nbsp;{error?.message && error.message}
      </FormErrorMessage>
    </FormControl>
  );
}

FormSelect.defaultProps = {
  selectProps: {},
  disabled: false,
};
