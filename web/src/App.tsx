import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import { BookmarkCard, Title } from './components/BookmarkCard';
import { Bookmark, BookmarkApi, BookmarkProps } from './domain/bookmark';
import { AddBookmarkForm } from './components/AddBookmarkForm';

export const App: React.FC = () => {
  const [bookmarks, setBookmarks] = useState<Bookmark[]>([]);

  useEffect(() => {
    //BookmarkApi.getAll().then((bookmarks) => setBookmarks(bookmarks));
    setBookmarks(Array.from({ length: 30 }, () => generate()));
  }, []);

  const saveBookmark = (e: BookmarkProps) => {
    BookmarkApi.save(e).then((b) => {
      setBookmarks([...bookmarks, b]);
    });
  };

  const deleteBookmark = (b: Bookmark) =>
    BookmarkApi.delete(b).then(() => {
      setBookmarks(bookmarks.filter((bmk) => bmk.url !== b.url));
    });

  return (
    <Panel>
      <Header>
        <Title>My Bookmarks!</Title>
      </Header>
      <Main>
        <Grid>
          {bookmarks.map((b) => (
            <BookmarkCard key={b.id} {...b} onRemove={() => deleteBookmark(b)} />
          ))}
        </Grid>
      </Main>
      <Footer>
        <AddBookmarkForm onSubmit={saveBookmark} />
      </Footer>
    </Panel>
  );
};

const Panel = styled.div`
  width: 100%;
  height: 100vh;
  background: linear-gradient(to right, #0f2027, #203a43, #2c5364);
`;

const Header = styled.div`
  height: 45px;
  padding: 0 15px;
  margin-bottom: 1px;
  box-shadow: 0 1px 3px rgb(0 0 0 / 12%), 0 1px 2px rgb(0 0 0 / 24%);
  display: flex;
  align-items: center;
`;

const Main = styled.div`
  height: calc(100vh - 85px);
  padding: 15px;
  overflow: auto;
`;

const Grid = styled.div`
  display: grid;
  justify-items: center;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  grid-gap: 10px;
`;

const Footer = styled.div`
  height: 40px;
  width: 100%;
  padding: 5px 15px;
  position: absolute;
  bottom: 0;
`;

const generate = () => {
  const number = Math.random().toString(36).substr(2, 5);
  return {
    id: number,
    title: number,
    description: 'Lorem ipsum uashuidhsaiuh sandushi uhdudhu hs ushuji isjdo hnusdh',
    url:
      'https://medium.com/@steveruiz/using-a-javascript-library-without-type-declarations-in-a-typescript-project-3643490015f3',
    image: 'https://picsum.photos/250/100',
  };
};
