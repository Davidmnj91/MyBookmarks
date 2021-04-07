import React from 'react';
import ReactDOM from 'react-dom';
import reset from 'styled-reset';
import { App } from './App';
import { createGlobalStyle } from 'styled-components';

const GlobalStyle = createGlobalStyle`
  ${reset}
  html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, dl, dt, dd, menu, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, main, menu, nav, output, ruby, section, summary, time, mark, audio, video {
    box-sizing: border-box;
    font-family: 'Poppins', sans-serif;
  }

  ::-webkit-scrollbar-track {
    background-color: #2c5364;
  }

  ::-webkit-scrollbar {
    width: 12px;
    background-color: #F5F5F5;
  }

  ::-webkit-scrollbar-thumb {
    border-radius: 10px;
    border: 3px solid #2c5364;
    background-color: #858585;
  }
`;

ReactDOM.render(
  <React.StrictMode>
    <GlobalStyle />
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
