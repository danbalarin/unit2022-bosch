import React from 'react';
import type { FC } from 'react';

import { SEO } from '~/components/seo';
import { Logo } from '~/components/logo';

import { Page } from './style';

export const Home: FC = () => {
  return (
    <Page>
      <SEO
        title="STRV Next.js app"
        description="change me before going to production"
      />
      <Logo />
      <span>Home page</span>
    </Page>
  );
};
