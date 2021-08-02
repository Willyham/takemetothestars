import './App.css';
import { useState } from 'react';

function App() {
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [candidates, setCandidates] = useState([]);

  const onClick = async () => {
    try {
      setIsLoading(true)
      const res = await fetch('/api/candidates?' + new URLSearchParams({
          lat: 51.509865,
          long: -0.118092,
          limit: 3,
      }))
      const data = await res.json()
      setCandidates(data)

    } catch (error) {
      setError(error)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="App">
      <header className="App-header"></header>
      <main>
        <p>
          Take me to the stars
        </p>
        {isLoading ? <div>Loading</div> : <button
          className="App-link"
          onClick={onClick}
        >
          Test
        </button>}
      
      
        {error ? <p>{error.message}</p> : null }

        {candidates.map(c => {
          return <p key={c.DistanceMetes}>
            {c.Point[0]} {", "} {c.Point[1]}
          </p>
        })}
      </main>
    </div>
  );
}

export default App;
