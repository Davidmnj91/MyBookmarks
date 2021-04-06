import React from 'react';
import ReactDOM from 'react-dom';
import styled, {css} from 'styled-components';

const Red = styled.h1`
  color: red;
`;

const App = () => (
  <div
    css={css`
      background-color: green;
    `}
  >
    <Red>My React and TypeScript App!</Red>
  </div>
);

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
