import { useLazyQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import React from 'react';
import Modal, { Props as ModalProps } from 'react-modal';
import styled from 'styled-components';

import Newznab from '../components/Newznab';

type Props = ModalProps & {
  episode: GraphQLTypes.Episode | null;
  onClose: () => void;
};

const NEWZNAB_SEARCH = gql`
  query($id: ID!, $categories: [NewznabCategory]!) {
    nzbSearchEpisode(episodeID: $id, categories: $categories) {
      id
      title
      size
      resolution
      is_torrent
      download_url
      seeders
      peers

      parsed {
        season_number
        episode_number
        video_codec
        audio_codec
        resolution
        scene_name
      }
    }
  }
`;

const CloseButton = styled.button``;

const SearchButton = styled.button``;

const ResultsContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

const EpisodeModal: React.FC<Props> = function (props) {
  const { episode, onClose } = props;

  const [searchNewznab, { data, loading, error }] = useLazyQuery(
    NEWZNAB_SEARCH,
    {
      variables: {
        id: episode?.id,
        categories: ['TV_ALL'],
      },
    }
  );

  return (
    <Modal {...props}>
      <CloseButton onClick={onClose}>Close</CloseButton>
      <span>{episode?.title}</span>
      <SearchButton onClick={() => searchNewznab()}>
        Search for episodes
      </SearchButton>
      <ResultsContainer>
        {data?.nzbSearchEpisode?.map((newznab: GraphQLTypes.Newznab) => (
          <Newznab key={newznab.title} newznab={newznab} />
        ))}
        {data !== null && data?.nzbSearchEpisode?.length === 0 && (
          <span>No results</span>
        )}
        {loading && <span>Loading</span>}
        {error && <span>{error}</span>}
      </ResultsContainer>
    </Modal>
  );
};

export default EpisodeModal;
