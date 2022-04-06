import React, { useState } from 'react';
import IconButton from '@mui/material/IconButton';
import FavoriteBorder from '@mui/icons-material/FavoriteBorder';
import fetchGraphQL from './fetchGraphQL';
import MDEditor from '@uiw/react-md-editor';
import Author from './Author';
import TweetMenuButton from './TweetMenu';
import { graphql } from "babel-plugin-relay/macro";
import { useMutation } from 'react-relay';

const TweetUpdateContentMutation = graphql`
mutation TweetUpdateMutation($tweetID: ID!, $content: String!) {
  updateTweet(id: $tweetID, tweet: { content: $content }) {
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
    history{
      id
      editedAt
      diff
    }
  }
}
`

const likeTweetMutation = graphql`
mutation TweetLikeMutation($tweetID: ID!) {
  likeTweet(id: $tweetID) {
    id
    likesCount
  }
}
`

const Tweet = ({ tweet }) => {
  // We'll load the name of a repository, initially setting it to null
  const [content, setContent] = useState(tweet.content);
  const [isEdit, setIsEdit] = useState(false);
  const [commitLikeTweet] = useMutation(likeTweetMutation);
  const [commitUpdateTweetContent] = useMutation(TweetUpdateContentMutation);

  const saveNewContent = () => {
    fetchGraphQL(TweetUpdateContentMutation, { tweetID: tweet.id, content }).then(response => {
      const data = response.data.updateTweet;
      setContent(data.content);
      setIsEdit(false);
    }).catch(error => {
      console.error(error);
    })
  }

  if (!tweet) {
    return (
      <p>
        Loading...
      </p>
    )
  }

  return (
    <div className="tweet-wrapper">
      {tweet?.author?.image_url && <img className="user-avatar" src={tweet.author.image_url} alt="user pic" />}
      <div style={{ width: "50%" }}>
        <div style={{ display: 'flex', justifyContent: 'space-between' }}>
          <Author tweet={tweet} author={tweet.author} />
          <TweetMenuButton onEditClicked={() => setIsEdit(true)} />
        </div>
        {isEdit && <MDEditor
          value={content}
          onChange={setContent}
          preview="edit"
        />}
        {!isEdit && <MDEditor.Markdown style={{ color: 'black' }} source={content} />}
        <br />
        {!isEdit && (
          <div style={{ display: 'flex', justifyContent: 'space-evenly'}}>
            <div>
              <IconButton onClick={() => commitLikeTweet({
                variables: { tweetID: tweet.id }
              })}>
                <FavoriteBorder fontSize="small" />
              </IconButton>
              <span style={{ fontSize: 14 }}>{tweet.likesCount}</span>
            </div>
          </div>
        )}
        {isEdit && <button onClick={() => setIsEdit(false)}>Cancel</button>}
        {isEdit && <button style={{ marginLeft: 5 }} onClick={() => {
          commitUpdateTweetContent({
            variables: { tweetID: tweet.id, content },
            onCompleted: (res, errors) => {
              const data = res.updateTweet;
              setContent(data.content);
              setIsEdit(false);
            }
          })
        }} disabled={tweet.content === content}>Save</button>}
      </div>
    </div>
  );
}

export default Tweet;