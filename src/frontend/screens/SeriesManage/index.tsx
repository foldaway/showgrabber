import { useQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import React, { useCallback, useState } from 'react';
import { useParams } from 'react-router-dom';
import styled from 'styled-components';

import Episode from '../../components/Episode';
import Series from '../../components/Series';
import EpisodeModal from '../../modals/EpisodeModal';

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

const SeriesManage: React.FC = function () {
  const { id } = useParams();

  const { data } = useQuery(SERIES_BY_ID, {
    variables: {
      id,
    },
  });

  const [modalEpisode, setModalEpisode] = useState<GraphQLTypes.Episode | null>(
    null
  );

  const series: GraphQLTypes.Series = data?.seriesByID;

  const handleEpisodeClicked = useCallback((episode: GraphQLTypes.Episode) => {
    setModalEpisode(episode);
  }, []);

  const handleModalClosed = useCallback(() => {
    setModalEpisode(null);
  }, []);

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
              <Episode
                key={episode.id}
                episode={episode}
                onSearchClicked={handleEpisodeClicked}
              />
            ))}
          </Season>
        ))}
      </Seasons>
      <EpisodeModal
        isOpen={modalEpisode !== null}
        episode={modalEpisode}
        onClose={handleModalClosed}
      />
    </Wrapper>
  );
};

export default SeriesManage;
