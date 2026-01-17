import Hero from "./components/Hero";
import Metrics from "./components/Metrics";
import Research from "./components/Research";
import Links from "./components/Links";
import Experience from "./components/Experience";

function App() {
     const experiences = [
        {
          position: "Associate Professor",
          organization: "Astana IT University",
          start: "Dec 2025",
          end: "Present",
          description: "Teaching, research, leading grants funded by Ministry."
        },
        {
          position: "Assistant Professor",
          organization: "Kazakh-British Technical University",
          start: "Sep 2024",
          end: "Dec 2025",
          description: "Lecturing, curriculum development, student supervision."
        },
        {
          position: "Senior Lecturer",
          organization: "SDU University",
          start: "Sep 2018",
          end: "Sep 2024",
          description: "Teaching software courses, research projects, mentoring students."
        }
      ];
  return (
    <main className="container">
      <Hero />
      <Metrics />
      <Research />
      <Links />
      <Experience />
    </main>
  );
}

export default App;
