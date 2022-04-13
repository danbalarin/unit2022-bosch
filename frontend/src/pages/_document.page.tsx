import { ColorModeScript } from '@chakra-ui/react';
import NextDocument, { Html, Head, Main, NextScript } from 'next/document';

import { theme } from '~/styles/theme';

export default class MyDocument extends NextDocument {
  render() {
    return (
      <Html lang="en">
        <Head>
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-Black.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-BlackItalic.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-Bold.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-BoldItalic.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-Light.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-LightItalic.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-Medium.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-MediumItalic.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-Regular.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/BoschSans-RegularItalic.woff"
            as="font"
            crossOrigin=""
          />
          <link
            rel="preload"
            href="/fonts/BoschSans/stylesheet.css"
            as="stylesheet"
            crossOrigin=""
          />
        </Head>
        <body>
          <ColorModeScript initialColorMode={theme.config.initialColorMode} />
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}
