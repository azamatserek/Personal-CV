import { useEffect, useState } from "react";

export default function Publications() {
  const [publications, setPublications] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8081/api/publications")
      .then((res) => res.json())
      .then((data) => setPublications(data))
      .catch((err) => console.error("Failed to fetch publications:", err));
  }, []);

  return (
    <section className="section">
      <h2>Publications</h2>
      {publications.length === 0 ? (
        <p>Loading publications...</p>
      ) : (
        <ul>
          {publications.map((pub, idx) => (
            <li key={idx}>
              <strong>{pub.Title}</strong> ({pub.Year}) — {pub.Type} [{pub.Q}] <br />
              DOI: <a href={`https://doi.org/${pub.DOI}`} target="_blank">{pub.DOI}</a>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
