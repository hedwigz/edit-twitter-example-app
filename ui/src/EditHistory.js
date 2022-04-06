import React, { useMemo, useState } from 'react';
import TimeAgo from 'javascript-time-ago';
import _ from 'lodash';
import en from 'javascript-time-ago/locale/en.json'
import { Popover } from 'react-tiny-popover';
import MDEditor from '@uiw/react-md-editor';
import gitDiff from 'git-diff';

TimeAgo.addDefaultLocale(en);
const timeAgo = new TimeAgo('en-IL');

const HistoryItem = ({ tweet, item }) => {
  const [isPopoverOpen, setIsPopoverOpen] = useState(false);
  const ago = timeAgo.format(Date.parse(item.editedAt), 'round');
  return (
    <Popover
      containerStyle={{backgroundColor: '#282c34'}}
      isOpen={isPopoverOpen}
      onClickOutside={() => setIsPopoverOpen(false)}
      positions={['right']} // preferred positions by priority
      content={(
        <div className="tweet-wrapper">
          <MDEditor.Markdown style={{ color: 'white' }} source={item.diff} />
          {/* <MDEditor.Markdown style={{ color: 'white' }} source={gitDiff(item.diff, tweet.content)} /> */}
        </div>
      )}
    >
      <span onClick={() => setIsPopoverOpen(!isPopoverOpen)} style={{ padding: 5, color: '#8b949e', borderBottom: 'solid 1px gray' }} title={item.diff}>{tweet.author.name} edited {ago}</span>
    </Popover>
  );
}

const EditHistory = ({ tweet, editHistory }) => {

  return (
    <div style={{ display: 'flex', flexFlow: 'column', backgroundColor: '#282c34', border: 'solid 1px gray', borderBottom: 'none' }}>
      {editHistory.map(editChange =>
        <HistoryItem key={editChange.id} item={editChange} tweet={tweet} />
      )}
    </div>
  )
}

export default EditHistory;