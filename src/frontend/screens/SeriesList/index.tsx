import React from 'react';
import styled from 'styled-components';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';
import Series from '../../components/Series';

const SERIES = gql`
  query {
    series {
      id
      name
    }
  }
`;

const Wrapper = styled.div``;

const Title = styled.h4``;

const SeriesList: React.FC = function () {
  const { data, loading, error } = useQuery<{ series: GraphQLTypes.Series[] }>(
    SERIES
  );

  return (
    <Wrapper>
      <Title>TV Series</Title>
      {data?.series?.map((series) => (
        <Series key={series.id} series={series} />
      ))}
    </Wrapper>
  );
};

export default SeriesList;
