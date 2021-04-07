export type BookmarkProps = {
  url: string;
  title?: string;
};

export type Bookmark = {
  id: string;
  url: string;
  title: string;
  description?: string;
  image: string;
  createdAt?: string;
  updatedAt?: string;
};

const BOOKMARKS_PATH = `${process.env.API_ENDPOINT}/bookmarks`;

export const BookmarkApi = {
  getAll: (): Promise<Bookmark[]> =>
    fetch(BOOKMARKS_PATH).then((response) => {
      if (response.ok) {
        return response.json().then((data) => data.data);
      }
    }),
  save: (b: BookmarkProps): Promise<Bookmark> =>
    fetch(BOOKMARKS_PATH, { method: 'POST', body: JSON.stringify(b) }).then((response) => {
      if (response.ok) {
        return response.json().then((data) => data);
      }
    }),
  delete: ({ id }: Bookmark): Promise<boolean> =>
    fetch(`${BOOKMARKS_PATH}/${id}`, { method: 'DELETE' }).then((response) => {
      return response.ok;
    }),
};
