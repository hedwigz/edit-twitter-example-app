import React, { useMemo, useState } from 'react';
import TimeAgo from 'javascript-time-ago';
import { Popover } from 'react-tiny-popover';
import _ from 'lodash';
import en from 'javascript-time-ago/locale/en.json'
import EditHistory from './EditHistory';

TimeAgo.addDefaultLocale(en);

const Author = ({ author, tweet }) => {
  const [popoverOpen, setPopoverOpen] = useState(false);
  const created_at = useMemo(() => {
    const timeAgo = new TimeAgo('en-IL');
    return timeAgo.format(Date.parse(tweet.created), 'round')
  }, [tweet.created]);
  const edited = _.size(tweet.history) !== 0;

  return (
    <div style={{ display: 'flex' }}>
      <Popover
        onClickOutside={() => setPopoverOpen(false)}
        isOpen={popoverOpen}
        positions={['right']} // preferred positions by priority
        content={<EditHistory tweet={tweet} editHistory={tweet.history} />}
      >
        <div onClick={() => setPopoverOpen(!popoverOpen)}>
          <span style={{ fontWeight: 'bold', color: 'black'}}>{author.name}&nbsp;</span>
          <span style={{ color: 'rgb(83, 100, 113)'}} onClick={() => setPopoverOpen(true)}>• {created_at} • {edited ? 'edited' : ''}</span>
        </div>
      </Popover>
    </div>
  )
}

export default Author;