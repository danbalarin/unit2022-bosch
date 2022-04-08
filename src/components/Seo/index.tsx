import React from 'react';
import Head from 'next/head';

interface SeoProps {
  title: string;
  description: string;
}

export const SEO = ({ title, description }: SeoProps) => {
  return (
    <Head>
      {/* Primary Meta Tags */}
      <title>{title}</title>
      <meta name="title" content={title} />
      <meta name="description" content={description} />
    </Head>
  );
};
