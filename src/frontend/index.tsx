import React from 'react';
import ReactDOM from 'react-dom';
import Modal from 'react-modal';

import App from './screens/App';

const rootElement = document.querySelector('.root');

Modal.setAppElement('.root');

if (rootElement === null) {
  throw Error('cannot find root element');
}

ReactDOM.render(<App />, rootElement);
