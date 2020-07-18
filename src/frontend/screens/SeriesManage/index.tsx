import { useQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import moment from 'moment';
import React from 'react';
import { useParams } from 'react-router-dom';
import styled from 'styled-components';

import Series from '../../components/Series';

const SERIES_BY_ID = gql`
  query($id: ID!) {
    seriesByID(id: $id) {
      id
      name
      network
      poster
      overview

      seasons {
        id
        number

        episodes {
          id
          title
          number
          airDate
        }
      }
    }
  }
`;

const Wrapper = styled.div`
  display: grid;
  grid-template-areas:
    'series metadata'
    'seasons seasons';
  grid-template-columns: 1fr 5fr;
  grid-template-rows: auto 1fr;
  height: 100%;
`;

const StyledSeries = styled(Series)`
  grid-area: series;
`;

const Metadata = styled.div`
  grid-area: metadata;

  display: flex;
  flex-direction: column;
  padding: 0 16px;
`;

const Overview = styled.span`
  font-weight: 200;
`;

const Seasons = styled.div`
  grid-area: seasons;
  overflow-y: scroll;
`;

const Season = styled.div`
  margin-bottom: 16px;

  display: flex;
  flex-direction: column;
`;

const SeasonNumber = styled.span`
  background-color: #262626;
  padding: 4px 8px;
  color: white;
  border-radius: 3px;
  position: sticky;
  top: 0;
  font-weight: 200;
  font-size: 1.2em;
`;

const Episode = styled.div`
  display: grid;
  grid-template-columns: 1fr 10fr 1fr;
  padding: 6px 0;

  &:hover {
    background-color: #ddd;
  }
`;

const EpisodeNumber = styled.span`
  font-family: monospace;
`;

const EpisodeTitle = styled.span`
  font-size: 0.9em;
  font-weight: 200;
`;

const EpisodeAirDate = styled.span`
  font-family: monospace;
`;

const SeriesManage: React.FC = function () {
  const { id } = useParams();

  const { data } = useQuery(SERIES_BY_ID, {
    variables: {
      id,
    },
  });

  const series: GraphQLTypes.Series = data?.seriesByID;

  return (
    <Wrapper>
      {series && <StyledSeries series={series} />}
      <Metadata>
        <Overview>{series?.overview}</Overview>
      </Metadata>
      <Seasons>
        {series?.seasons?.map((season) => (
          <Season key={season.id}>
            <SeasonNumber>Season {season.number}</SeasonNumber>
            {season.episodes.map((episode) => (
              <Episode key={episode.id}>
                <EpisodeNumber>{episode.number}</EpisodeNumber>
                <EpisodeTitle>{episode.title}</EpisodeTitle>
                <EpisodeAirDate>
                  {moment(episode.airDate).format('L')}
                </EpisodeAirDate>
              </Episode>
            ))}
          </Season>
        ))}
      </Seasons>
    </Wrapper>
  );
};

export default SeriesManage;
