import 'react-hot-loader';

import { ApolloProvider } from '@apollo/react-hooks';
import ApolloClient from 'apollo-boost';
import React from 'react';
import { hot } from 'react-hot-loader/root';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import styled from 'styled-components';

import Header from '../components/Header';
import SeriesAdd from './SeriesAdd';
import SeriesList from './SeriesList';

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

export default hot(App);
