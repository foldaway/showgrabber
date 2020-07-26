import moment from 'moment';
import React from 'react';
import styled from 'styled-components';

const Wrapper = styled.div`
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

interface Props {
  episode: GraphQLTypes.Episode;
}

const Episode: React.FC<Props> = function (props) {
  const { episode } = props;

  return (
    <Wrapper>
      <EpisodeNumber>{episode.number}</EpisodeNumber>
      <EpisodeTitle>{episode.title}</EpisodeTitle>
      <EpisodeAirDate>{moment(episode.airDate).format('L')}</EpisodeAirDate>
    </Wrapper>
  );
};

export default Episode;
