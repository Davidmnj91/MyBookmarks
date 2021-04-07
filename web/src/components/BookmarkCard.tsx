import React from 'react';
import { Bookmark } from '../domain/bookmark';
import styled from 'styled-components';

type BookmarkProps = {
  onRemove: () => void;
} & Bookmark;

export const BookmarkCard: React.FC<BookmarkProps> = (props: BookmarkProps) => {
  return (
    <Panel href={props.url} target="_blank" rel="noreferrer">
      <Picture src={props.image} alt={props.title} />
      <TitleAndDescriptionPanel>
        <Title>{props.title}</Title>
        {props?.description && <Description>{props.description}</Description>}
        {/*<button onClick={props.onRemove}>D</button>*/}
      </TitleAndDescriptionPanel>
    </Panel>
  );
};

const Panel = styled.a`
  width: 250px;
  display: inline-block;
  text-decoration: none;
  background: rgba(255, 255, 255, 0.25);
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
  backdrop-filter: blur(4px);
  -webkit-backdrop-filter: blur(4px);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.18);
`;

const TitleAndDescriptionPanel = styled.div`
  padding: 10px;
`;

const Picture = styled.img`
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
`;

export const Title = styled.h2`
  color: #fff;
  font-size: 24px;
  font-weight: 500;
  text-shadow: 0 2px 4px rgb(71 97 206 / 36%);
`;

const Description = styled.p`
  font-style: normal;
  font-weight: 400;
  font-size: 16px;
  color: #fff;
  text-shadow: 0 2px 4px rgb(71 97 206 / 36%);
  margin-top: 10px;
`;
