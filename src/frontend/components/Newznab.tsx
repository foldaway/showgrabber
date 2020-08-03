import React from 'react';
import styled from 'styled-components';

import { formatBytes } from '../util/formatBytes';

const Wrapper = styled.div`
  display: grid;
  grid-template-areas: 'title size metadata seeders peers';
  grid-template-columns: 5fr 1fr 2fr 1fr 1fr;
  align-items: center;
`;

const Title = styled.span`
  grid-area: title;
  font-family: monospace;
`;

const Size = styled.span`
  grid-area: size;
  font-family: monospace;
`;

const Metadata = styled.div`
  display: flex;
`;

const Seeders = styled.span`
  grid-area: seeders;
  font-family: monospace;
`;

const Peers = styled.span`
  grid-area: peers;
  font-family: monospace;
`;

const Resolution = styled.span`
  background-color: #7aaed6;
  font-family: monospace;
  border-radius: 8px;
  padding: 3px 4px;
`;

const VideoCodec = styled.span`
  background-color: #7be07d;
  font-family: monospace;
  border-radius: 8px;
  padding: 3px 4px;
`;

const SceneName = styled.span`
  background-color: #d895b4;
  font-family: monospace;
  border-radius: 8px;
  padding: 3px 4px;
`;

const ReleaseFormat = styled.span`
  background-color: #d8ae82;
  font-family: monospace;
  border-radius: 8px;
  padding: 3px 4px;
`;

interface Props {
  newznab: GraphQLTypes.Newznab;
}

const Newznab: React.FC<Props> = function (props) {
  const { newznab } = props;

  return (
    <Wrapper>
      <Title>{newznab.title}</Title>
      <Size>{formatBytes(newznab.size)}</Size>
      <Metadata>
        {newznab.parsed?.resolution && (
          <Resolution>{newznab.parsed?.resolution}</Resolution>
        )}
        {newznab.parsed?.video_codec && (
          <VideoCodec>{newznab.parsed?.video_codec}</VideoCodec>
        )}
        {newznab.parsed?.scene_name && (
          <SceneName>{newznab.parsed?.scene_name}</SceneName>
        )}
        {newznab.parsed?.release_format && (
          <ReleaseFormat>{newznab.parsed?.release_format}</ReleaseFormat>
        )}
      </Metadata>
      <Seeders>{newznab.seeders}</Seeders>
      <Peers>{newznab.peers}</Peers>
    </Wrapper>
  );
};

export default Newznab;
