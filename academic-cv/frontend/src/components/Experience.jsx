import React, { useEffect, useState } from "react";

export default function Experience() {
  const [experiences, setExperiences] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Use environment variable for API URL (works in Docker) or fallback to localhost
    const apiUrl = process.env.REACT_APP_API_URL || "http://localhost:8081";

    fetch(`${apiUrl}/api/experience`)
      .then((res) => {
        if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
        return res.json();
      })
      .then((data) => {
        setExperiences(data);
      })
      .catch((err) => {
        console.error("Error fetching experiences:", err);
        setExperiences([]); // make sure it's not undefined
      })
      .finally(() => setLoading(false)); // always stop loading
  }, []);

  if (loading) return <p>Loading experiences...</p>;

  return (
    <section className="section" id="experience">
      <h2>Experience</h2>
      {experiences.length === 0 ? (
        <p>No experience data available.</p>
      ) : (
        <ul>
          {experiences.map((exp, idx) => (
            <li key={idx} className="mb-4">
              <strong>{exp.position}</strong> / {exp.organization} <br />
              <em>
                {exp.start} - {exp.end}
              </em>
              <p>{exp.description}</p>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
