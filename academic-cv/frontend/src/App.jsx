import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Hero from "./components/Hero";
import Metrics from "./components/Metrics";
import Research from "./components/Research";
import Links from "./components/Links";
import Experience from "./components/Experience";

// Placeholder for your future Publications component
const Publications = () => (
  <div className="container" style={{ padding: "2rem 0" }}>
    <h1>Publications</h1>
    <p>The publications list will be integrated here soon.</p>
  </div>
);

// New Navbar Component
const Navbar = () => (
  <nav className="navbar" style={{ display: 'flex', gap: '20px', padding: '1rem 0', borderBottom: '1px solid #eee' }}>
    <Link to="/" style={{ fontWeight: 'bold', textDecoration: 'none', color: 'inherit' }}>Home</Link>
    <Link to="/publications" style={{ textDecoration: 'none', color: 'inherit' }}>Publications</Link>
  </nav>
);

function App() {
  return (
    <Router>
      <main className="container">
        <Navbar />

        <Routes>
          {/* Main Home Page */}
          <Route path="/" element={
            <>
              <Hero />
              <Metrics />
              <Research />
              <Links />
              <Experience />
            </>
          } />

          {/* Separate Publications Page */}
          <Route path="/publications" element={<Publications />} />
        </Routes>
      </main>
    </Router>
  );
}

export default App;