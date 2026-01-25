import React, { useEffect, useState } from 'react';

const Publications = () => {
  const [pubs, setPubs] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('http://localhost:8081/api/publications')
      .then(res => res.json())
      .then(data => {
        setPubs(data);
        setLoading(false);
      })
      .catch(err => console.error("Error fetching publications:", err));
  }, []);

  if (loading) return <div>Loading publications...</div>;

  // Logic to group publications by Year
  const groupedByYear = pubs.reduce((acc, pub) => {
    const year = pub.year;
    if (!acc[year]) acc[year] = [];
    acc[year].push(pub);
    return acc;
  }, {});

  const years = Object.keys(groupedByYear).sort((a, b) => b - a);

  return (
    <section className="publications-section">
      <h1 className="section-title">Publications</h1>
      {years.map(year => (
        <div key={year} className="year-group">
          <h2 className="year-header">{year}</h2>
          <ul className="pub-list">
            {groupedByYear[year].map(pub => (
              <li key={pub.id} className="pub-item">
                <span className="pub-type">[{pub.type}]</span>
                <span className="pub-title">{pub.title}</span>
                <span className="pub-authors">{pub.authors}</span>
                <span className="pub-venue"><em>{pub.venue}</em></span>
                {pub.link && (
                  <a href={pub.link} className="pub-link" target="_blank" rel="noreferrer">
                    [View Paper]
                  </a>
                )}
              </li>
            ))}
          </ul>
        </div>
      ))}
    </section>
  );
};

export default Publications;