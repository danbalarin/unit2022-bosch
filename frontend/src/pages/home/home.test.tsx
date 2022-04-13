import React from 'react';
import { render, screen } from '@testing-library/react';

import { HomePage } from '.';

describe('pages/home', () => {
  it('should render home page', () => {
    render(<HomePage />);
    expect(screen.queryByText('Home page')).not.toBeNull();
  });
});
