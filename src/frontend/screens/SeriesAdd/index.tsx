import { useMutation, useQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import React, { useCallback, useState } from 'react';
import styled from 'styled-components';
import { useDebounce } from 'use-debounce';

import TVDBSeries from '../../components/TVDBSeries';

const Wrapper = styled.div`
  display: grid;
  grid-auto-flow: row;
  grid-template-rows: auto 1fr;
  overflow: hidden;
  height: 100%;
`;

const Input = styled.input`
  font-size: 1em;
  border-radius: 3px;
`;

const ResultsContainer = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  column-gap: 32px;
  row-gap 32px;
  margin-top: 16px;
  overflow-y: scroll;
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

const SERIES_ADD = gql`
  mutation($input: SeriesAddInput!) {
    seriesAdd(input: $input) {
      ok
      series {
        id
        status
        network
        poster
        tvdbID
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

  const [seriesAdd] = useMutation(SERIES_ADD);

  const handleAddClicked = useCallback(
    (series: GraphQLTypes.TVDBSeries) => {
      seriesAdd({
        variables: {
          input: {
            tvdbID: series.id,
          },
        },
      });
    },
    [seriesAdd]
  );

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
          <TVDBSeries
            key={series.id}
            series={series}
            onAddClicked={handleAddClicked}
          />
        ))}
      </ResultsContainer>
    </Wrapper>
  );
};

export default SeriesAdd;
