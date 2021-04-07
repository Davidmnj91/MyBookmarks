import React, { ChangeEvent, FormEvent, useState } from 'react';
import { BookmarkProps } from '../domain/bookmark';
import styled from 'styled-components';
import { AiOutlineSend } from 'react-icons/all';

type AddBookmarkFormProps = {
  onSubmit: (b: BookmarkProps) => void;
};

export const AddBookmarkForm: React.FC<AddBookmarkFormProps> = (props) => {
  const [url, setUrl] = useState('');

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setUrl(e.target.value);
  };

  const doSubmit = (e: FormEvent) => {
    e.preventDefault();
    props.onSubmit({ url });
  };

  return (
    <Form onSubmit={doSubmit}>
      <UrlInput
        type="url"
        name="url"
        placeholder="http://bookmark.url.com"
        value={url}
        onChange={handleChange}
        required
      />
      <Submit type="submit">
        <AiOutlineSend />
      </Submit>
    </Form>
  );
};

const Form = styled.form`
  display: flex;
  height: 100%;
`;

const UrlInput = styled.input`
  flex: 1;
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
  padding: 0 5px;
  background: white;
  color: black;
  box-shadow: 0 8px 32px 0 rgb(31 38 135 / 37%);
  -webkit-backdrop-filter: blur(4px);
  backdrop-filter: blur(4px);
  border: 1px solid rgba(255, 255, 255, 0.18);
  font-size: 16px;
`;

const Submit = styled.button`
  background: transparent;
  color: white;
  box-shadow: 0 1px 3px rgb(0 0 0 / 12%), 0 1px 2px rgb(0 0 0 / 24%);
  backdrop-filter: blur(4px);
  border: 1px solid rgba(255, 255, 255, 0.18);
  border-top-right-radius: 10px;
  border-bottom-right-radius: 10px;
  font-size: 22px;
  display: flex;
  align-items: center;
`;
