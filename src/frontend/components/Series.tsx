import React from 'react';
import styled from 'styled-components';

const Wrapper = styled.div`
  display: grid;
  grid-template-areas:
    'image'
    'title';
  grid-template-columns: 1fr;
  grid-template-rows: 6fr 1fr;
  row-gap: 8px;
`;

const ImageContainer = styled.div`
  grid-area: image;
  position: relative;
`;

const Image = styled.img`
  width: 100%;
  height: 100%;
  border-radius: 3px;
  background-color: #888;
  object-fit: cover;
  box-shadow: 0px 0px 50px 0px rgba(0, 0, 0, 0.2);
`;

const Network = styled.span`
  position: absolute;
  bottom: 8px;
  left: 8px;
  padding: 5px;
  border-radius: 3px;
  background-color: #262626;
  color: white;
  font-size: 0.65em;
  font-weight: 900;
`;

const Title = styled.h3`
  grid-area: title;
  margin: 0;
`;

interface Props {
  series: GraphQLTypes.Series;
}

const Series: React.FC<Props> = function (props) {
  const { series } = props;

  return (
    <Wrapper>
      <ImageContainer>
        <Image src={`https://thetvdb.com/banners/${series.poster}`} />
        <Network>{series.network}</Network>
      </ImageContainer>
      <Title>{series.name}</Title>
    </Wrapper>
  );
};

export default React.memo(Series);
