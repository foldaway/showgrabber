import React from 'react';
import styled from 'styled-components';

const Wrapper = styled.div``;

const Title = styled.span``;
const Size = styled.span``;
const Seeders = styled.span``;
const Peers = styled.span``;

interface Props {
  newznab: GraphQLTypes.Newznab;
}

const Newznab: React.FC<Props> = function (props) {
  const { newznab } = props;

  return (
    <Wrapper>
      <Title>{newznab.title}</Title>
      <Size>{newznab.size}</Size>
      <Seeders>{newznab.seeders}</Seeders>
      <Peers>{newznab.peers}</Peers>
    </Wrapper>
  );
};

export default Newznab;
