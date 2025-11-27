import SettingsRoundedIcon from '@mui/icons-material/SettingsRounded';
import CalendarMonthRoundedIcon from '@mui/icons-material/CalendarMonthRounded';
import Card from '@mui/joy/Card';
import IconButton from '@mui/joy/IconButton';
import type { JSX } from 'react';
import { useNavigate } from 'react-router';

export function SideBar(): JSX.Element {
  const navigate = useNavigate();

  const navRoutes = [
    {
      path: '/',
      icon: <CalendarMonthRoundedIcon />
    }
  ]

  return (
    <>
      <Card
        style={{
          display: 'flex',
          justifyContent: 'space-between',
        }}
      >
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
          }}
        >
          {navRoutes.map((item, index) => (
            <IconButton
              key={index}
              onClick={() => navigate(item.path)}
            >
              {item.icon}
            </IconButton>
          ))}
        </div>
        <IconButton>
          <SettingsRoundedIcon />
        </IconButton>
      </Card>
    </>
  );
}
