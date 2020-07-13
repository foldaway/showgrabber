import 'react-hot-loader';
import { hot } from 'react-hot-loader/root';
import React from 'react';
import ReactDOM from 'react-dom';
import styled from 'styled-components';
import Header from './components/Header';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from '@apollo/react-hooks';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import SeriesList from './screens/SeriesList';
import SeriesAdd from './screens/SeriesAdd';

const Wrapper = styled.div`
  font-family: -apple-system, 'Segoe UI', sans-serif;
`;

const client = new ApolloClient({
  uri: '/graphql',
});

function App() {
  return (
    <ApolloProvider client={client}>
      <Wrapper>
        <Header />
        <BrowserRouter>
          <Switch>
            <Route path="/add" component={SeriesAdd} exact />
            <Route path="/" component={SeriesList} />
          </Switch>
        </BrowserRouter>
      </Wrapper>
    </ApolloProvider>
  );
}

const rootElement = document.querySelector('.root');

if (rootElement === null) {
  throw Error('cannot find root element');
}

const HotApp = hot(App);

ReactDOM.render(<HotApp />, rootElement);
