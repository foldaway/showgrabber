import moment from 'moment';
import React, { useCallback } from 'react';
import styled from 'styled-components';

const Wrapper = styled.div`
  display: grid;
  grid-template-columns: 1fr 10fr 1fr 1fr;
  padding: 6px 0;
  align-items: center;

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

const SearchButton = styled.button``;

interface Props {
  episode: GraphQLTypes.Episode;
  onSearchClicked: (episode: GraphQLTypes.Episode) => void;
}

const Episode: React.FC<Props> = function (props) {
  const { episode, onSearchClicked } = props;

  const handleSearchClicked = useCallback(() => {
    onSearchClicked(episode);
  }, [episode, onSearchClicked]);

  return (
    <Wrapper>
      <EpisodeNumber>{episode.number}</EpisodeNumber>
      <EpisodeTitle>{episode.title}</EpisodeTitle>
      <EpisodeAirDate>{moment(episode.airDate).format('L')}</EpisodeAirDate>
      <SearchButton onClick={handleSearchClicked}>Search</SearchButton>
    </Wrapper>
  );
};

export default Episode;
