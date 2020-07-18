import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 16px 0;
`;

const Title = styled(Link)`
  text-decoration: none;
  color: black;
  font-style: italic;
  font-weight: 700;
  font-size: 4em;
  margin-bottom: 16px;
`;

const PageLink = styled(Link)`
  text-decoration: none;
  color: black;
  padding: 4px 8px;
  border-radius: 3px;

  &:hover {
    background-color: #444;
    color: white;
    cursor: pointer;
  }
`;

const PageLinksContainer = styled.div`
  display: grid;
  grid-auto-flow: column;
  column-gap: 8px;
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
