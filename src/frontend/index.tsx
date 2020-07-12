import 'react-hot-loader';
import { hot } from 'react-hot-loader/root';
import React from 'react';
import ReactDOM from 'react-dom';

function App() {
  return <span>Hello world</span>;
}

const rootElement = document.querySelector('.root');

if (rootElement === null) {
  throw Error('cannot find root element');
}

const HotApp = hot(App);

ReactDOM.render(<HotApp />, rootElement);
