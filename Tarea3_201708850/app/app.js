import React, { useState } from "react";
import Axios from "axios";

const MusicForm = () => {
  const [title, setTitle] = useState("");
  const [artist, setArtist] = useState("");
  const [year, setYear] = useState(0);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = {
      title,
      artist,
      year,
    };

    await Axios.post("http://localhost:8080/api/insert", data);

    setTitle("");
    setArtist("");
    setYear(0);
  };

  return (
    <div>
      <h1>Add New Music</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="title"
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <input
          type="text"
          name="artist"
          placeholder="Artist"
          value={artist}
          onChange={(e) => setArtist(e.target.value)}
        />
        <input
          type="number"
          name="year"
          placeholder="Year"
          value={year}
          onChange={(e) => setYear(e.target.value)}
        />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

const MusicTable = () => {
  const [music, setMusic] = useState([]);

  const fetchMusic = async () => {
    const response = await Axios.get("http://localhost:8080/api/select");
    setMusic(response.data);
  };

  useEffect(() => {
    fetchMusic();
  }, []);

  return (
    <div>
      <h1>Music List</h1>
      <table>
        <thead>
          <tr>
            <th>Title</th>
            <th>Artist</th>
            <th>Year</th>
          </tr>
        </thead>
        <tbody>
          {music.map((item) => (
            <tr key={item.id}>
              <td>{item.title}</td>
              <td>{item.artist}</td>
              <td>{item.year}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

const App = () => {
  return (
    <div>
      <MusicForm />
      <MusicTable />
    </div>
  );
};

export default App;
