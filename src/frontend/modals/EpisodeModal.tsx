import { useLazyQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import React from 'react';
import Modal, { Props as ModalProps } from 'react-modal';
import styled from 'styled-components';

import Newznab from '../components/Newznab';

type Props = ModalProps & {
  series?: GraphQLTypes.Series;
  season: GraphQLTypes.Season | null;
  episode: GraphQLTypes.Episode | null;
  onClose: () => void;
};

const NEWZNAB_SEARCH = gql`
  query($term: String!, $categories: [NewznabCategory]!) {
    nzbSearch(term: $term, categories: $categories) {
      id
      title
      size
      resolution
      is_torrent
      download_url
      seeders
      peers
    }
  }
`;

const CloseButton = styled.button``;

const SearchButton = styled.button``;

const EpisodeModal: React.FC<Props> = function (props) {
  const { series, season, episode, onClose } = props;

  const [searchNewznab, { data, loading, error }] = useLazyQuery(
    NEWZNAB_SEARCH,
    {
      variables: {
        term: `${series?.name} S${season?.number
          ?.toString()
          ?.padStart(2, '0')}E${episode?.number?.toString()?.padStart(2, '0')}`,
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
      <div>
        {data?.nzbSearch?.map((newznab: GraphQLTypes.Newznab) => (
          <Newznab key={newznab.title} newznab={newznab} />
        ))}
        {data !== null && data?.nzbSearch?.length === 0 && (
          <span>No results</span>
        )}
        {loading && <span>Loading</span>}
        {error && <span>{error}</span>}
      </div>
    </Modal>
  );
};

export default EpisodeModal;
