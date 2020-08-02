import React from 'react';
import styled from 'styled-components';

import { formatBytes } from '../util/formatBytes';

const Wrapper = styled.div`
  display: grid;
  grid-template-areas: 'title size seeders peers';
  grid-template-columns: 5fr 1fr 1fr 1fr;
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
const Seeders = styled.span`
  grid-area: seeders;
  font-family: monospace;
`;
const Peers = styled.span`
  grid-area: peers;
  font-family: monospace;
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
      <Seeders>{newznab.seeders}</Seeders>
      <Peers>{newznab.peers}</Peers>
    </Wrapper>
  );
};

export default Newznab;
