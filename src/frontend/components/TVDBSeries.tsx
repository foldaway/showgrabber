import React from 'react';
import styled from 'styled-components';

const Wrapper = styled.div`
  display: grid;
  grid-template-areas:
    'art title'
    'art overview';
  grid-template-columns: auto 1fr;
  column-gap: 16px;
  padding: 16px;
  border: 1px solid blue;
`;

const Title = styled.h3`
  grid-area: title;
`;

const Overview = styled.span`
  grid-area: overview;
`;

const Image = styled.img`
  grid-area: art;
  background-color: #888;
  width: 200px;
`;

interface Props {
  series: GraphQLTypes.TVDBSeries;
}

const TVDBSeries: React.FC<Props> = function (props) {
  const { series } = props;

  return (
    <Wrapper>
      <Image
        src={`https://thetvdb.com/banners/${series.posterImages?.[0]?.thumbnail}`}
      />
      <Title>{series.seriesName}</Title>
      <Overview>{series.overview}</Overview>
    </Wrapper>
  );
};

export default React.memo(TVDBSeries);
