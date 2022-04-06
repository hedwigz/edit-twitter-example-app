// your-app-name/src/fetchGraphQL.js
async function fetchGraphQL(text, variables, userid = -1) {
  const REACT_APP_GITHUB_AUTH_TOKEN = process.env.REACT_APP_GITHUB_AUTH_TOKEN;
  if (userid === -1) {
    const urlSearchParams = new URLSearchParams(window.location.search);
    const uid = urlSearchParams.get("userid");
    if (uid) {
      userid = uid;
    }
  }

  // Fetch data from GitHub's GraphQL API:
  const response = await fetch(`/query?userid=${userid}`, {
    method: 'POST',
    headers: {
      Authorization: `bearer ${REACT_APP_GITHUB_AUTH_TOKEN}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      query: text,
      variables,
    }),
  });

  // Get the response as JSON
  return await response.json();
}

export default fetchGraphQL;