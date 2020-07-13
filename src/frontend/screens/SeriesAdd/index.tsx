import React, { useState, useEffect } from 'react';
import styled from 'styled-components';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';
import TVDBSeries from '../../components/TVDBSeries';
import { useDebounce } from 'use-debounce';

const Wrapper = styled.div``;

const Input = styled.input``;

const ResultsContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

const SERIES_SEARCH = gql`
  query($term: String!) {
    tvdbSeriesSearch(term: $term) {
      id
      seriesName
      status
      network
      banner
      siteRating
      siteRatingCount
      overview
      summary {
        airedEpisodes
        airedSeasons
      }

      posterImages {
        thumbnail
      }
    }
  }
`;

const SeriesAdd: React.FC = function () {
  const [value, setValue] = useState('');
  const [term] = useDebounce(value, 300);

  const { data, loading, error } = useQuery<{
    tvdbSeriesSearch: GraphQLTypes.TVDBSeries[];
  }>(SERIES_SEARCH, {
    variables: {
      term,
    },
    skip: term.length === 0,
  });

  return (
    <Wrapper>
      <Input
        placeholder="Search term"
        value={value}
        onChange={(e) => {
          setValue(e.target.value);
        }}
      />
      {loading && <span>Loading</span>}
      <ResultsContainer>
        {data?.tvdbSeriesSearch?.map((series) => (
          <TVDBSeries key={series.id} series={series} />
        ))}
      </ResultsContainer>
    </Wrapper>
  );
};

export default SeriesAdd;
