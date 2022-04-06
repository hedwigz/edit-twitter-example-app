// your-app-name/src/App.js
import React from 'react';
import './App.css';
import { loadQuery, RelayEnvironmentProvider, usePreloadedQuery } from 'react-relay';
import { graphql } from "babel-plugin-relay/macro";
import Tweet from './Tweet';
import RelayEnvironment from './RelayEnvironment';

const TweetsQuery = graphql`
  query AppAllTweetsQuery {
    tweets {
      id
      content
      created
      likesCount
      author {
        name
        image_url
        created
        email
      }
      history {
        id
        editedAt
        diff
      }
    }
  }
`;

const preloadedQuery = loadQuery(RelayEnvironment, TweetsQuery);

function App() {
  const { tweets } = usePreloadedQuery(TweetsQuery, preloadedQuery);

  return (
    
      <div className="App">
        {!tweets && (
          <header className="App-header">
            <p>
              Loading...
            </p>
          </header>
        )}
        {tweets && tweets.map(t => {
          return <Tweet key={t.id} tweet={t} />
        })}
      </div>
  );
}

export default App;