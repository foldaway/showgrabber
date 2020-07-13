import React from 'react';
import styled from 'styled-components';

const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
`;

const Title = styled.h1`
  font-weight: 300;
`;

const Header: React.FC = function () {
  return (
    <Wrapper>
      <Title>showgrabber</Title>
    </Wrapper>
  );
};

export default Header;
