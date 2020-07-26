import React from 'react';
import Modal, { Props as ModalProps } from 'react-modal';
import styled from 'styled-components';

type Props = ModalProps & {
  episode: GraphQLTypes.Episode | null;
  onClose: () => void;
};

const CloseButton = styled.button``;

const EpisodeModal: React.FC<Props> = function (props) {
  const { episode, onClose } = props;

  return (
    <Modal {...props}>
      <CloseButton onClick={onClose}>Close</CloseButton>
      <span>{episode?.title}</span>
    </Modal>
  );
};

export default EpisodeModal;
