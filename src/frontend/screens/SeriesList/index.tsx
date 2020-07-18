import { useQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

import Series from '../../components/Series';

const SERIES = gql`
  query {
    series {
      id
      name
      poster
      status
      network
      overview
    }
  }
`;

const Wrapper = styled.div`
  display: grid;
  grid-auto-flow: row;
  grid-template-rows: auto 1fr;
  grid-template-areas:
    'title'
    'list';
  height: 100%;
`;

const Title = styled.h4`
  grid-area: title;
`;

const List = styled.div`
  grid-area: list;
  overflow-y: scroll;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  grid-template-rows:  repeat(auto-fill, minmax(350px, max-content));
  column-gap: 32px;
  row-gap 32px;
  margin-top: 16px;
`;

const SeriesList: React.FC = function () {
  const { data, loading, error } = useQuery<{ series: GraphQLTypes.Series[] }>(
    SERIES
  );

  return (
    <Wrapper>
      <Title>TV Series</Title>
      <List>
        {data?.series?.map((series) => (
          <Link key={series.id} to={`/series/${series.id}`}>
            <Series series={series} />
          </Link>
        ))}
      </List>
    </Wrapper>
  );
};

export default SeriesList;
