import React from 'react';
import styled from 'styled-components';

const Wrapper = styled.div`
  border: 1px solid blue;
`;

const Title = styled.h3``;

interface Props {
  series: GraphQLTypes.Series;
}

const Series: React.FC<Props> = function (props) {
  const { series } = props;

  return (
    <Wrapper>
      <Title>{series.name}</Title>
    </Wrapper>
  );
};

export default React.memo(Series);
