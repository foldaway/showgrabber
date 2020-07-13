import React from 'react';
import ReactDOM from 'react-dom';
import App from './screens/App';

const rootElement = document.querySelector('.root');

if (rootElement === null) {
  throw Error('cannot find root element');
}

ReactDOM.render(<App />, rootElement);
