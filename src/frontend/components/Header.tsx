import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const Title = styled(Link)`
  text-decoration: none;
  color: black;
  font-style: italic;
  font-weight: 700;
  font-size: 4em;
`;

const PageLink = styled(Link)`
  text-decoration: none;
  color: black;

  &:hover {
    background-color: #444;
    color: white;
    cursor: pointer;
  }
`;

const PageLinksContainer = styled.div`
  display: flex;
  flex-direction: row;
`;

const Header: React.FC = function () {
  return (
    <Wrapper>
      <Title to="/">showgrabber</Title>
      <PageLinksContainer>
        <PageLink to="/">Home</PageLink>
        <PageLink to="/add">Add</PageLink>
      </PageLinksContainer>
    </Wrapper>
  );
};

export default Header;
