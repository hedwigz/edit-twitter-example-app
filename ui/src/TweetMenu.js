import * as React from 'react';
import Button from '@mui/material/Button';
import Menu from '@mui/material/Menu';
import MoreHoriz from '@mui/icons-material/MoreHoriz';
import MenuItem from '@mui/material/MenuItem';
import { IconButton } from '@mui/material';

export default function TweetMenuButton({ onEditClicked }) {
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div>
      <IconButton
        style={{ padding: 0 }}
        aria-controls="tweet-menu"
        aria-haspopup="true"
        aria-expanded={open ? 'true' : undefined}
        onClick={handleClick}>
        <MoreHoriz fontSize="small" />
      </IconButton>
      <Menu
        id="tweet-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
        MenuListProps={{
          'aria-labelledby': 'tweet-menu',
        }}
      >
        <MenuItem onClick={() => {
          onEditClicked();
          handleClose();
        }}>Edit</MenuItem>
      </Menu>
    </div>
  );
}