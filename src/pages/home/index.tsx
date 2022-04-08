import React from 'react';
import type { FC } from 'react';

import { SEO } from '~/components/Seo';
import { NavbarLayout } from '~/components/NavbarLayout';

export const Home: FC = () => {
  return (
    <NavbarLayout>
      <SEO
        title="STRV Next.js app"
        description="change me before going to production"
      />
      <span>Home page</span>
    </NavbarLayout>
  );
};
